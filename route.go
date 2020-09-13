package main

import (
	"encoding/json"
	"math/rand"
	"net/http"

	"github.com/alakiani/HelloREST/entity"
	"github.com/alakiani/HelloREST/repository"
)

var (
	repo repository.PostsRepository = repository.NewPostsRepository()
)

func getPosts(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")
	posts, err := repo.FindAll()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"error": "error getting posts"}`))
		return
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(posts)
}

func addPost(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")
	var p entity.Post
	err := json.NewDecoder(req.Body).Decode(&p)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"error": "error unmarshalling the post"}`))
		return
	}

	p.ID = rand.Int()
	repo.Save(&p)

	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(p)
}
