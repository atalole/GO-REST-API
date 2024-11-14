package test

import model "gin/models"

type Repository interface {
	CreateTestRepository(input model.Test)
}
