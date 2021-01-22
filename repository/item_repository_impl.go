package repository

import (
	"evermos-online-store/config"
	"evermos-online-store/entity"
	"evermos-online-store/exception"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type itemRepositoryImpl struct {
	Collection *mongo.Collection
}

func NewItemRepository(database *mongo.Database) ItemRepository {
	return &itemRepositoryImpl{
		Collection: database.Collection("items"),
	}
}

func (repository *itemRepositoryImpl) Insert(item entity.Item) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	models := mongo.IndexModel{
		Keys: bson.D{{Key: "name", Value: "text"}},
	}
	_, err := repository.Collection.Indexes().CreateOne(ctx, models)

	_, err = repository.Collection.InsertOne(ctx, bson.M{
		"_id":         item.Id,
		"name":        item.Name,
		"price":       item.Price,
		"quantity":    item.Quantity,
		"is_deleted":  item.IsDeleted,
		"created_at":  item.CreatedAt,
		"modified_at": item.ModifiedAt,
	})
	exception.PanicIfNeeded(err)
}

func (repository *itemRepositoryImpl) FindAll(params ...string) (items []entity.Item) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	filters := bson.M{"is_deleted": false}
	if params[0] != "" {
		filters = bson.M{
			"$text": bson.M{
				"$search": params[0],
			},
		}
	}

	cursor, err := repository.Collection.Find(ctx, filters)
	exception.PanicIfNeeded(err)

	var documents []bson.M
	err = cursor.All(ctx, &documents)
	exception.PanicIfNeeded(err)

	for _, document := range documents {
		items = append(items, entity.Item{
			Id:         document["_id"].(string),
			Name:       document["name"].(string),
			Price:      document["price"].(int64),
			Quantity:   document["quantity"].(int32),
			IsDeleted:  document["is_deleted"].(bool),
			CreatedAt:  document["created_at"].(int64),
			ModifiedAt: document["modified_at"].(int64),
		})
	}

	return items
}

func (repository *itemRepositoryImpl) Show(params ...string) (items []entity.Item) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	filters := bson.M{"_id": params[0]}

	cursor, err := repository.Collection.Find(ctx, filters)
	// exception.PanicIfNeeded(err)

	var documents []bson.M
	err = cursor.All(ctx, &documents)
	if err != nil {
		panic(exception.ValidationError{
			Message: "data not found",
		})
	}

	for _, document := range documents {
		items = append(items, entity.Item{
			Id:         document["_id"].(string),
			Name:       document["name"].(string),
			Price:      document["price"].(int64),
			Quantity:   document["quantity"].(int32),
			IsDeleted:  document["is_deleted"].(bool),
			CreatedAt:  document["created_at"].(int64),
			ModifiedAt: document["modified_at"].(int64),
		})
	}

	return items
}

func (repository *itemRepositoryImpl) Update(item entity.Item) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	filter := bson.M{"_id": item.Id}
	update := bson.M{
		"_id":         item.Id,
		"name":        item.Name,
		"price":       item.Price,
		"quantity":    item.Quantity,
		"is_deleted":  item.IsDeleted,
		"modified_at": item.ModifiedAt,
	}

	_, err := repository.Collection.UpdateOne(ctx, filter, bson.M{"$set": update})
	exception.PanicIfNeeded(err)
}

func (repository *itemRepositoryImpl) PutDelete(item entity.Item) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	filter := bson.M{"_id": item.Id}
	update := bson.M{
		"_id":         item.Id,
		"is_deleted":  item.IsDeleted,
		"modified_at": item.ModifiedAt,
	}

	_, err := repository.Collection.UpdateOne(ctx, filter, bson.M{"$set": update})
	exception.PanicIfNeeded(err)
}
