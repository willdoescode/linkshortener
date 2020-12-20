package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/fatih/color"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Url struct {
	Id  string `json:"id,omitempty"`
	Url string `json:"url,omitempty"`
}

var (
	clientOptions = options.Client().ApplyURI("mongodb://localhost:27017")
	client, e     = mongo.Connect(context.TODO(), clientOptions)
	collection    = client.Database("test").Collection("urls")
)

func init() {
	CheckError(e)
	e = client.Ping(context.TODO(), nil)
	CheckError(e)
}

func CreateShort(w http.ResponseWriter, r *http.Request) {
	var url Url
	_ = json.NewDecoder(r.Body).Decode(&url)
	color.Red(fmt.Sprintf("POST: %v", url))
	collection.InsertOne(context.TODO(), url)
}

func GetShortUrl(w http.ResponseWriter, r *http.Request) {
	var res Url
	collection.FindOne(context.TODO(), bson.D{primitive.E{Key: "id", Value: mux.Vars(r)["id"]}}).Decode(&res)
	j, _ := json.Marshal(res)
	color.Green(fmt.Sprintf("GET:  %v", res))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func CheckError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/create", CreateShort).Methods("POST")
	router.HandleFunc("/{id}", GetShortUrl).Methods("GET")
	handler := cors.Default().Handler(router)

	log.Fatal(http.ListenAndServe(":8080", handler))
}
