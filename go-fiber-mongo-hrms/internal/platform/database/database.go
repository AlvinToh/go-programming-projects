package database

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

func Connect() {
	mongoURI := fmt.Sprintf("mongodb://%s:%s@%s:%s",
		viper.GetString("database.user"),
		viper.GetString("database.password"),
		viper.GetString("database.host"),
		viper.GetString("database.port"),
	)

	clientOptions := options.Client().ApplyURI(mongoURI).
		SetConnectTimeout(viper.GetDuration("database.timeout") * time.Second).
		SetMaxPoolSize(uint64(viper.GetInt("database.poolsize")))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, clientOptions)
	defer cancel()

	if err != nil {
		panic("Failed to connect to MongoDB: " + err.Error())
	}

	db = client.Database(viper.GetString("database.dbname"))
}

func GetDB() *mongo.Database {
	return db
}
