package service

import (
	"github.com/ysdy/Go_Learning/db"
	"github.com/ysdy/Go_Learning/entity"
)

// User is alias of entity.Pet struct
type Pet entity.Pet

type PetService struct{}

// List is get all Pet
func (s PetService) List() ([]Pet, error) {
	db := db.GetDB()
	var u []Pet
	db.Find(&u)
	return u, nil
}
