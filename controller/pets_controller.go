package controller

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ysdy/Go_Learning/service"
)

// Controller is pet controlller
type PetController struct {
	Service service.PetService
}

// List action: GET /pets
func (pc PetController) List(c *gin.Context) {
	p, err := pc.Service.List()

	if err != nil {
		c.AbortWithStatus(404)
		log.Println(err)
	} else {
		c.JSON(200, p)
	}
}
