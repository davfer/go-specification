package main

import (
	"context"
	"fmt"

	"github.com/davfer/go-specification"
	"github.com/davfer/go-specification/mongo/repository"
	mongoSpec "github.com/davfer/go-specification/mongo/resolver"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Id       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username"`
	Age      int                `bson:"age"`
}

type AdultUsersCriteria struct {
}

func (a *AdultUsersCriteria) IsSatisfiedBy(u any) bool {
	return a.GetPrimitive().IsSatisfiedBy(u)
}

func (a *AdultUsersCriteria) GetPrimitive() specification.Criteria {
	return specification.Attr{
		Name:       "Age",
		Value:      18,
		Comparison: specification.ComparisonGte,
	}
}

func main() {
	mongoConn, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://root:example@localhost:27017"))
	if err != nil {
		panic(err)
	}

	adultUsersCriteria := &AdultUsersCriteria{}
	user1 := &User{
		Username: "user1",
		Age:      20,
	}
	user2 := &User{
		Username: "user2",
		Age:      15,
	}

	db := mongoConn.Database("test").Collection("users")
	_, err = db.InsertOne(context.Background(), user1)
	if err != nil {
		panic(err)
	}
	_, err = db.InsertOne(context.Background(), user2)
	if err != nil {
		panic(err)
	}

	repo := repository.CriteriaRepository[*User]{
		Collection: db,
		Converter:  mongoSpec.NewMongoConverter(),
	}
	res, err := repo.Match(context.Background(), adultUsersCriteria)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Adult users: %v\n", res[0].Username) // Adult users: user1

	if _, err = db.DeleteMany(context.Background(), bson.M{}); err != nil {
		panic(err)
	}
}
