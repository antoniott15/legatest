package main

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RepositoryDB struct {
	client   *mongo.Client
	Investor *mongo.Collection
	Project  *mongo.Collection
}

func New(uri string) (*RepositoryDB, error) {

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	err = client.Connect(context.Background())
	if err != nil {
		return nil, err
	}

	investor := client.Database("legagreen").Collection("investor")
	project := client.Database("legagreen").Collection("project")

	return &RepositoryDB{
		client:   client,
		Investor: investor,
		Project:  project,
	}, nil
}
