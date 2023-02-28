package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserDetails struct {
	Username string
	Password string
}

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Username  string
	Password  string
	Email     string
	Chytmail  string
	Roles     []string
	Twofactor bool
}
