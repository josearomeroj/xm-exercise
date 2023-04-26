package utils

import (
	"context"
	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func ConnectToMongoDatabase(ctx context.Context, dbUrl string) (*mongo.Client, error) {
	client := options.Client().ApplyURI(dbUrl)
	m, err := mongo.Connect(ctx, client)
	if err != nil {
		return nil, err
	}

	ct, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	return m, m.Ping(ct, nil)
}

func ConnectToKafka(ctx context.Context, kafkaUrl string, kafkaTopic string) (*kafka.Conn, error) {
	conn, err := kafka.DialContext(ctx, "tcp", kafkaUrl)
	if err != nil {
		return nil, err
	}

	err = conn.CreateTopics(kafka.TopicConfig{Topic: kafkaTopic, NumPartitions: 1, ReplicationFactor: 1})
	if err == kafka.TopicAlreadyExists {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return conn, nil
}
