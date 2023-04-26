package store

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Company struct {
	Id           uuid.UUID `bson:"_id,omitempty"`
	Name         string    `bson:"name,omitempty"`
	Description  string    `bson:"description,omitempty"`
	EmployeesNum *int32    `bson:"employees_num,omitempty"`
	Registered   *bool     `bson:"registered,omitempty"`
	Type         int32     `bson:"type,omitempty"`
}

func (s *store) setUpCompanyIndexes(ctx context.Context) error {
	return nil
}

func (s *store) RemoveCompany(ctx context.Context, id uuid.UUID) error {
	c := s.companyCollection.FindOneAndDelete(ctx, bson.M{"_id": id})
	if err := c.Err(); err == mongo.ErrNoDocuments {
		return err
	} else if err != nil {
		return err
	}

	return nil
}

func (s *store) GetCompany(ctx context.Context, id uuid.UUID) (*Company, error) {
	sr := s.companyCollection.FindOne(ctx, bson.M{"_id": id})
	if err := sr.Err(); err == mongo.ErrNoDocuments {
		return nil, ErrNoExist
	} else if err != nil {
		s.log.Errorf("error getting company: %s", err)
		return nil, err
	}

	c := &Company{}
	if err := sr.Decode(&c); err != nil {
		s.log.Errorf("error decoding company: %s", err)
		return nil, err
	}

	return c, nil
}

func (s *store) CreateUpdateCompany(ctx context.Context, c *Company) (uuid.UUID, error) {
	if c.Id == uuid.Nil {
		c.Id = uuid.New()
	}

	if _, err := s.companyCollection.UpdateByID(ctx, c.Id, bson.M{"$set": c}, options.Update().SetUpsert(true)); err != nil {
		s.log.Errorf("error creating/updating company: %s", err)
		return uuid.Nil, err
	}
	return c.Id, nil
}
