package main

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

//Mongo blog interface code https://www.mongodb.com/blog/post/mongodb-go-driver-tutorial

func close(client *mongo.Client, ctx context.Context,
	cancel context.CancelFunc) {
	logger.Info("--> closing DB connection")
	// CancelFunc to cancel to context
	defer cancel()

	// client provides a method to close
	// a mongoDB connection.
	defer func() {

		// client.Disconnect method also has deadline.
		// returns error if any,
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

func connect(uri string) (*mongo.Client, context.Context,
	context.CancelFunc, error) {
	logger.Info("--> creating DB connection")
	// ctx will be used to set deadline for process, here
	// deadline will of 30 seconds.
	ctx, cancel := context.WithTimeout(context.Background(),
		30*time.Second)

	// mongo.Connect return mongo.Client method
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	return client, ctx, cancel, err
}

// This is a user defined method that accepts
// mongo.Client and context.Context
// This method used to ping the mongoDB, return error if any.
func ping(client *mongo.Client, ctx context.Context) error {
	logger.Info("--> running ping")

	// mongo.Client has Ping to ping mongoDB, deadline of
	// the Ping method will be determined by cxt
	// Ping method return error if any occored, then
	// the error can be handled.
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}
	logger.Info("--> DB connected successfully")
	return nil
}

func insertOne(client *mongo.Client, ctx context.Context, dataBase, col string, doc interface{}) (*mongo.InsertOneResult, error) {

	// select database and collection ith Client.Database method
	// and Database.Collection method
	collection := client.Database(dataBase).Collection(col)

	// InsertOne accept two argument of type Context
	// and of empty interface
	result, err := collection.InsertOne(ctx, doc)
	return result, err
}

func insertMany(client *mongo.Client, ctx context.Context, dataBase, col string, docs []interface{}) (*mongo.InsertManyResult, error) {

	// select database and collection ith Client.Database
	// method and Database.Collection method
	collection := client.Database(dataBase).Collection(col)

	// InsertMany accept two argument of type Context
	// and of empty interface
	result, err := collection.InsertMany(ctx, docs)
	return result, err
}

func queryAll(client *mongo.Client, ctx context.Context, dataBase, col string, query, field interface{}) (result *mongo.Cursor, err error) {

	// select database and collection.
	collection := client.Database(dataBase).Collection(col)

	// collection has an method Find,
	// that returns a mongo.cursor
	// based on query and field.
	result, err = collection.Find(ctx, bson.M{})
	return
}

func query(client *mongo.Client, ctx context.Context, dataBase, col string, query, field interface{}) (result *mongo.Cursor, err error) {

	// select database and collection.
	collection := client.Database(dataBase).Collection(col)

	// collection has an method Find,
	// that returns a mongo.cursor
	// based on query and field.
	result, err = collection.Find(ctx, query,
		options.Find().SetProjection(field))
	return
}

func writeToMongo(collection string, account string, amount string, data Traffic) (string, error) {
	client, ctx, cancel, err := connect("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}

	// Release resource when main function is returned.
	defer close(client, ctx, cancel)

	// Create  a object of type interface to  store
	// the bson values, that  we are inserting into database.
	var document interface{}

	if collection == "users" {
		document = bson.D{{"Account", account}, {"Password", data.Password}}
	} else if collection == "ledger" {
		document = bson.D{{"Account", account}, {"Amount", amount}}
	}

	// insertOne accepts client , context, database
	// name collection name and an interface that
	// will be inserted into the  collection.
	// insertOne returns an error and aresult of
	// insertina single document into the collection.
	insertOneResult, err := insertOne(client, ctx, "crypto",
		collection, document)

	// handle the error
	if err != nil {
		panic(err)
	}

	// print the insertion id of the document,
	// if it is inserted.
	//logger.Info("Result of InsertOne")
	//logger.Info(insertOneResult.InsertedID)

	return fmt.Sprintf("inserted: %s at id: %s", account, insertOneResult.InsertedID), nil
}

func writeToHashHistory(collection string, hash string, timestamp string, iteration int, previousHash string, ledger string) (string, error) {
	client, ctx, cancel, err := connect("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}

	// Release resource when main function is returned.
	defer close(client, ctx, cancel)

	// Create  a object of type interface to  store
	// the bson values, that  we are inserting into database.
	var document interface{}

	document = bson.D{{"Iteration", iteration}, {"Timestamp", timestamp}, {"Hash", hash}, {"PreviousHash", previousHash}, {"ledger", ledger}}

	// insertOne accepts client , context, database
	// name collection name and an interface that
	// will be inserted into the  collection.
	// insertOne returns an error and aresult of
	// insertina single document into the collection.
	insertOneResult, err := insertOne(client, ctx, "crypto",
		collection, document)

	// handle the error
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("hashed ledger: %s", insertOneResult.InsertedID), nil
}

func queryMongo(traffic Traffic) (returnedResults bson.D, err error) {

	// Get Client, Context, CalcelFunc and err from connect method.
	client, ctx, cancel, err := connect("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}

	// Free the resource when mainn dunction is  returned
	defer close(client, ctx, cancel)

	// create a filter an option of type interface,
	// that stores bjson objects.
	var filter, option interface{}

	// filter  gets all document,
	// with maths field greater that 70
	//filter = bson.D{
	//	{"maths", bson.D{{"$gt", 70}}},
	//}

	var cursor *mongo.Cursor

	var results []bson.D

	value := traffic.SourceAccount
	collection := "ledger"
	if traffic.Role == "USER" {
		collection = "users"
	}

	filter = bson.D{
		{"Account", value},
	}

	//  option remove id field from all documents
	option = bson.D{{"_id", 0}}

	if value == "hash" {
		filter = bson.D{}
		cursor, err = query(client, ctx, "crypto", collection, filter, option)
		// handle the errors.
		if err != nil {
			panic(err)
		}
	} else {
		// call the query method with client, context,
		// database name, collection  name, filter and option
		// This method returns momngo.cursor and error if any.
		cursor, err = query(client, ctx, "crypto", collection, filter, option)
		// handle the errors.
		if err != nil {
			panic(err)
		}
	}

	// to get bson object  from cursor,
	// returns error if any.
	if err := cursor.All(ctx, &results); err != nil {

		// handle the error
		panic(err)
	}

	// printing the result of query.
	if results == nil {
		myMessage := bson.D{{"message", "No Match"}}
		return myMessage, errors.New("No Match")
	}
	var docSlice []bson.D
	for _, doc := range results {
		//logger.Info(doc)
		docSlice = append(docSlice, doc)
		return doc, nil
	}
	return
}

func queryHash(iteration int) (returnedResults bson.D, err error) {

	// Get Client, Context, CalcelFunc and err from connect method.
	client, ctx, cancel, err := connect("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}

	// Free the resource when mainn dunction is  returned
	defer close(client, ctx, cancel)

	// create a filter an option of type interface,
	// that stores bjson objects.
	var filter, option interface{}

	// filter  gets all document,
	// with maths field greater that 70
	//filter = bson.D{
	//	{"maths", bson.D{{"$gt", 70}}},
	//}

	var cursor *mongo.Cursor

	var results []bson.D

	filter = bson.D{
		{"Iteration", iteration},
	}

	//  option remove id field from all documents
	option = bson.D{{"_id", 0}}

	// call the query method with client, context,
	// database name, collection  name, filter and option
	// This method returns momngo.cursor and error if any.
	cursor, err = query(client, ctx, "crypto", "hashHistory", filter, option)
	// handle the errors.
	if err != nil {
		panic(err)
	}

	// to get bson object  from cursor,
	// returns error if any.
	if err := cursor.All(ctx, &results); err != nil {

		// handle the error
		panic(err)
	}

	// printing the result of query.
	//logger.Debug("--> Query Result")
	if results == nil {
		myMessage := bson.D{{"Message", "No Match"}}
		return myMessage, errors.New("No Match")
	}
	var docSlice []bson.D
	for _, doc := range results {
		//logger.Info(doc)
		docSlice = append(docSlice, doc)
		return doc, nil
	}
	return
}

func queryMongoAll(collection string) (returnedResults []bson.D, err error) {

	// Get Client, Context, CalcelFunc and err from connect method.
	client, ctx, cancel, err := connect("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}

	// Free the resource when mainn dunction is  returned
	defer close(client, ctx, cancel)

	// create a filter an option of type interface,
	// that stores bjson objects.
	var filter, option interface{}

	// filter  gets all document,
	// with maths field greater that 70
	//filter = bson.D{
	//	{"maths", bson.D{{"$gt", 70}}},
	//}

	var cursor *mongo.Cursor

	var results []bson.D

	//  option remove id field from all documents
	option = bson.D{{"_id", 0}}

	filter = bson.D{}
	cursor, err = query(client, ctx, "crypto", collection, filter, option)
	// handle the errors.
	if err != nil {
		panic(err)
	}

	// to get bson object  from cursor,
	// returns error if any.
	if err := cursor.All(ctx, &results); err != nil {

		// handle the error
		panic(err)
	}

	return results, nil
}

func UpdateOne(client *mongo.Client, ctx context.Context, dataBase, col string, filter, update interface{}) (result *mongo.UpdateResult, err error) {

	// select the databse and the collection
	collection := client.Database(dataBase).Collection(col)

	// A single document that match with the
	// filter will get updated.
	// update contains the filed which should get updated.
	result, err = collection.UpdateOne(ctx, filter, update)
	return
}

// UpdateMany is a user defined method, that update
// a multiple document matching the filter.
// This methos accepts client, context, databse,
// collection, filter and update filter and update
// is of type interface this method returns
// UpdateResult and an error if any.
func UpdateMany(client *mongo.Client, ctx context.Context, dataBase, col string, filter, update interface{}) (result *mongo.UpdateResult, err error) {

	// select the databse and the collection
	collection := client.Database(dataBase).Collection(col)

	// All the documents that match with the filter will
	// get updated.
	// update contains the filed which should get updated.
	result, err = collection.UpdateMany(ctx, filter, update)
	return
}

func updateMongo(account string, amount string) (string, error) {
	// get Client, Context, CalcelFunc and err from connect method.
	client, ctx, cancel, err := connect("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}

	// Free the resource when main function in returned
	defer close(client, ctx, cancel)

	// filter object is used to select a single
	// document matching that matches.
	filter := bson.D{
		{"Account", account},
	}

	// The field of the document that need to updated.
	update := bson.D{
		{"$set", bson.D{
			{"Amount", amount},
		}},
	}

	// Returns result of updated document and a error.
	result, err := UpdateOne(client, ctx, "crypto", "ledger", filter, update)

	// handle error
	if err != nil {
		logger.Error("updateMongo error")
		return "updateMongo error", errors.New("updateMongo error")
	}

	// print count of documents that affected
	//logger.Info("update single document")
	logger.Debug(fmt.Sprintf("--> updated this many records %d", result.ModifiedCount))
	return fmt.Sprintf("--> updated this many records %d", result.ModifiedCount), nil
}

func deleteOne(client *mongo.Client, ctx context.Context, dataBase, col string, query interface{}) (result *mongo.DeleteResult, err error) {

	// select document and collection
	collection := client.Database(dataBase).Collection(col)

	// query is used to match a document  from the collection.
	result, err = collection.DeleteOne(ctx, query)
	return
}

// deleteMany is a user defined function that delete,
// multiple documents from the collection.
// Returns DeleteResult and an  error if any.
func deleteMany(client *mongo.Client, ctx context.Context, dataBase, col string, query interface{}) (result *mongo.DeleteResult, err error) {

	// select document and collection
	collection := client.Database(dataBase).Collection(col)

	// query is used to match  documents  from the collection.
	result, err = collection.DeleteMany(ctx, query)
	return
}

func deleteMongo(jsonResponse Traffic, ctx context.Context) {
	client, ctx, cancel, err := connect("mongodb://localhost:27017")
	collection := "ledger"
	if err != nil {
		panic(err)
	}

	//  free resource when main function is returned
	defer close(client, ctx, cancel)

	// This query delete document when the maths
	// field is greater than  60
	query := bson.D{
		{"Account", jsonResponse.SourceAccount},
	}

	if jsonResponse.Role == "USER" {
		collection = "users"
	}

	// Returns result of deletion and error
	result, err := deleteOne(client, ctx, "crypto", collection, query)

	// print the count of affected documents
	logger.Debug(fmt.Sprintf("rows affected by DeleteOne %s", result.DeletedCount))

	// This query deletes deletes documts that has
	// science field greater that 0
	query = bson.D{
		{"science", bson.D{{"$gt", 0}}},
	}

	// Returns result of deletion and error
	result, err = deleteMany(client, ctx, "crypto", "ledger", query)

	// print the count of affected documents
	logger.Debug(fmt.Sprintf("rows affected by DeleteMany %s", result.DeletedCount))
}

func connectMongo() {
	// Get Client, Context, CalcelFunc and
	// err from connect method.
	client, ctx, cancel, err := connect("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}

	// Release resource when the main
	// function is returned.
	defer close(client, ctx, cancel)

	// Ping mongoDB with Ping method
	ping(client, ctx)
}

func checkDoc(account string) {
	var coll *mongo.Collection
	var id primitive.ObjectID

	// find the document for which the _id field matches id
	// specify the Sort option to sort the documents by age
	// the first document in the sorted order will be returned
	opts := options.FindOne().SetSort(bson.D{{"Account", account}})
	var result bson.M
	err := coll.FindOne(context.TODO(), bson.D{{"_id", id}}, opts).Decode(&result)
	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in the collection
		if err == mongo.ErrNoDocuments {
			return
		}
		log.Fatal(err)
	}
	fmt.Printf("found document %v", result)
}
