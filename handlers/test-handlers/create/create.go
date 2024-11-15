package test

import (
	"fmt"
	createTest "gin/controllers/test-controllers/create"
	util "gin/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	gpc "github.com/restuwahyu13/go-playground-converter"
)

type handler struct {
	service createTest.Service
}

func NewHandlerCreateTest(service createTest.Service) *handler {
	return &handler{service: service}
}

func (h *handler) CreateTestHandler(ctx *gin.Context) {
	fmt.Println("call2")

	var input createTest.InputCreateTest
	ctx.ShouldBindJSON(&input)

	config := gpc.ErrorConfig{
		Options: []gpc.ErrorMetaConfig{
			gpc.ErrorMetaConfig{
				Tag:     "required",
				Field:   "Name",
				Message: "name is required on body",
			},
		},
	}
	errResponse, errCount := util.GoValidator(&input, config.Options)

	if errCount > 0 {
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, errResponse)
		return
	}
	_, errCreateTest := h.service.CreateTestService(&input)
	fmt.Println("errCreateTest", errCreateTest)
	switch errCreateTest {

	case "CREATE_TESTS_CONFLICT_409":
		util.APIResponse(ctx, "Name TEST already exist", http.StatusConflict, http.MethodPost, nil)
		return

	case "CREATE_TESTS_FAILED_403":
		util.APIResponse(ctx, "Create new Test failed", http.StatusForbidden, http.MethodPost, nil)
		return

	default:
		util.APIResponse(ctx, "Create new Test successfully", http.StatusCreated, http.MethodPost, nil)
	}
}
