package repository

import (
	"context"

	"github.com/tuanbui-n9/project-2/internal/app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	Collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) *UserRepository {
	return &UserRepository{
		Collection: db.Collection("users"),
	}
}

func (ur *UserRepository) Create(ctx context.Context, user *models.User) (*models.User, error) {
	_, err := ur.Collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *UserRepository) GetByID(ctx context.Context, id primitive.ObjectID) (*models.User, error) {
	var user models.User
	err := ur.Collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepository) Update(ctx context.Context, id primitive.ObjectID, user *models.User) error {
	_, err := ur.Collection.ReplaceOne(ctx, bson.M{"_id": id}, user)
	return err
}

func (ur *UserRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	_, err := ur.Collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
