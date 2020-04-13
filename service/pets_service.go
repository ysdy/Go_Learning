package service

import (
	"github.com/ysdy/Go_Learning/entity"
)

type PetService interface {
	List() ([]entity.Pet, error)
}

type petService struct {
	petRepository entity.PetRepository
}

func NewPetService(db entity.PetRepository) PetService {
	return &petService{petRepository: db}
}

// List is get all Pet
func (s petService) List() ([]entity.Pet, error) {
	pr := s.petRepository
	u := pr.Find()
	return u, nil
}
