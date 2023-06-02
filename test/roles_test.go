package test

import (
	"encoding/json"
	"fmt"
	"go/ems/app"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	db := app.DbConnection()
	validate := validator.New()
	router := app.Router(db, validate)
	return router
}

func TestSuccessCreateRoles(t *testing.T) {
	router := setupRouter()

	name := "admin"
	payload := fmt.Sprintf(`{ "name": "%s" }`, name)

	requestBody := strings.NewReader(payload)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/roles", requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	fmt.Println(responseBody["code"])

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
}

func TestFailCreateRoles(t *testing.T) {
	router := setupRouter()

	name := "admin"
	payload := fmt.Sprintf(`{ "name": "%s" }`, name)

	requestBody := strings.NewReader(payload)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/roles", requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	fmt.Println(responseBody["code"])

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
}

func TestSuccessUpdateRoles(t *testing.T) {
	router := setupRouter()

	id := 18
	name := "admin"

	payload := fmt.Sprintf(`{ "id": %d "name": "%s" }`, id, name)

	requestBody := strings.NewReader(payload)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:8080/api/v1/roles", requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
}
