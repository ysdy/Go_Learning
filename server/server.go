package server

import (
	"github.com/gin-gonic/gin"
	"github.com/ysdy/Go_Learning/controller"
	"github.com/ysdy/Go_Learning/db"
	"github.com/ysdy/Go_Learning/entity"
	"github.com/ysdy/Go_Learning/service"
)

// Init is initialize server
func Init() {
	r := router()
	r.Run()
}

func router() *gin.Engine {
	r := gin.Default()

	p := r.Group("/api/v1")
	{
		pet_repo := entity.PetRepository{DB: db.GetDB()}
		pet_ctrl := controller.PetController{Petservice: service.NewPetService(pet_repo)}
		p.GET("/pets", pet_ctrl.List)
	}
	return r
}
