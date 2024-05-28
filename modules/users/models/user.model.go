package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
	"learn-go/connections"
)

type User struct {
}

func NewUserModel() *User {
	return &User{}
}

func (md *User) GetDatabase(connection string, collection string) *mongo.Collection {
	cnn := new(connections.MongoDBConnection)
	cnn_instance := cnn.NewConnection()
	return cnn_instance.Database(connection).Collection(collection)
}

func getColumns(isGetPassword bool) bson.D {
	columns := bson.D{
		{Key: "_id", Value: 1},
		{Key: "username", Value: 1},
		{Key: "email", Value: 1},
	}
	if isGetPassword {
		columns = append(columns, bson.E{"password", 1})
	}

	return columns
}

func (md *User) GetUser(username string, isGetPassword bool) map[string]any {
	var result bson.M
	query := bson.M{"username": username}
	columns := getColumns(isGetPassword)
	opts := options.FindOne().SetProjection(columns)
	coll := md.GetDatabase("ecomobi_cps_main", "users")
	_ = coll.FindOne(context.Background(), query, opts).Decode(&result)

	return result
}

func (md *User) CreateUser(username string, password string) bool {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	upsert := bson.M{
		"$set": bson.M{
			"username": username,
			"password": string(hashedPassword),
		},
	}

	filter := bson.M{"username": username}
	opts := options.Update().SetUpsert(true)
	coll := md.GetDatabase("ecomobi_cps_main", "users")
	_, e := coll.UpdateOne(context.Background(), filter, upsert, opts)
	if e == nil {
		return true
	}

	return false
}
