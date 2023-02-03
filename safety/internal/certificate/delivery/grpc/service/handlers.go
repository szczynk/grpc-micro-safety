package service

import (
	"context"
	"safety/config"
	"safety/internal/domain"
	"safety/internal/models"
	"safety/pb"
	"safety/pkg/grpc_errors"
	"safety/pkg/logger"
	"safety/pkg/utils"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type certificatesService struct {
	pb.UnimplementedCertificateServiceServer
	logger        logger.Logger
	cfg           *config.Config
	certificateUC domain.CertificateUseCase
}

// Certificate service constructor
func NewCertificateGRPCService(
	logger logger.Logger,
	cfg *config.Config,
	certificateUC domain.CertificateUseCase,
) *certificatesService {
	return &certificatesService{
		logger:        logger,
		cfg:           cfg,
		certificateUC: certificateUC,
	}
}

// *Command

func (u *certificatesService) CreateCertificate(ctx context.Context, r *pb.CreateCertificateRequest) (*pb.CreateCertificateResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "certificate.CreateCertificate")
	defer span.Finish()

	parsedUserID, err := uuid.Parse(r.GetUserId())
	if err != nil {
		u.logger.Errorf("ConvertUserID: %v", err)
		return nil, status.Errorf(codes.Internal, "ConvertUserID: %v", err)
	}

	newCertificate := &models.Certificate{
		UserID:      parsedUserID,
		Dose:        r.GetDose(),
		Description: r.GetDescription(),
		ImageUrl:    r.GetImageUrl(),
	}

	createdCertificate, err := u.certificateUC.CreateCertificate(ctx, newCertificate)
	if err != nil {
		u.logger.Errorf("certificateUC.CreateCertificate: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "CreateCertificate: %v", err)
	}

	err = u.SendHeader(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.CreateCertificateResponse{Certificate: u.CertificateModelToProto(createdCertificate)}, nil
}

func (u *certificatesService) UpdateCertificateById(ctx context.Context, r *pb.UpdateCertificateByIdRequest) (*pb.UpdateCertificateByIdResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "certificate.UpdateCertificateByID")
	defer span.Finish()

	_, userRole, err := u.ExtractUserId(ctx)
	if err != nil {
		u.logger.Errorf("ExtractUserId: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "ExtractUserId: %v", err)
	}

	certificateId := r.GetId()
	// scheduleId := r.GetScheduleId()

	_, err = u.certificateUC.FindByID(ctx, certificateId, u.cfg.Cache.Expire)
	if err != nil {
		u.logger.Errorf("certificateUC.FindByID: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "certificateUC.FindByID: %v", err)
	}

	var updates models.Certificate
	switch userRole {
	case "admin":
		updates = models.Certificate{
			Description:   r.GetDescription(),
			ImageUrl:      r.GetImageUrl(),
			AdminUsername: r.GetAdminUsername(),
			Status:        r.GetStatus(),
			StatusAt:      time.Now(),
			StatusInfo:    r.GetStatusInfo(),
		}
	default:
		updates = models.Certificate{
			Description: r.GetDescription(),
			ImageUrl:    r.GetImageUrl(),
			Status:      "pending",
		}
	}

	updatedCertificate, err := u.certificateUC.UpdateByID(ctx, certificateId, updates)
	if err != nil {
		u.logger.Errorf("certificateUC.UpdateByID: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "certificateUC.UpdateByID: %v", err)
	}

	err = u.SendHeader(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateCertificateByIdResponse{Certificate: u.CertificateModelToProto(updatedCertificate)}, nil
}

func (u *certificatesService) DeleteCertificateById(ctx context.Context, r *pb.DeleteCertificateByIdRequest) (*pb.DeleteCertificateByIdResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "certificate.DeleteCertificateByID")
	defer span.Finish()

	certificateId := r.GetId()

	err := u.certificateUC.DeleteByID(ctx, certificateId)
	if err != nil {
		u.logger.Errorf("certificateUC.DeleteByID: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "certificateUC.DeleteByID: %v", err)
	}

	err = u.SendHeader(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteCertificateByIdResponse{Message: "Certificate deleted successfully"}, nil
}

// *Query

func (u *certificatesService) FindCertificates(ctx context.Context, r *pb.FindCertificatesRequest) (*pb.FindCertificatesResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "certificate.FindCertificates")
	defer span.Finish()

	// filter
	// TODO(Bagus): if user, userId from extractUserId, else, u can input userId
	userID, userRole, err := u.ExtractUserId(ctx)
	if err != nil {
		u.logger.Errorf("ExtractUserId: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "ExtractUserId: %v", err)
	}

	var userId string
	switch userRole {
	case "admin":
		userId = r.GetUserId()
	default:
		userId = userID
	}

	aStatus := r.GetStatus()

	sort := r.GetSort()
	if len(sort) == 0 {
		sort = "created_at asc"
	}

	paginateQuery := utils.NewPaginationQuery(r.GetLimit(), r.GetPage(), sort)

	filters := map[string]string{
		"user_id": userId,
		"status":  aStatus,

		"limit": strconv.FormatUint(uint64(r.GetLimit()), 10),
		"page":  strconv.FormatUint(uint64(r.GetPage()), 10),
		"sort":  sort,
	}

	certificateList, totalCount, err := u.certificateUC.Find(ctx, filters, paginateQuery, u.cfg.Cache.Expire)
	if err != nil {
		u.logger.Errorf("certificateUC.Find: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "certificateUC.Find: %v", err)
	}

	parsedCertificateList := make([]*pb.CertificateDetail, 0, totalCount)
	for _, certificate := range certificateList {
		parsedCertificateList = append(parsedCertificateList, u.CertificateDetailToProto(certificate))
	}

	err = u.SendHeader(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.FindCertificatesResponse{
		TotalCount:   totalCount,
		TotalPages:   paginateQuery.GetTotalPages(totalCount),
		Page:         paginateQuery.GetPage(),
		Limit:        paginateQuery.GetLimit(),
		HasMore:      paginateQuery.GetHasMore(totalCount),
		Certificates: parsedCertificateList,
	}, nil
}

func (u *certificatesService) FindCertificateById(ctx context.Context, r *pb.FindCertificateByIdRequest) (*pb.FindCertificateByIdResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "certificate.FindCertificateByID")
	defer span.Finish()

	certificateId := r.GetId()

	certificate, err := u.certificateUC.FindByID(ctx, certificateId, u.cfg.Cache.Expire)
	if err != nil {
		u.logger.Errorf("certificateUC.FindByID: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "certificateUC.FindByID: %v", err)
	}

	err = u.SendHeader(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.FindCertificateByIdResponse{Certificate: u.CertificateModelToProto(certificate)}, nil
}

// * Utils

func (u *certificatesService) CertificateModelToProto(certificate *models.Certificate) *pb.Certificate {
	return &pb.Certificate{
		Id:            uint32(certificate.ID),
		UserId:        certificate.UserID.String(),
		Dose:          certificate.Dose,
		ImageUrl:      certificate.ImageUrl,
		Description:   certificate.Description,
		AdminUsername: certificate.AdminUsername,
		Status:        certificate.Status,
		StatusAt:      timestamppb.New(certificate.StatusAt),
		StatusInfo:    certificate.StatusInfo,
		CreatedAt:     timestamppb.New(certificate.CreatedAt),
		UpdatedAt:     timestamppb.New(certificate.UpdatedAt),
	}
}

func (u *certificatesService) CertificateDetailToProto(certificate *models.Certificate) *pb.CertificateDetail {
	return &pb.CertificateDetail{
		Id:            uint32(certificate.ID),
		UserId:        certificate.UserID.String(),
		UserUsername:  certificate.UserUsername,
		UserAvatar:    certificate.UserAvatar,
		Dose:          certificate.Dose,
		ImageUrl:      certificate.ImageUrl,
		Description:   certificate.Description,
		AdminUsername: certificate.AdminUsername,
		Status:        certificate.Status,
		StatusAt:      timestamppb.New(certificate.StatusAt),
		StatusInfo:    certificate.StatusInfo,
		CreatedAt:     timestamppb.New(certificate.CreatedAt),
		UpdatedAt:     timestamppb.New(certificate.UpdatedAt),
	}
}
