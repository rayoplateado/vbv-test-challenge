package database

import (
	"context"
	"factorialapi/log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ICollection[T any] interface {
	FindOne(databaseName, collectionName, id string) (T, error)
	Find(databaseName, collectionName string, t T) ([]T, error)
	List(databaseName, collectionName string) ([]T, error)
	Add(databaseName, collectionName string, t T) (string, error)
}

type Collection[T any] struct {
}

func (c Collection[T]) FindOne(databaseName, collectionName, id string) (T, error) {
	log.Info("searching by id: " + id)
	var t T

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://"+os.Getenv("DB_USER")+":"+os.Getenv("DB_PASS")+"@"+os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT")))
	defer client.Disconnect(context.TODO())

	if err != nil {
		log.Error("error connecting to Database: " + err.Error())
		return t, err
	}
	collection := client.Database(databaseName).Collection(collectionName)
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		log.Error("error converting hex to objectid: " + err.Error())
		return t, err
	}

	err = collection.FindOne(context.TODO(), bson.M{"_id": objectId}).Decode(&t)

	if err != nil {
		log.Error("error fetching object by id: " + err.Error())
		return t, err
	}

	return t, nil
}

func (c Collection[T]) Find(databaseName, collectionName string, t T) ([]T, error) {
	log.Info("searching list by filter")

	var list = []T{}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://"+os.Getenv("DB_USER")+":"+os.Getenv("DB_PASS")+"@"+os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT")))
	defer client.Disconnect(context.TODO())

	if err != nil {
		log.Error("error connecting to Database: " + err.Error())
		return nil, err
	}
	collection := client.Database(databaseName).Collection(collectionName)
	filterBson, err := bson.Marshal(t)
	if err != nil {
		log.Error("error marshaling object to bson: " + err.Error())
		return nil, err
	}
	var filter bson.M
	err = bson.Unmarshal(filterBson, &filter)
	if err != nil {
		log.Error("error unmarshaling bson to primitive: " + err.Error())
		return nil, err
	}
	cur, err := collection.Find(context.TODO(), filter, options.Find())

	if err != nil {
		log.Error("error finding data: " + err.Error())
		return nil, err
	}

	for cur.Next(context.TODO()) {
		var t T
		err := cur.Decode(&t)
		if err != nil {
			log.Error("error decoding object: " + err.Error())
			return nil, err
		}
		list = append(list, t)
	}

	return list, nil
}

func (c Collection[T]) List(databaseName, collectionName string) ([]T, error) {
	log.Info("searching list")
	var list = []T{}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://"+os.Getenv("DB_USER")+":"+os.Getenv("DB_PASS")+"@"+os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT")))
	defer client.Disconnect(context.TODO())

	if err != nil {
		log.Error("error connecting to Database: " + err.Error())
		return nil, err
	}
	collection := client.Database(databaseName).Collection(collectionName)
	cur, err := collection.Find(context.TODO(), bson.D{{}}, options.Find())

	if err != nil {
		log.Error("error finding data: " + err.Error())
		return nil, err
	}

	for cur.Next(context.TODO()) {
		var t T
		err := cur.Decode(&t)
		if err != nil {
			log.Error("error decoding object: " + err.Error())
			return nil, err
		}
		list = append(list, t)
	}

	return list, nil
}

func (c Collection[T]) Add(databaseName, collectionName string, t T) (string, error) {
	log.Info("adding document")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://"+os.Getenv("DB_USER")+":"+os.Getenv("DB_PASS")+"@"+os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT")))
	defer client.Disconnect(context.TODO())

	if err != nil {
		log.Error("error connecting to Database: " + err.Error())
		return "", err
	}
	collection := client.Database(databaseName).Collection(collectionName)
	result, err := collection.InsertOne(context.TODO(), t)
	if err != nil {
		log.Error("error inserting data: " + err.Error())
		return "", err
	}

	return result.InsertedID.(primitive.ObjectID).Hex(), err
}
