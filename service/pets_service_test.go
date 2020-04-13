package service

import (
	"regexp"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"github.com/ysdy/Go_Learning/entity"
)

func getDBMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	gdb, err := gorm.Open("mysql", db)
	if err != nil {
		return nil, nil, err
	}
	return gdb, mock, nil
}

func TestList(t *testing.T) {
	db, mock, err := getDBMock()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	db.LogMode(true)

	request := []entity.Pet{
		{Id: "test_id1", Name: "test name1", Age: 5, StoreId: "test_store_id1"},
		{Id: "test_id2", Name: "test name2", Age: 6, StoreId: "test_store_id2"},
	}

	// Mock設定
	mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM `pets`")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "age", "store_id"}).
			AddRow(request[0].Id, request[0].Name, request[0].Age, request[0].StoreId).
			AddRow(request[1].Id, request[1].Name, request[1].Age, request[1].StoreId))

	// 実行
	pet_repo := entity.PetRepository{DB: db}
	pet_service := NewPetService(pet_repo)
	response, _ := pet_service.List()

	assert.ElementsMatch(t, request, response)
}
