package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/ysdy/Go_Learning/mocks"
	"github.com/ysdy/Go_Learning/service"
	"net/http/httptest"
	"testing"
)

func TestPetList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	serviceMock := mocks.NewMockPetService(ctrl)
	u := []service.Pet{
		{Id: "test_id1", Name: "test name1"},
		{Id: "test_id2", Name: "test name2"},
	}

	serviceMock.EXPECT().List().Return(u, nil)
	controller := PetController{Petservice: serviceMock}

	controller.List(c)
	assert.Equal(t, 200, c.Writer.Status())
}
