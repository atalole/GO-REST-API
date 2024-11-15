package test

import (
	"fmt"
	model "gin/models"
)

type Service interface {
	CreateTestService(input *InputCreateTest) (*model.Tests, int)
}

type service struct {
	repository Repository
}

func NewServiceCreate(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) CreateTestService(input *InputCreateTest) (*model.Tests, int) {
	fmt.Println("call3")

	test := model.Tests{
		Name: input.Name,
	}

	fmt.Println("test", test)
	resultCreateTest, errCreateTest := s.repository.CreateTestRepository(&test)

	fmt.Println("resultCreateTesterrCreateTest", errCreateTest)
	return resultCreateTest, errCreateTest
}
