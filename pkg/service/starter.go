package service

import (
	"context"
	"fmt"
	"github.com/josearomeroj/xm-exercise/pkg/auth"
	api "github.com/josearomeroj/xm-exercise/pkg/gen/company_api"
	"github.com/josearomeroj/xm-exercise/pkg/internal/utils"
	"github.com/josearomeroj/xm-exercise/pkg/store"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

func StartService(
	ctx context.Context,
	conf *Config,
	serviceLogger *zap.SugaredLogger,
	dbLogger *zap.SugaredLogger,
	log *zap.SugaredLogger,
	async bool) {

	var cleanup []func()

	log.Infof("Connecting to database...")
	db, err := utils.ConnectToMongoDatabase(ctx, conf.DBConfig.Url)
	cleanup = append(cleanup, func() { db.Disconnect(ctx) })
	if err != nil {
		log.Fatalf("Could not connect to database, error: %s", err)
	}

	log.Infof("Connecting to kafka...")
	kafkaConn, err := utils.ConnectToKafka(ctx, conf.KafkaConfig.KafakaUrl, conf.KafkaConfig.KafkaTopic)
	cleanup = append(cleanup, func() { kafkaConn.Close() })
	if err != nil {
		log.Fatalf("Could not connect to kafka, error: %s", err)
	}

	log.Infof("Creating jwt manager...")
	jwtManager, err := auth.NewJWTManager(conf.AuthConfig.PrivateKey, conf.AuthConfig.PublicKey, conf.AuthConfig.JWTValidityMillis)
	if err != nil {
		log.Fatalf("Could not create jwt manager: %s", err)
	}

	store, err := store.NewMongoStore(ctx, conf.DBConfig.DbName, db, dbLogger)
	if err != nil {
		log.Fatalf("Could not create mongodb store: %s", err)
	}

	addr, err := net.ResolveTCPAddr("tcp", conf.KafkaConfig.KafakaUrl)
	if err != nil {
		log.Fatalf("error resolving kafka address: %s", err)
	}

	svc, err := NewService(ctx, jwtManager, &kafka.Writer{
		Addr:  addr,
		Topic: conf.KafkaConfig.KafkaTopic,
	}, store, serviceLogger)
	if err != nil {
		log.Fatalf("Could not create service, error: %s", err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", conf.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer(grpc.ChainUnaryInterceptor(svc.Middlewares()...))
	api.RegisterCompanyServiceServer(grpcServer, svc)
	api.RegisterUserServiceServer(grpcServer, svc)

	log.Infof("Starting service at port: %d", conf.Port)
	if async {
		go func() {
			log.Fatalf("server error: %s", grpcServer.Serve(lis))
			execute(cleanup)
		}()
	} else {
		log.Fatalf("server error: %s", grpcServer.Serve(lis))
		execute(cleanup)
	}
}

func execute(funcs []func()) {
	for _, f := range funcs {
		f()
	}
}
