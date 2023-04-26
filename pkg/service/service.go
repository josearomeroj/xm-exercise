package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/josearomeroj/xm-exercise/pkg/auth"
	api "github.com/josearomeroj/xm-exercise/pkg/gen/company_api"
	"github.com/josearomeroj/xm-exercise/pkg/internal/utils"
	"github.com/josearomeroj/xm-exercise/pkg/store"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Service interface {
	api.CompanyServiceServer
	api.UserServiceServer

	Middlewares() []grpc.UnaryServerInterceptor
}

type service struct {
	api.UnimplementedCompanyServiceServer
	api.UnimplementedUserServiceServer

	kafkaEventWriter *kafka.Writer

	jwtManager *auth.JWTManager
	store      store.Store
	log        *zap.SugaredLogger
}

func NewService(ctx context.Context,
	jwtManager *auth.JWTManager,
	kafkaEventClient *kafka.Writer,
	store store.Store,
	log *zap.SugaredLogger) (Service, error) {

	return &service{
		store:            store,
		log:              log,
		kafkaEventWriter: kafkaEventClient,
		jwtManager:       jwtManager,
	}, nil
}

func (s *service) Login(ctx context.Context, request *api.LoginRequest) (*api.LoginResponse, error) {
	// password + username has to be 16 bytes as defined in proto file
	id, _ := uuid.FromBytes(append([]byte(request.Username)[:8], []byte(request.Password)[:8]...))
	token, err := s.jwtManager.Generate(id)
	if err != nil {
		s.log.Errorf("error creating user bearer token: %s", err)
		return nil, internalError()
	}

	return &api.LoginResponse{Id: id.String(), AuthToken: token}, nil
}

func (s *service) GetCompany(ctx context.Context, request *api.GetCompanyRequest) (*api.Company, error) {
	company, err := s.store.GetCompany(ctx, uuid.MustParse(request.Id))
	if err != nil {
		if err == store.ErrNoExist {
			return nil, status.Errorf(codes.NotFound, "company not found or")
		} else {
			return nil, status.Errorf(codes.Internal, "internal error")
		}
	}

	return companyParse(company), nil
}

func (s *service) UpdateCompany(ctx context.Context, request *api.UpdateCompanyRequest) (*emptypb.Empty, error) {
	usr := GetUser(ctx)
	targetId := uuid.MustParse(request.Id)

	_, err := s.createUpdateCompany(ctx, &store.Company{
		Id:           targetId,
		Name:         utils.EmptyOrValue(request.Name),
		Description:  utils.EmptyOrValue(request.Description),
		EmployeesNum: request.EmployeesNum,
		Registered:   request.Registered,
		Type:         int32(utils.EmptyOrValue(request.Type)),
	})

	s.sendEvent(ctx, &event{PerformerId: targetId, TargetId: usr, EventType: EventUpdate})
	return &emptypb.Empty{}, err
}

func (s *service) CreateCompany(ctx context.Context, request *api.CreateCompanyRequest) (*api.Company, error) {
	id, err := s.createUpdateCompany(ctx, &store.Company{
		Id:           uuid.Nil,
		Name:         request.Name,
		Description:  request.Description,
		EmployeesNum: &request.EmployeesNum,
		Registered:   &request.Registered,
		Type:         int32(request.Type),
	})

	if err != nil {
		return nil, err
	}

	s.sendEvent(ctx, &event{PerformerId: GetUser(ctx), TargetId: id, EventType: EventCreate})
	return s.GetCompany(ctx, &api.GetCompanyRequest{Id: id.String()})
}

func (s *service) createUpdateCompany(ctx context.Context, request *store.Company) (uuid.UUID, error) {
	id, err := s.store.CreateUpdateCompany(ctx, request)
	if err != nil {
		return uuid.Nil, internalError()
	}

	return id, nil
}

func (s *service) RemoveCompany(ctx context.Context, request *api.RemoveCompanyRequest) (*emptypb.Empty, error) {
	targetId := uuid.MustParse(request.Id)
	if err := s.store.RemoveCompany(ctx, targetId); err != nil {
		if err == store.ErrNoExist {
			return nil, status.Errorf(codes.NotFound, "company with id %s does not exist", request.Id)
		} else {
			return nil, internalError()
		}
	}

	s.sendEvent(ctx, &event{PerformerId: GetUser(ctx), TargetId: targetId, EventType: EventRemove})
	return &emptypb.Empty{}, nil
}

func companyParse(c *store.Company) *api.Company {
	return &api.Company{
		Id:           c.Id.String(),
		Name:         c.Name,
		Description:  &c.Description,
		EmployeesNum: utils.EmptyOrValue(c.EmployeesNum),
		Registered:   utils.EmptyOrValue(c.Registered),
		Type:         api.CompanyType(c.Type),
	}
}

func internalError() error {
	return status.Errorf(codes.Internal, "internal error")
}

func (s *service) sendEvent(ctx context.Context, e *event) {
	err := s.kafkaEventWriter.WriteMessages(ctx, kafka.Message{Value: e.bytes()})
	if err != nil {
		s.log.Errorf("error sending kafka event: %s", err)
	} else {
		s.log.Infof("kafka event %+v sent successfully", *e)
	}
}
