package model

import "testing"

//TestUser...
func TestUser(t *testing.T) *User {
	return &User{
		Email:    "ajsanaolpan12@gmail.com",
		Password: "password",
	}
}
