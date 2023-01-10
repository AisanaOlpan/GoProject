package store

import "github.com/AisanaOlpan/GoProject/internal/app/model"

//UserRepository
type UserRepository interface {
	Create(*model.User) error
	FindByEmail(string) (*model.User, error)
}
