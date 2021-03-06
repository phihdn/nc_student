package models

import "github.com/dgrijalva/jwt-go"

// Sequence for mongodb
type Sequence struct {
	ID    string `bson:"_id"`
	Value int    `bson:"sequence_value"`
}

// Error with code and message
type Error struct {
	Code int
	Msg  string
}

// Student contains information for a student
type Student struct {
	ID        int    `json:"id" bson:"id"`
	FirstName string `json:"first_name" bson:"first_name"`
	LastName  string `json:"last_name" bson:"last_name"`
	ClassName string `json:"class_name" bson:"class_name"`
	Email     string `json:"email" bson:"email"`
	Age       int    `json:"age" bson:"age"`
}

type StudentAddRequest struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	ClassName string `json:"class_name"`
	Email     string `json:"email" validate:"email,min=4"`
}

// StudentUpdateRequest contains information for a request to update a student
type StudentUpdateRequest struct {
	ID        int    `json:"id" bson:"id"`
	FirstName string `json:"first_name" bson:"first_name"`
	LastName  string `json:"last_name" bson:"last_name"`
	ClassName string `json:"class_name" bson:"class_name"`
	Email     string `json:"email" bson:"email"`
	Age       int    `json:"age" bson:"age"`
}

type StudentDeleteRequest struct {
	ID        int    `json:"id,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	ClassName string `json:"class_name,omitempty"`
	Email     string `json:"email,omitempty"`
	Age       int    `json:"age,omitempty"`
}

type StudentSearchRequest struct {
	ID        int    `json:"id,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	ClassName string `json:"class_name,omitempty"`
	Email     string `json:"email,omitempty"`
	Age       int    `json:"age,omitempty"`
	Name      string `json:"name,omitempty"`
}

type UserClaims struct {
	UserID int    `json:"uid"`
	Phone  string `json:"p"`
	Email  string `json:"e"`
	jwt.StandardClaims
}
