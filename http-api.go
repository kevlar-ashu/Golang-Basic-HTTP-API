package main

import (
	"net/http"
	"strconv"

	"encoding/json"

	"github.com/gorilla/mux"
)

// User is a struct which represents a user in our application
type User struct {
	FullName string `json:"fullName"`
	UserName string `json:"userName"`
	Email    string `json:"email"`
}

// Post is struct that represent a single post
type Post struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	Author User   `json:"author"`
}

/* type Item struct {
	Data  string `json:"data"`
	Count int    `json:"count"`
} */

//var data []Item = []Item{}
var posts []Post = []Post{}

func main() {
	router := mux.NewRouter()

	// router.HandleFunc("/test", test)

	// router.HandleFunc("/add/{item}", addItem)

	router.HandleFunc("/posts", addItem).Methods("POST")

	router.HandleFunc("/posts", getAllPosts).Methods("GET")

	router.HandleFunc("/posts/{id}", getPost).Methods("GET")

	router.HandleFunc("/posts/{id}", updatePost).Methods("PUT")

	router.HandleFunc("/posts/{id}", patchPost).Methods("PATCH")

	router.HandleFunc("/posts/{id}", deletePost).Methods("DELETE")

	http.ListenAndServe(":5000", router)
}

func getPost(w http.ResponseWriter, r *http.Request) {
	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)

	if err != nil {
		// there was an error
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to integer"))
		return
	}

	// error checking
	if id >= len(posts) {
		w.WriteHeader(404)
		w.Write([]byte("No post found with specified id"))
		return
	}

	post := posts[id]
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

func getAllPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

/* func test(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(struct {
		ID string
	}{"555"})
} */

func addItem(w http.ResponseWriter, r *http.Request) {

	/* routerVariable := mux.Vars(r)["item"]
	data = append(data, routerVariable) */

	var newPost Post
	json.NewDecoder(r.Body).Decode(&newPost)
	posts = append(posts, newPost)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(posts)
}

func updatePost(w http.ResponseWriter, r *http.Request) {
	// get the id of the post from the route parameters

	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)

	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted into integer"))
		return
	}
	// error checking
	if id >= len(posts) {
		w.WriteHeader(404)
		w.Write([]byte("No post found with specified id"))
		return
	}

	// get the value from json body
	var updatedPost Post

	json.NewDecoder(r.Body).Decode(&updatedPost)

	posts[id] = updatedPost

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedPost)

}

func patchPost(w http.ResponseWriter, r *http.Request) {

	// get the id of the post from the route parameters

	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)

	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted into integer"))
		return
	}
	// error checking
	if id >= len(posts) {
		w.WriteHeader(404)
		w.Write([]byte("No post found with specified id"))
		return
	}

	// get the current value
	/* post := posts[id]
	json.NewDecoder(r.Body).Decode(&post)
	posts[id] = post */

	post := &posts[id]
	json.NewDecoder(r.Body).Decode(post)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

func deletePost(w http.ResponseWriter, r *http.Request) {

	// get the id of the post from the route parameters

	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)

	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted into integer"))
		return
	}
	// error checking
	if id >= len(posts) {
		w.WriteHeader(404)
		w.Write([]byte("No post found with specified id"))
		return
	}

	// delete the post from the slice
	// https://github.com/golang/go/wiki/SliceTricks#delete
	posts = append(posts[:id], posts[id+1:]...)

	w.WriteHeader(200)
}
