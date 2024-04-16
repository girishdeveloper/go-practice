package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	//"github.com/mongodb/mongo-go-driver/mongo"
)

type Teacher struct {
	Firstname   string
	Lastname    string
	Create_time time.Time
}

func main() {
	credential := options.Credential{
		Username: "admin",
		Password: "egi_123",
	}
	clientOpts := options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(credential)
	client, err := mongo.Connect(context.TODO(), clientOpts)
	//"mongodb://admin:egi_123@localhost:27017/trialdb"
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("mongodb connection", client)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	//select the collection in the database
	collection := client.Database("trialdb").Collection("school")
	//insert a record
	res, err := collection.InsertOne(context.Background(), bson.M{"firstname": "Girish", "lastname": "Madhavan", "create_time": time.Now()})
	if err != nil {
		log.Fatal(err)
	}
	id := res.InsertedID
	fmt.Printf("Row added with _id=%s\n", id)
	//Reading a document
	filter := bson.D{{"firstname", "Girish"}}
	result := Teacher{}
	err = collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", result)
	// query by _id
	filter = bson.D{{"_id", id}}
	result = Teacher{}
	err = collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", result)
	q1 := bson.M{"firstname": bson.M{"$eq": "Girish"}}
	q2 := bson.M{"lastname": bson.M{"$eq": "Madhavan"}}
	clauses := bson.A{q1, q2}
	filter = bson.D{{"$and", clauses}}
	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		raw := Teacher{}
		err = cur.Decode(&raw)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("1011 %+v\n", raw)
		if err = cur.Err(); err != nil {
			log.Fatal(err)
		}
	}
	//update an entry
	filter = bson.D{{"lastname", "Madhavan"}}
	update := bson.M{"$set": bson.M{"lastname": "Hoeslkan"}}
	opts := options.FindOneAndUpdate().SetUpsert(false)
	resDecoded := Teacher{}
	err = collection.FindOneAndUpdate(
		context.TODO(),
		filter,
		update,
		opts).Decode(&resDecoded)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("0022 %+v\n", resDecoded)
}
