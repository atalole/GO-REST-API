package test

import (
	"fmt"
	model "gin/models"
	"net/http"

	"gorm.io/gorm"
)

// interface
type Repository interface {
	CreateTestRepository(input *model.Tests) (*model.Tests, int)
}

// struct
// struct is likely intended to serve as a data access layer
type repository struct {
	db *gorm.DB
}

// NewRepositoryCreate  constructor function for the repository struct
// constructor use for Dependency Injection and Encapsulation

func NewRepositoryCreate(db *gorm.DB) *repository {
	return &repository{db: db}
}

// (r *repository) defines this as a method of the repository type.The repository struct likely manages database operations or data storage.
// CreateTestRepository function
// input *model.Tests input
// *model.Tests,string Return Values like a return response

// * (Dereference): Access the value stored at a pointer.
// & (Address-of): Get the memory address of a variable.

func (r *repository) CreateTestRepository(input *model.Tests) (*model.Tests, int) {
	fmt.Println("call4")

	var tests model.Tests
	db := r.db.Model(&tests)
	// creates a buffered channel in Go that can hold integers, with a buffer size of 1.
	errorCode := make(chan int, 1)

	checkTestsExist := db.Debug().Select("*").Where("Name = ?", input.Name).Find(&tests)

	fmt.Println("checkTestsExist", checkTestsExist.RowsAffected)

	//  <- receive operation
	if checkTestsExist.RowsAffected > 0 {
		errorCode <- http.StatusConflict
		return &tests, <-errorCode
	}

	tests.Name = input.Name

	addNewTest := db.Debug().Create(&tests)
	db.Commit()

	if addNewTest.Error != nil {
		errorCode <- http.StatusForbidden
		return &tests, <-errorCode
	} else {
		errorCode <- http.StatusOK
	}
	return &tests, <-errorCode
}
