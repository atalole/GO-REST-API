package route

import (
	"fmt"
	createTest "gin/controllers/test-controllers/create"
	handlerCreateStudent "gin/handlers/test-handlers/create"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitTestRoutes(db *gorm.DB, route *gin.Engine) {

	createTestRepository := createTest.NewRepositoryCreate(db)
	createTestService := createTest.NewServiceCreate(createTestRepository)
	createTestHandler := handlerCreateStudent.NewHandlerCreateTest(createTestService)

	/**
	@description All Student Route
	*/
	groupRoute := route.Group("/api/v1")
	fmt.Println("call1")
	groupRoute.POST("/test", createTestHandler.CreateTestHandler)
}
