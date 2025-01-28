package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	model "github.com/Utsav_Kasvala/Netflix_backend/models" // replace 'your_project_path' with the actual path to your model package
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionstring = "connectionstringhere"

const dbname = "netflix"
const colName = "watchlist"

// connect with mongoDB
var collection *mongo.Collection

func init() {
	// client options
	clientOption := options.Client().ApplyURI(connectionstring)
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection = client.Database(dbname).Collection(colName)

	fmt.Println("Collection instance created!")

}

func insertone(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", inserted.InsertedID)
}

func updateone(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Updated the document: ", result.ModifiedCount)
}

func deleteone(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}

	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted the document: ", result.DeletedCount)
}

func deleteall() {
	result, err := collection.DeleteMany(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted all documents: ", result.DeletedCount)
}

func getall() []model.Netflix {
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var movies []model.Netflix
	if err = cursor.All(context.Background(), &movies); err != nil {
		log.Fatal(err)
	}

	return movies
}

func GetMyAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")

	allmovies := getall()
	json.NewEncoder(w).Encode(allmovies)
}

func AddMovie(w http.ResponseWriter, r *http.Request) {

	fmt.Println("AddMovie")

	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")

	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)

	insertone(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkasWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")

	params := mux.Vars(r)
	updateone(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")

	params := mux.Vars(r)
	deleteone(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")

	deleteall()
	json.NewEncoder(w).Encode("Deleted all movies")
}
func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("<h1>Hello</h1>"))
}
