package main

import (
	"auth/utils"
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io"
	"net/http"
	"net/url"
)

type cloudflareResponse struct {
	Success bool `json:"success"`
}

func checkLogin(details UserDetails, cloudflare string) (bool, string) {
	resp, err := http.PostForm("https://challenges.cloudflare.com/turnstile/v0/siteverify",
		url.Values{
			"secret":   {getSecretByName("cloudflare_secret")},
			"response": {cloudflare},
		})

	b, _ := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("err", err.Error())
		return false, "CAPTCHA error#"
	}

	res := cloudflareResponse{}
	json.Unmarshal(b, &res)

	if !res.Success {
		return false, "CAPTCHA error"
	}

	ctx := context.TODO()
	client, dataErr := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27018"))

	dataErr = client.Ping(ctx, nil)
	if dataErr != nil {
		panic(dataErr)
	}
	defer client.Disconnect(ctx)

	collection := client.Database("auth").Collection("users")

	password, _ := utils.HashPassword(details.Password)

	cursor, _ := collection.Find(ctx, bson.M{
		"username": details.Username,
		"password": password,
	})

	var loginUser []User

	if err := cursor.All(ctx, &loginUser); err != nil {
		return false, "idk"
	}

	if len(loginUser) > 0 {
		return true, generateToken(loginUser[0].ID, loginUser[0].Username)
	}

	return false, "wrong"
}
