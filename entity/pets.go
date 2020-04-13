package entity

import (
	"github.com/jinzhu/gorm"
)

type Pet struct {
	Id      string `json:"id"`
	Species string `json:"species"`
	Name    string `json:"name"`
	Age     int32  `json:"age"`
	StoreId string `json:"store_id"`
}

type PetRepository struct {
	DB *gorm.DB
}

func (p *PetRepository) Find() []Pet {
	var r []Pet
	p.DB.Find(&r)
	return r
}
