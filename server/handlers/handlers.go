package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fatih/color"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
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
	ClientOptions = options.Client().ApplyURI("mongodb://localhost:27017")
	Client, e     = mongo.Connect(context.TODO(), ClientOptions)
	Collection    = Client.Database("linkshort").Collection("urls")
	ctx           = context.Background()
	rdb           = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
)

func init() {
	CheckError(e)
	e = Client.Ping(context.TODO(), nil)
	CheckError(e)
}

func CheckError(e error) {
	if e != nil {
		panic(e)
	}
}

func CreateShort(w http.ResponseWriter, r *http.Request) {
	var url Url
	_ = json.NewDecoder(r.Body).Decode(&url)
	color.Red(fmt.Sprintf("POST: %v", url))
	id, err := rdb.Get(ctx, url.Url).Result()
	w.Header().Set("Content-type", "application/json")
	if err == redis.Nil {
		Collection.InsertOne(context.TODO(), url)
		err := rdb.Set(ctx, url.Url, url.Id, 0).Err()
		CheckError(err)
		w.WriteHeader(http.StatusOK)
		j, _ := json.Marshal(Url{
			Id:  url.Id,
			Url: "",
		})
		w.Write(j)
	} else {
		w.WriteHeader(http.StatusOK)
		j, _ := json.Marshal(Url{
			Id:  id,
			Url: "",
		})
		w.Write(j)
	}
}

func GetShortUrl(w http.ResponseWriter, r *http.Request) {
	var res Url
	Collection.FindOne(context.TODO(), bson.D{primitive.E{Key: "id", Value: mux.Vars(r)["id"]}}).Decode(&res)
	j, _ := json.Marshal(res)
	color.Green(fmt.Sprintf("GET:  %v", res))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
