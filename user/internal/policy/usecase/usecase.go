package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"time"
	"user/internal/domain"
	"user/pb"
	"user/pkg/logger"

	"github.com/casbin/casbin-go-client/client"
	"github.com/opentracing/opentracing-go"
)

// Policy UseCase
type policyUseCase struct {
	logger          logger.Logger
	policyRedisRepo domain.PolicyRedisRepository
	casbin          *client.Enforcer
}

// New Policy UseCase
func NewPolicyUseCase(logger logger.Logger, policyRedisRepo domain.PolicyRedisRepository, casbin *client.Enforcer) domain.PolicyUseCase {
	return &policyUseCase{logger: logger, policyRedisRepo: policyRedisRepo, casbin: casbin}
}

// *Command

// create new policy
func (u *policyUseCase) CreatePolicy(ctx context.Context, policy *pb.Policy) (bool, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "PolicyUseCase.CreatePolicy")
	defer span.Finish()

	param := []string{policy.Role, policy.Service, policy.Method}
	ruleAdded, err := u.casbin.AddPolicy(ctx, param)
	if err != nil {
		u.logger.Errorf("casbin.CreatePolicy: %v", err)
		return ruleAdded, fmt.Errorf("casbin.CreatePolicy: %v", err)
	}

	return ruleAdded, nil
}

func (u *policyUseCase) DeletePolicy(ctx context.Context, policy *pb.Policy) (bool, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "PolicyUseCase.DeletePolicy")
	defer span.Finish()

	param := []string{policy.Role, policy.Service, policy.Method}
	ruleRemoved, err := u.casbin.RemovePolicy(ctx, param)
	if err != nil {
		u.logger.Errorf("casbin.DeletePolicy: %v", err)
		return ruleRemoved, fmt.Errorf("casbin.DeletePolicy: %v", err)
	}

	return ruleRemoved, nil
}

// *Query

func (u *policyUseCase) Find(ctx context.Context, filters map[string]string, expire time.Duration) ([][]string, uint32, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "PolicyUseCase.Find")
	defer span.Finish()

	keys := make([]string, 0, len(filters))
	for k := range filters {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var filterKey string
	parsedFilters := make(map[string]interface{}, len(filters))
	var fieldIndex int32
	fieldVal := make([]string, 0, 3)

	for _, k := range keys {
		if len(filters[k]) > 0 {
			filterKey += fmt.Sprintf("%v_%v-", k, filters[k])

			if k != "limit" && k != "page" && k != "sort" {
				parsedFilters[k] = filters[k]
			}

			fieldVal = append(fieldVal, filters[k])
		}
	}
	filterKey = strings.TrimSuffix(filterKey, "-")

	fieldValues := strings.Join(fieldVal, " ")

	if _, found := parsedFilters["tmethod"]; found {
		fieldIndex = 2
	}
	if _, found := parsedFilters["service"]; found {
		fieldIndex = 1
	}
	if _, found := parsedFilters["role"]; found {
		fieldIndex = 0
	}

	var cachedPolicyList [][]string
	cachedByte, er := u.policyRedisRepo.FindByID(ctx, "policy_list:", filterKey)
	if er != nil {
		foundPolicyList, err := u.casbin.GetFilteredPolicy(ctx, fieldIndex, fieldValues)
		if err != nil {
			u.logger.Errorf("casbin.Find: %v", err)
			// return nil, 0, fmt.Errorf("policyRepo.Find: %v", err)
		}

		foundPolicyByte, err := json.Marshal(foundPolicyList)
		if err != nil {
			u.logger.Errorf("json.Marshal: %v", err)
			return nil, 0, err
		}

		err = u.policyRedisRepo.CreatePolicy(ctx, "policy_list:", filterKey, foundPolicyByte, expire)
		if err != nil {
			u.logger.Errorf("policyRedisRepo.CreatePolicy", err)
		}

		return foundPolicyList, uint32(len(foundPolicyList)), nil
	}

	err := json.Unmarshal(cachedByte, &cachedPolicyList)
	if err != nil {
		u.logger.Errorf("json.Unmarshal: %v", err)
		return nil, 0, err
	}

	return cachedPolicyList, uint32(len(cachedPolicyList)), nil
}
