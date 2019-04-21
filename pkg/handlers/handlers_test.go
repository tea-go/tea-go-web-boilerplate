package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/tea-go/tea-go-web-boilerplate/pkg/app"
	"github.com/tea-go/tea-go-web-boilerplate/pkg/router"
	"github.com/tea-go/tea-go-web-boilerplate/pkg/testdata"
)

type apiTestCase struct {
	tag                string
	method             string
	url                string
	body               string
	status             int
	expectResponseBody func(body string)
}

func newRouter() router.Routes {
	logger := logrus.New()
	logger.Level = logrus.PanicLevel

	router := router.NewRouter()

	router.Use(
		[]gin.HandlerFunc{
			app.Init(logger), app.Transaction(testdata.DB),
		},
	)

	return router
}

func testAPI(router router.Routes, method, URL, body string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, URL, bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	router.Router().ServeHTTP(res, req)
	return res
}

func runAPITests(t *testing.T, router router.Routes, tests []apiTestCase) {
	for _, test := range tests {
		res := testAPI(router, test.method, test.url, test.body)
		// assert response status
		assert.Equal(t, test.status, res.Code, test.tag)
		// assert response body
		if test.expectResponseBody != nil {
			test.expectResponseBody(res.Body.String())
		}
	}
}

type baseResponseBody struct {
	Status     string `json:"status"`
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

// ListResponseBody define a list api response basic body structure
type ListResponseBody struct {
	baseResponseBody
	Count int `json:"count"`
}

// DetailResponseBody define a detail api response basic body structure
type DetailResponseBody struct {
	baseResponseBody
}
