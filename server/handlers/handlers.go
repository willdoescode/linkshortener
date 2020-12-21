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
	Collection    = Client.Database("test").Collection("urls")
	ctx           = context.Background()
	rdb           = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
)

func init() {
	CheckError(e)
	e = Client.Ping(context.TODO(), nil)
	CheckError(e)
}

func CheckError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func CreateShort(w http.ResponseWriter, r *http.Request) {
	var url Url
	_ = json.NewDecoder(r.Body).Decode(&url)
	color.Red(fmt.Sprintf("POST: %v", url))
	id, err := rdb.Get(ctx, url.Url).Result()
	if err == redis.Nil {
		Collection.InsertOne(context.TODO(), url)
		err := rdb.Set(ctx, url.Url, url.Id, 0).Err()
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		u := Url{
			Id:  url.Id,
			Url: "",
		}
		j, _ := json.Marshal(u)
		w.Write(j)
	} else {
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		u := Url{
			Id:  id,
			Url: "",
		}
		j, _ := json.Marshal(u)
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
