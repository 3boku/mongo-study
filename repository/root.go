package repository

import (
	"context"
	"github.com/3boku/mongo-study/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"reflect"
)

type Repository struct {
	config *config.Config

	client *mongo.Client
	db     *mongo.Database

	account      *mongo.Collection
	customer     *mongo.Collection
	transactions *mongo.Collection
}

func NewRepository(config *config.Config) (*Repository, error) {
	r := &Repository{
		config: config,
	}

	ctx := context.Background()

	opt := options.Client().ApplyURI(config.MongoDB.Uri)

	var err error
	if r.client, err = mongo.Connect(ctx, opt); err != nil {
		return nil, err
	} else if err = r.client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	r.registerCollection()

	return r, err
}

func (r *Repository) registerCollection() {
	db := r.client.Database("super-mongo")

	r.db = db
	r.account = db.Collection("account")
	r.customer = db.Collection("customer")
	r.transactions = db.Collection("transactions")
}

var decimalParser = reflect.TypeOf(decimal.Deciaml{})
