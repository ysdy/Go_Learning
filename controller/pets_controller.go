package controller

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ysdy/Go_Learning/service"
)

// Controller is pet controlller
type PetController struct {
	Petservice service.PetService
}

func (pc PetController) List(c *gin.Context) {
	p, err := pc.Petservice.List()

	if err != nil {
		c.AbortWithStatus(404)
		log.Println(err)
	} else {
		c.JSON(200, p)
	}
}
