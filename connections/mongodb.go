package connections

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"learn-go/configs"
)

type MongoDBConnection struct {
}

var (
	ctx_mg      context.Context
	mongoclient *mongo.Client
)

func (us *MongoDBConnection) NewConnection() *mongo.Client {

	if mongoclient != nil {
		return mongoclient
	}
	conf := configs.LoadConfigJson(".")
	configApplication := conf["application"].(map[string]interface{})
	env := fmt.Sprintf("%v", configApplication["enviroment"])
	var db_uri string
	if env == "local" {
		db_uri = "mongodb://localhost:27017/?serverSelectionTimeoutMS=5000&connectTimeoutMS=10000&3t.uriVersion=3&3t.connection.name=local&3t.alwaysShowAuthDB=true&3t.alwaysShowDBFromUserRole=true"
	} else {
		configMongo := conf["mongodb"].(map[string]interface{})
		db_username := fmt.Sprintf("%v", configMongo["username"])
		db_password := fmt.Sprintf("%v", configMongo["password"])
		db_port := fmt.Sprintf("%v", configMongo["port"])
		db_auth := fmt.Sprintf("%v", configMongo["auth"])
		db_host := fmt.Sprintf("%v", configMongo["host"])
		db_uri = "mongodb+srv://" + db_username + ":" + db_password + "@" + db_host + ":" + db_port + "/?authMechanism=SCRAM-SHA-256&authSource=" + db_auth + "&directConnection=true&ssl=false"
	}
	mongoconn := options.Client().ApplyURI(db_uri)
	mg_conn, err_mongo := mongo.Connect(ctx_mg, mongoconn)
	mongoclient = mg_conn

	if err_mongo != nil {

	}

	return mongoclient

}
func (us *MongoDBConnection) CloseConnection(c *gin.Context) {
	fmt.Println("Connection to MongoDB closed.")
	if mongoclient == nil {
		return
	}

	mongoclient.Disconnect(c)
	mongoclient = nil

}
