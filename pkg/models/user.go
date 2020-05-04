package models

import (
	"github.com/beevik/guid"
)

//User models
type User struct {
	ID       string `bson:"_id" json:"id"`
	UserName string `bson:"user_name"`
	Password string `bson:"password"`
}

//NewUser constructor
func NewUser() (User, error) {
	newguid := guid.New()
	return User{
		ID: newguid.String(),
	}, nil
}
