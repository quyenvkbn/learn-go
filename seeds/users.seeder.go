package seeds

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
	"learn-go/modules/users/models"
)

func UserSeed() {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("Quyen@123"), bcrypt.DefaultCost)
	upsert := bson.M{
		"$set": bson.M{
			"username": "quyennv",
			"email":    "quyennv@example.com",
			"password": string(hashedPassword),
		},
	}

	filter := bson.M{"email": "quyennv@example.com"}
	opts := options.Update().SetUpsert(true)
	md := models.User{}
	coll := md.GetDatabase("ecomobi_cps_main", "users")
	coll.UpdateOne(context.Background(), filter, upsert, opts)

}
