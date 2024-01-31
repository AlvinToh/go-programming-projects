package model

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var db *mongo.Database

type Employee struct {
	ID     string `bson:"_id,omitempty"`
	Name   string `bson:"name"`
	Salary string `bson:"salary"`
	Age    string `bson:"age"`
}

func SetDB(database *mongo.Database) {
	db = database
}

func collection() *mongo.Collection {
	return db.Collection("employees")
}

func GetAllEmployees() ([]Employee, error) {
	var employees []Employee
	cursor, err := collection().Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	if err = cursor.All(context.Background(), &employees); err != nil {
		return nil, err
	}
	return employees, nil
}

func GetEmployeeById(id string) (*Employee, error) {
	var employee Employee
	objectID, _ := primitive.ObjectIDFromHex(id)
	err := collection().FindOne(context.Background(), bson.D{{Key: "_id", Value: objectID}}).Decode(&employee)
	return &employee, err
}

func (e *Employee) CreateEmployee() (*Employee, error) {
	result, err := collection().InsertOne(context.Background(), e)
	if err != nil {
		return nil, err
	}
	e.ID = result.InsertedID.(primitive.ObjectID).Hex()
	return e, nil
}

func (e *Employee) UpdateEmployee() error {
	objectID, _ := primitive.ObjectIDFromHex(e.ID)
	update := bson.D{{Key: "$set", Value: bson.M{"name": e.Name, "age": e.Age, "salary": e.Salary}}}
	result := collection().FindOneAndUpdate(context.Background(), bson.D{{Key: "_id", Value: objectID}}, update)
	if err := result.Err(); err != nil {
		return err
	}
	err := result.Decode(e)
	return err
}

func DeleteEmployee(id string) error {
	objectID, _ := primitive.ObjectIDFromHex(id)
	_, err := collection().DeleteOne(context.Background(), bson.D{{Key: "_id", Value: objectID}})
	return err
}
