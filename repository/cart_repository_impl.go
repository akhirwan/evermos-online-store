package repository

import (
	"evermos-online-store/config"
	"evermos-online-store/entity"
	"evermos-online-store/exception"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type cartRepositoryImpl struct {
	Collection *mongo.Collection
}

func NewCartRepository(database *mongo.Database) CartRepository {
	return &cartRepositoryImpl{
		Collection: database.Collection("user-cart"),
	}
}

func (repository *cartRepositoryImpl) FindAll(params ...string) (carts []entity.Cart) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	filters := bson.M{}
	if params[0] != "" {
		filters = bson.M{"user_email": params[0]}
	}

	cursor, err := repository.Collection.Find(ctx, filters)
	exception.PanicIfNeeded(err)

	var documents []bson.M
	err = cursor.All(ctx, &documents)
	exception.PanicIfNeeded(err)

	for _, document := range documents {
		carts = append(carts, entity.Cart{
			Id:          document["_id"].(string),
			UserEmail:   document["user_email"].(string),
			CreatedAt:   document["created_at"].(int64),
			ModifiedAt:  document["modified_at"].(int64),
			DetailItems: document["detail_items"].([]entity.DetailItems),
		})
	}

	return carts
}

func (repository *cartRepositoryImpl) Show(params ...string) (carts []entity.Cart) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()
	log.Println(params[0])

	cursor, err := repository.Collection.Find(ctx, bson.M{"user_email": params[0]})
	exception.PanicIfNeeded(err)

	var documents []bson.M
	err = cursor.All(ctx, &documents)
	exception.PanicIfNeeded(err)

	for _, document := range documents {
		carts = append(carts, entity.Cart{
			Id:          document["_id"].(string),
			UserEmail:   document["user_email"].(string),
			CreatedAt:   document["created_at"].(int64),
			ModifiedAt:  document["modified_at"].(int64),
			DetailItems: document["detail_items"].([]entity.DetailItems),
		})
	}

	return carts
}

func (repository *cartRepositoryImpl) Insert(cart entity.Cart) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	_, err := repository.Collection.InsertOne(ctx, bson.M{
		"_id":         cart.Id,
		"user_email":  cart.UserEmail,
		"detail_item": cart.DetailItems,
		"created_at":  cart.CreatedAt,
		"modified_at": cart.ModifiedAt,
	})
	exception.PanicIfNeeded(err)
}

func (repository *cartRepositoryImpl) Update(cart entity.Cart) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	filter := bson.M{"user_email": cart.UserEmail}
	update := bson.M{
		"_id":         cart.Id,
		"user_email":  cart.UserEmail,
		"detail_item": cart.DetailItems,
		"created_at":  cart.CreatedAt,
		"modified_at": cart.ModifiedAt,
	}

	_, err := repository.Collection.UpdateOne(ctx, filter, bson.M{"$set": update})
	exception.PanicIfNeeded(err)
}

func (repository *cartRepositoryImpl) Remove(cart entity.Cart) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	_, err := repository.Collection.DeleteOne(ctx, bson.M{"_id": cart.Id})
	exception.PanicIfNeeded(err)
}
