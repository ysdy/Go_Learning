package server

import (
	"github.com/gin-gonic/gin"
	"github.com/ysdy/Go_Learning/controller"
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
		pet_ctrl := controller.PetController{Petservice: service.NewPetService()}
		p.GET("/pets", pet_ctrl.List)
	}
	return r
}
