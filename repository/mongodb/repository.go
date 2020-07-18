package mongodb

import (
	"context"
	"time"

	"github.com/jmlattanzi/rex/services/inventory-service/inventory"
	"gopkg.in/mgo.v2/bson"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type mongoRepository struct {
	client   *mongo.Client
	database string
	timeout  time.Duration
}

func newMongoClient(mongoURL string, mongoTimeout int) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(mongoTimeout)*time.Second)
	defer cancel() // properly time out

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, readpref.Primary()) // confirm that we can read from our db
	if err != nil {
		return nil, err
	}

	return client, err
}

func NewMongoRepository(mongoURL, mongoDB string, mongoTimeout int) (inventory.InventoryRepository, error) {
	repo := &mongoRepository{
		database: mongoDB,
		timeout:  time.Duration(mongoTimeout) * time.Second,
	}

	client, err := newMongoClient(mongoURL, mongoTimeout)
	if err != nil {
		return nil, errors.Wrap(err, "repository.NewMongoRepository")
	}

	repo.client = client
	return repo, nil
}

func (r *mongoRepository) Get(id string) (*inventory.Entry, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	entry := &inventory.Entry{}
	collection := r.client.Database(r.database).Collection("inventory")
	filter := bson.M{"id": id}

	err := collection.FindOne(ctx, filter).Decode(entry)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.Wrap(inventory.ErrEntryNotFound, "repository.Entry.Get")
		}
		return nil, errors.Wrap(err, "repository.Entry.Get")
	}

	return entry, nil
}

func (r *mongoRepository) Post(entry *inventory.Entry) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	collection := r.client.Database(r.database).Collection("inventory")
	// _, err := collection.InsertOne(ctx, bson.M{
	// 	"id":          entry.ID,
	// 	"name":        entry.Name,
	// 	"category":    entry.Category,
	// 	"tags":        entry.Tags,
	// 	"image_url":   entry.ImageURL,
	// 	"created_at":  entry.CreatedAt,
	// 	"modified_at": entry.ModifiedAt,
	// })

	_, err := collection.InsertOne(ctx, bson.M{
		"id":              entry.ID,
		"publisher":       entry.Publisher,
		"title":           entry.Title,
		"issue":           entry.Issue,
		"condition":       entry.Condition,
		"cover_price":     entry.CoverPrice,
		"quantity":        entry.Quantity,
		"total":           entry.Total,
		"cover_variation": entry.CoverVariation,
		"category":        entry.Category,
		"tags":            entry.Tags,
		"image_url":       entry.ImageURL,
		"created_at":      entry.CreatedAt,
		"modified_at":     entry.ModifiedAt,
	})

	if err != nil {
		return errors.Wrap(err, "repository.Entry.Post")
	}

	return nil
}

func (r *mongoRepository) Update(entry *inventory.Entry, id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	collection := r.client.Database(r.database).Collection("inventory")
	filter := bson.M{"id": id}
	update := bson.M{"$set": bson.M{
		"publisher":       entry.Publisher,
		"title":           entry.Title,
		"issue":           entry.Issue,
		"condition":       entry.Condition,
		"cover_price":     entry.CoverPrice,
		"quantity":        entry.Quantity,
		"total":           entry.Total,
		"cover_variation": entry.CoverVariation,
		"category":        entry.Category,
		"tags":            entry.Tags,
		"image_url":       entry.ImageURL,
		"created_at":      entry.CreatedAt,
		"modified_at":     entry.ModifiedAt,
	}}
	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return errors.Wrap(err, "repository.Entry.Update")
	}

	return nil
}

func (r *mongoRepository) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	collection := r.client.Database(r.database).Collection("inventory")
	filter := bson.M{"id": id}
	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return errors.Wrap(err, "repository.Entry.Delete")
	}

	return nil
}

func (r *mongoRepository) GetAll() ([]*inventory.Entry, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	collection := r.client.Database(r.database).Collection("inventory")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, errors.Wrap(err, "repository.Entry.GetAll")
	}

	var results []*inventory.Entry
	if err = cursor.All(ctx, &results); err != nil {
		return nil, errors.Wrap(err, "repository.Entry.GetAll")
	}
	return results, nil
}

func (r *mongoRepository) GetCategory(category string) ([]*inventory.Entry, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	collection := r.client.Database(r.database).Collection("inventory")
	cursor, err := collection.Find(ctx, bson.M{
		"category": category,
	})
	if err != nil {
		return nil, errors.Wrap(err, "repository.Entry.GetCategory")
	}

	var results []*inventory.Entry
	if err = cursor.All(ctx, &results); err != nil {
		return nil, errors.Wrap(err, "repository.Entry.GetCategory")
	}

	return results, nil
}
