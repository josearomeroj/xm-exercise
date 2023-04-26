package store

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

const (
	companiesCollectionName = "companies"
)

var (
	ErrNoExist = errors.New("one or more documents does not exist")
)

type Store interface {
	CreateUpdateCompany(ctx context.Context, c *Company) (uuid.UUID, error)
	GetCompany(ctx context.Context, id uuid.UUID) (*Company, error)
	RemoveCompany(ctx context.Context, id uuid.UUID) error
}

type store struct {
	client *mongo.Client

	log               *zap.SugaredLogger
	companyCollection *mongo.Collection
}

func NewMongoStore(
	ctx context.Context,
	dbName string,
	client *mongo.Client,
	log *zap.SugaredLogger) (Store, error) {
	log.Infof("creating mongodb store with database name: %s", dbName)

	db := client.Database(dbName)
	s := &store{
		client:            client,
		log:               log,
		companyCollection: db.Collection(companiesCollectionName),
	}

	if err := s.setUpCompanyIndexes(ctx); err != nil {
		return nil, err
	}
	return s, nil
}
