package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	// the db orm was generated by github.com/steebchen/prisma-client-go
	"github.com/POABOB/go-discord/prisma/db"
	"github.com/joho/godotenv"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Error loading .env file\n")
	}
	client := db.NewClient(
		db.WithDatasourceURL(os.Getenv("DATABASE_URL")),
	)
	if err := client.Prisma.Connect(); err != nil {
		fmt.Printf("%v\n", err)
		return err
	}

	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()

	ctx := context.Background()

	// // create a post
	// createdPost, err := client.Profile.CreateOne(
	// 	db.Profile.UserID.Set("Test User"),
	// 	db.Profile.Name.Set("Test User"),
	// 	db.Profile.ImageURL.Set("https://avatars.githubusercontent.com/u/38785340?v=4"),
	// 	db.Profile.Email.Set("abc@example.com"),
	// ).Exec(ctx)
	// if err != nil {
	// 	return err
	// }

	// result, _ := json.MarshalIndent(createdPost, "", "  ")
	// fmt.Printf("created post: %s\n", result)

	// find a single post
	post, err := client.Profile.FindUnique(
		db.Profile.ID.Equals("dc10f2ae-5c7b-4bb9-abd0-f5bec0d2d9be"),
	).Exec(ctx)
	if err != nil {
		return err
	}

	result, _ := json.MarshalIndent(post, "", "  ")
	fmt.Printf("post: %s\n", result)

	// for optional/nullable values, you need to check the function and create two return values
	// `desc` is a string, and `ok` is a bool whether the record is null or not. If it's null,
	// `ok` is false, and `desc` will default to Go's default values; in this case an empty string (""). Otherwise,
	// `ok` is true and `desc` will be "my description".
	desc := post.ImageURL
	fmt.Printf("The posts's description is: %s\n", desc)

	return nil
}
