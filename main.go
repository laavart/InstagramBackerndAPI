package InstagramBackerndAPI

import (
	"context"
	"encoding/json"
	fmt "fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	//"fmt"
	"httprouter"
	"log"
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
	id        primitive.ObjectID  `json:"id"                  bson:"_id"`
	caption   string              `json:"caption,omitempty"   bson:"caption"`
	imageurl  string              `json:"imageurl,omitempty"  bson:"imageurl"`
	timestamp primitive.Timestamp `json:"timestamp,omitempty" bson:"timestamp"`
}

var (
	clientOptions = options.Client().ApplyURI("mongodb+srv://user:<password>@users.thoxw.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")
	ctx, _        = context.WithTimeout(context.Background(), 10*time.Second)
	client, _     = mongo.Connect(ctx, clientOptions)
)

func addUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	user := userModel{}

	json.NewDecoder(r.Body).Decode(&user)

	user.id = primitive.NewObjectID()

	client.Database("Users").Collection("UserList").InsertOne(ctx, bson.M{})
}

func getUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uid := ps.ByName("uid")

	if !primitive.IsValidObjectID(uid) {
		w.WriteHeader(http.StatusNotFound)
	}

	filterCursor, err := client.Database("Users").Collection("UserList").Find(ctx, bson.M{"_id": uid})
	if err != nil {
		log.Fatal(err)
	}
	var user bson.M
	if err = filterCursor.All(ctx, &user); err != nil {
		log.Fatal(err)
	}
	fmt.Println(user)
}

func addPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uid := ps.ByName("uid")
	post := postModel{}

	json.NewDecoder(r.Body).Decode(&post)

	post.id = primitive.NewObjectID()

	client.Database("Users").Collection(uid).InsertOne(ctx, bson.M{})
}

func getPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uid := ps.ByName("uid")
	pid := ps.ByName("pid")

	if !primitive.IsValidObjectID(uid) || !primitive.IsValidObjectID(pid) {
		w.WriteHeader(http.StatusNotFound)
	}

	filterCursor, err := client.Database("Users").Collection(uid).Find(ctx, bson.M{"_id": pid})
	if err != nil {
		log.Fatal(err)
	}
	var post bson.M
	if err = filterCursor.All(ctx, &post); err != nil {
		log.Fatal(err)
	}
	fmt.Println(post)
}

func getAllPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	if !primitive.IsValidObjectID(id) {
		w.WriteHeader(http.StatusNotFound)
	}

	cursor, err := client.Database("Users").Collection(id).Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var post bson.M
		if err = cursor.Decode(&post); err != nil {
			log.Fatal(err)
		}
		pj, _ := json.Marshal(post)
		fmt.Fprint(w, "%s\n", pj)
		time.Sleep(2 * time.Second)
	}
}

func main() {
	router := httprouter.New()
	router.POST("/user/:id", addUser)
	router.GET("/user/:id", getUser)
	router.POST("/post/:id", addPost)
	router.GET("/post/:uid", getPost)
	router.GET("/user/post/:uid", getAllPost)

	log.Fatal(http.ListenAndServe(":8080", router))

}
