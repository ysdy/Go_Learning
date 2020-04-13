package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/ysdy/Go_Learning/db"
	"net/http/httptest"
	"testing"
)

func TestPetList(t *testing.T) {
	// DBが必要
	db.Init()
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	var controller PetController
	controller.List(c)
	assert.Equal(t, 200, c.Writer.Status())
}
