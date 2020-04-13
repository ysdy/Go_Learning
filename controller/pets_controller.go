package controller

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ysdy/Go_Learning/service"
)

// Controller is pet controlller
type PetController struct {
}

// List action: GET /pets
func (pc PetController) List(c *gin.Context) {
	var sv service.PetService
	p, err := sv.List()

	if err != nil {
		c.AbortWithStatus(404)
		log.Println(err)
	} else {
		c.JSON(200, p)
	}
}
