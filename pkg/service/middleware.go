package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	api "github.com/josearomeroj/xm-exercise/pkg/gen/company_api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const userIdKey = "userId"

func (s *service) Middlewares() []grpc.UnaryServerInterceptor {
	return []grpc.UnaryServerInterceptor{
		ValidationInterceptor,
		s.AuthInterceptor,
	}
}

func (s *service) authFunc(ctx context.Context) (context.Context, error) {
	token, err := auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, err
	}

	userId, err := s.jwtManager.Verify(token)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid auth token: %v", err)
	}

	return context.WithValue(ctx, userIdKey, userId), nil
}

func (s *service) AuthInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	switch info.FullMethod {
	case api.CompanyService_RemoveCompany, api.CompanyService_CreateCompany, api.CompanyService_UpdateCompany:
		if ctx, err := s.authFunc(ctx); err != nil {
			return nil, err
		} else {
			return handler(ctx, req)
		}
	}

	return handler(ctx, req)
}

func ValidationInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	switch v := req.(type) {
	case interface{ Validate() error }:
		if err := v.Validate(); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "validation error: %s", err)
		}
	}

	return handler(ctx, req)
}

func GetUser(ctx context.Context) uuid.UUID {
	return ctx.Value(userIdKey).(uuid.UUID)
}
