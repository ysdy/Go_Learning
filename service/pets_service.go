package service

import (
	"github.com/ysdy/Go_Learning/db"
	"github.com/ysdy/Go_Learning/entity"
)

// User is alias of entity.Pet struct
type Pet entity.Pet

// Service procides pet's behavior
//type PetService struct{}
type PetService interface {
	List() ([]Pet, error)
}

func NewPetService() PetService {
	return &petService{}
}

type petService struct {
}

// List is get all Pet
func (s petService) List() ([]Pet, error) {
	db := db.GetDB()
	var u []Pet
	db.Find(&u)
	return u, nil
}
