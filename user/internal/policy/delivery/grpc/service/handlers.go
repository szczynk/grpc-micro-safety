package service

import (
	"context"
	"strings"
	"user/config"
	"user/internal/domain"
	"user/pb"
	"user/pkg/grpc_errors"
	"user/pkg/logger"

	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc/status"
)

type policiesService struct {
	pb.UnimplementedPolicyServiceServer
	logger   logger.Logger
	cfg      *config.Config
	policyUC domain.PolicyUseCase
}

// Policy service constructor
func NewPolicyGRPCService(
	logger logger.Logger,
	cfg *config.Config,
	policyUC domain.PolicyUseCase,
) *policiesService {
	return &policiesService{
		logger:   logger,
		cfg:      cfg,
		policyUC: policyUC,
	}
}

// * Command

func (u *policiesService) CreatePolicy(ctx context.Context, r *pb.CreatePolicyRequest) (*pb.CreatePolicyResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "policy.CreatePolicy")
	defer span.Finish()

	newPolicy := &pb.Policy{
		Role:    strings.ToLower(r.GetRole()),
		Service: r.GetService(),
		Method:  r.GetMethod(),
	}

	ruleAdded, err := u.policyUC.CreatePolicy(ctx, newPolicy)
	if err != nil {
		u.logger.Errorf("policyUC.CreatePolicy: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "CreatePolicy: %v", err)
	}
	if !ruleAdded {
		u.logger.Errorf("policyUC.CreatePolicy: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "CreatePolicy1: %v", err)
	}

	err = u.SendHeader(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.CreatePolicyResponse{Policy: newPolicy}, nil
}

func (u *policiesService) DeletePolicy(ctx context.Context, r *pb.DeletePolicyRequest) (*pb.DeletePolicyResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "role.DeletePolicy")
	defer span.Finish()

	policy := &pb.Policy{
		Role:    strings.ToLower(r.GetRole()),
		Service: r.GetService(),
		Method:  r.GetMethod(),
	}

	ruleRemoved, err := u.policyUC.DeletePolicy(ctx, policy)
	if err != nil {
		u.logger.Errorf("policyUC.DeletePolicy: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "DeletePolicy: %v", err)
	}

	err = u.SendHeader(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.DeletePolicyResponse{Res: ruleRemoved}, nil
}

// * Query

func (u *policiesService) FindPolicies(ctx context.Context, r *pb.FindPoliciesRequest) (*pb.FindPoliciesResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "policy.FindPolicies")
	defer span.Finish()

	// filters
	filters := map[string]string{
		"role":    strings.ToLower(r.GetRole()),
		"service": r.GetService(),
		"tmethod": r.GetMethod(),
	}

	policyList, totalCount, err := u.policyUC.Find(ctx, filters, u.cfg.Cache.Expire)
	if err != nil {
		u.logger.Errorf("policyUC.Find: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "policyUC.Find: %v", err)
	}

	if totalCount == 0 {
		return &pb.FindPoliciesResponse{}, nil
	}

	parsedPolicyList := make([]*pb.Policy, 0, totalCount)
	for _, policy := range policyList {
		parsedPolicyList = append(parsedPolicyList, u.PolicyModelToProto(policy))
	}

	err = u.SendHeader(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.FindPoliciesResponse{
		TotalCount: totalCount,
		Policies:   parsedPolicyList,
	}, nil
}

// * Utils

func (u *policiesService) PolicyModelToProto(policy []string) *pb.Policy {
	return &pb.Policy{
		Role:    policy[0],
		Service: policy[1],
		Method:  policy[2],
	}
}
