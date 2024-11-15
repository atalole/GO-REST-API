package test

import (
	"fmt"
	model "gin/models"
)

type Service interface {
	CreateTestService(input *InputCreateTest) (*model.Tests, string)
}

type service struct {
	repository Repository
}

func NewServiceCreate(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) CreateTestService(input *InputCreateTest) (*model.Tests, string) {
	fmt.Println("call3")

	test := model.Tests{
		Name: input.Name,
	}

	resultCreateTest, errCreateTest := s.repository.CreateTestRepository(&test)
	return resultCreateTest, errCreateTest
}
