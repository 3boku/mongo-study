package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID         primitive.ObjectID `bson:"_id"`
	Frist_name *string            `json:"first_name" validate:"required, min=2, max=100"`
	Last_name
	Password
	Email
	Phone
	Token
	User_type
	Refresh_token
	Created_at
	Updated_at
	User_id
}

func userHandroller