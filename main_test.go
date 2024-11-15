package main

import (
	util "gin/utils"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"

	"github.com/restuwahyu13/go-supertest/supertest"
	. "github.com/smartystreets/goconvey/convey"
)

var router = SetupRouter()

// var accessToken string
// var studentId interface{}

func TestLoginHandler(t *testing.T) {

	Convey("Add Data in Tests", t, func() {

		Convey("Data Insert In Test Table", func() {
			payload := gin.H{
				"Name": "Amitsss",
			}

			test := supertest.NewSuperTest(router, t)

			test.Post("/api/v1/test")
			test.Send(payload)
			test.Set("Content-Type", "application/json")
			test.End(func(req *http.Request, rr *httptest.ResponseRecorder) {

				response := util.Parse(rr.Body.Bytes())
				t.Log(response)

				assert.Equal(t, http.StatusCreated, rr.Code)
				assert.Equal(t, http.MethodPost, req.Method)
				assert.Equal(t, "Create new Test successfully", response.Message)
			})
		})
	})
}
