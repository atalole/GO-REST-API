package route

import (
	createStudent "gin/controllers/student-controllers/create"
	deleteStudent "gin/controllers/student-controllers/delete"
	resultStudent "gin/controllers/student-controllers/result"
	resultsStudent "gin/controllers/student-controllers/results"
	updateStudent "gin/controllers/student-controllers/update"
	handlerCreateStudent "gin/handlers/student-handlers/create"
	handlerDeleteStudent "gin/handlers/student-handlers/delete"
	handlerResultStudent "gin/handlers/student-handlers/result"
	handlerResultsStudent "gin/handlers/student-handlers/results"
	handlerUpdateStudent "gin/handlers/student-handlers/update"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

func InitStudentRoutes(db *gorm.DB, route *gin.Engine) {

	/**
	@description All Handler Student
	*/

	createStudentRepository := createStudent.NewRepositoryCreate(db)
	createStudentService := createStudent.NewServiceCreate(createStudentRepository)
	createStudentHandler := handlerCreateStudent.NewHandlerCreateStudent(createStudentService)

	resultsStudentRepository := resultsStudent.NewRepositoryResults(db)
	resultsStudentService := resultsStudent.NewServiceResults(resultsStudentRepository)
	resultsStudentHandler := handlerResultsStudent.NewHandlerResultsStudent(resultsStudentService)

	resultStudentRepository := resultStudent.NewRepositoryResult(db)
	resultStudentService := resultStudent.NewServiceResult(resultStudentRepository)
	resultStudentHandler := handlerResultStudent.NewHandlerResultStudent(resultStudentService)

	deleteStudentRepository := deleteStudent.NewRepositoryDelete(db)
	deleteStudentService := deleteStudent.NewServiceDelete(deleteStudentRepository)
	deleteStudentHandler := handlerDeleteStudent.NewHandlerDeleteStudent(deleteStudentService)

	updateStudentRepository := updateStudent.NewRepositoryUpdate(db)
	updateStudentService := updateStudent.NewServiceUpdate(updateStudentRepository)
	updateStudentHandler := handlerUpdateStudent.NewHandlerUpdateStudent(updateStudentService)

	/**
	@description All Student Route
	*/
	groupRoute := route.Group("/api/v1")
	groupRoute.POST("/student", createStudentHandler.CreateStudentHandler)
	groupRoute.GET("/student", resultsStudentHandler.ResultsStudentHandler)
	groupRoute.GET("/student/:id", resultStudentHandler.ResultStudentHandler)
	groupRoute.DELETE("/student/:id", deleteStudentHandler.DeleteStudentHandler)
	groupRoute.PUT("/student/:id", updateStudentHandler.UpdateStudentHandler)
}
