package InstagramBackerndAPI

import (
	"context"
	"encoding/json"
	"httprouter-master/httprouter-master"
	"log"
	"mongo-go-driver-master/mongo-go-driver-master/bson"
	"mongo-go-driver-master/mongo-go-driver-master/bson/primitive"
	"mongo-go-driver-master/mongo-go-driver-master/mongo"
	"mongo-go-driver-master/mongo-go-driver-master/mongo/options"
	"net/http"
	"time"
)

type userModel struct {
	id       primitive.ObjectID `json:"id"                 bson:"_id"`
	name     string             `json:"name,omitempty"     bson:"name"`
	email    string             `json:"email,omitempty"    bson:"email"`
	password string             `json:"password,omitempty" bson:"password"`
}

type postModel struct {
	id        primitive.ObjectID `json:"id"                  bson:"_id"`
	caption   string             `json:"caption,omitempty"   bson:"caption"`
	imageurl  string             `json:"imageurl,omitempty"  bson:"imageurl"`
	timestamp string             `json:"timestamp,omitempty" bson:"timestamp"`
}

var (
	clientOptions = options.Client().ApplyURI("mongodb+srv://user:<password>@users.thoxw.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")
	ctx, _        = context.WithTimeout(context.Background(), 10*time.Second)
	client, _     = mongo.Connect(ctx, clientOptions)
)

func addUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	if !bson.TypeObjectID(id) {
		w.WriteHeader(http.StatusNotFound)
	}
}

func getUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
}

func addPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
}

func getPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
}

func getAllPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
}

func main() {
	router := httprouter.New()
	router.POST("/adduser/:id", addUser)
	router.GET("/user/:id", getUser)
	router.POST("/addpost/:id", addPost)
	router.GET("/post/:uid", getPost)
	router.GET("/post/:uid", getAllPost)

	log.Fatal(http.ListenAndServe(":8080", router))

}
