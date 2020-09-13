package repository

import (
	"context"
	"log"

	"../entity"
	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"
)

const (
	projectID      = "GoREST"
	collectionName = "posts"
)

//PostsRepository .. interface
type PostsRepository interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}

type repo struct{}

//NewPostsRepository .. factory Method
func NewPostsRepository() PostsRepository {
	return &repo{}
}

func (*repo) Save(post *entity.Post) (*entity.Post, error) {

	ctx := context.Background()
	conf := &firebase.Config{ProjectID: projectID}
	app, err := firebase.NewApp(ctx, conf)

	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	client, err := app.Firestore(ctx)

	if err != nil {
		log.Fatalln(err)
	}

	defer client.Close()

	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"ID":    post.ID,
		"Title": post.Title,
		"Text":  post.Text,
	})

	if err != nil {
		log.Fatalln(err)
	}

	return post, nil
}

func (*repo) FindAll() ([]entity.Post, error) {

	ctx := context.Background()
	conf := &firebase.Config{ProjectID: projectID}
	app, err := firebase.NewApp(ctx, conf)

	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	defer client.Close()

	var posts []entity.Post

	iter := client.Collection(collectionName).Documents(ctx)

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}

		post := entity.Post{
			ID:    doc.Data()["ID"].(int),
			Title: doc.Data()["Title"].(string),
			Text:  doc.Data()["Text"].(string),
		}

		posts = append(posts, post)
	}

	return posts, nil
}
