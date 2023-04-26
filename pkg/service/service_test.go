package service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	pb "github.com/josearomeroj/xm-exercise/pkg/gen/company_api"
	"github.com/josearomeroj/xm-exercise/pkg/internal/utils"
	"github.com/josearomeroj/xm-exercise/pkg/logging"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"testing"
)

const confFile = "../../example_config.yaml"

var serverStarted = false

func getConnStartServer(t *testing.T) *grpc.ClientConn {
	conf, err := LoadConfig(confFile)
	if err != nil {
		t.Fatalf("Could not load config from %s, error: %s", confFile, err)
	}
	if !serverStarted {
		testLogger := logging.NewLogger("test")
		StartService(context.Background(), conf, testLogger, testLogger, testLogger, true)
		serverStarted = true
	}

	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", conf.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial gRPC server: %v", err)
	}
	return conn
}

func TestAuthInterceptor_NoToken_or_Invalid(t *testing.T) {
	// create a new CompanyServiceClient with a connection to the test server
	conn := getConnStartServer(t)
	defer conn.Close()

	client := pb.NewCompanyServiceClient(conn)

	// create a new RemoveCompanyRequest without a token
	req := &pb.RemoveCompanyRequest{Id: uuid.New().String()}

	// call the RemoveCompany RPC method with no token
	_, err := client.RemoveCompany(context.Background(), req)

	// check that the error is an "Unauthenticated" error
	statusErr, ok := status.FromError(err)
	require.True(t, ok)
	require.Equal(t, codes.Unauthenticated, statusErr.Code())

	createCtx := metadata.AppendToOutgoingContext(context.Background(), "Authorization", "Bearer invalidtoken")
	// call the RemoveCompany RPC method with invalid token
	_, err = client.RemoveCompany(createCtx, req)
	statusErr, ok = status.FromError(err)
	require.True(t, ok)
	require.Equal(t, codes.Unauthenticated, statusErr.Code())
}

func TestCompanyService(t *testing.T) {
	// Connect to the gRPC server
	conn := getConnStartServer(t)
	defer conn.Close()

	// Create a client for the CompanyService
	companyServiceClient := pb.NewCompanyServiceClient(conn)
	// Create a client for the UserService
	userServiceClient := pb.NewUserServiceClient(conn)

	// Test logging in
	loginReq := &pb.LoginRequest{Username: "testuser", Password: "testpassword"}
	loginResp, err := userServiceClient.Login(context.Background(), loginReq)
	assert.NoError(t, err)
	assert.NotEmpty(t, loginResp.AuthToken)

	// Test creating a company
	createReq := &pb.CreateCompanyRequest{
		Name:         "Test Company",
		Description:  "A test company",
		EmployeesNum: 5,
		Registered:   true,
		Type:         pb.CompanyType_Corporations,
	}
	ctx := metadata.AppendToOutgoingContext(context.Background(), "Authorization", "Bearer "+loginResp.AuthToken)
	createResp, err := companyServiceClient.CreateCompany(ctx, createReq)
	assert.NoError(t, err)
	assert.NotEmpty(t, createResp.Id)

	// Test getting a company
	getReq := &pb.GetCompanyRequest{Id: createResp.Id}
	getResp, err := companyServiceClient.GetCompany(context.Background(), getReq)
	assert.NoError(t, err)
	assert.Equal(t,
		[]interface{}{
			createReq.Name,
			createReq.Description,
			createReq.EmployeesNum,
			createReq.Registered,
			createReq.Type,
		},
		[]interface{}{
			getResp.Name,
			*getResp.Description,
			getResp.EmployeesNum,
			getResp.Registered,
			getResp.Type,
		})

	// Test updating a company
	updateReq := &pb.UpdateCompanyRequest{
		Id:           createResp.Id,
		Name:         utils.Ref("Updated Test Company"),
		Description:  utils.Ref("An updated test company"),
		EmployeesNum: utils.Ref[int32](10),
		Registered:   utils.Ref(true),
		Type:         utils.Ref(pb.CompanyType_NonProfit),
	}
	_, err = companyServiceClient.UpdateCompany(ctx, updateReq)
	assert.NoError(t, err)

	// Test removing a company
	removeReq := &pb.RemoveCompanyRequest{Id: createResp.Id}
	_, err = companyServiceClient.RemoveCompany(ctx, removeReq)
	assert.NoError(t, err)

	_, err = companyServiceClient.GetCompany(ctx, getReq)
	assert.Error(t, err)
	s, _ := status.FromError(err)
	assert.Equal(t, s.Code(), codes.NotFound)
}

func TestCompanyService_PartialUpdate(t *testing.T) {
	conn := getConnStartServer(t)
	client := pb.NewCompanyServiceClient(conn)

	// Perform login to obtain auth token
	loginClient := pb.NewUserServiceClient(conn)
	loginReq := &pb.LoginRequest{Username: "testuser", Password: "testpassword"}
	loginRes, err := loginClient.Login(context.Background(), loginReq)
	require.NoError(t, err)

	// Set auth token in metadata
	ctx := metadata.AppendToOutgoingContext(context.Background(), "Authorization", "Bearer "+loginRes.AuthToken)

	// Create a company to update
	createReq := &pb.CreateCompanyRequest{
		Name:         "Test Company",
		Description:  "A test company",
		EmployeesNum: 50,
		Registered:   true,
		Type:         pb.CompanyType_Corporations,
	}
	createRes, err := client.CreateCompany(ctx, createReq)
	require.NoError(t, err)

	// Update only the description field
	updateReq := &pb.UpdateCompanyRequest{
		Id:           createRes.Id,
		EmployeesNum: utils.Ref[int32](51),
		Name:         utils.Ref("acme"),
	}

	// Make partial update request using the same auth token from login
	_, err = client.UpdateCompany(ctx, updateReq)
	require.NoError(t, err)

	// Retrieve the updated company and check that only the description has changed
	getReq := &pb.GetCompanyRequest{
		Id: createRes.Id,
	}
	getRes, err := client.GetCompany(ctx, getReq)
	require.NoError(t, err)

	expectedCompany := &pb.Company{
		Id:           createRes.Id,
		Name:         "acme",
		EmployeesNum: 51,
		Registered:   true,
		Type:         pb.CompanyType_Corporations,
	}

	resCompany := &pb.Company{
		Id:           getRes.Id,
		Name:         getRes.Name,
		EmployeesNum: getRes.EmployeesNum,
		Registered:   getRes.Registered,
		Type:         pb.CompanyType_Corporations,
	}

	require.Equal(t, expectedCompany, resCompany)
}
