package test

import (
	"belajar-golang-restful-api/app"
	"belajar-golang-restful-api/controller"
	"belajar-golang-restful-api/helper"
	"belajar-golang-restful-api/middleware"
	"belajar-golang-restful-api/model/domain"
	"belajar-golang-restful-api/model/web"
	"belajar-golang-restful-api/repository"
	"belajar-golang-restful-api/service"
	"bytes"
	"context"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-playground/validator"
	"github.com/stretchr/testify/assert"
)

func dbConnection() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/belajar_golang_restful_api_test")
	helper.UserPanicIfError(err)

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(1)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(5 * time.Minute)

	return db
}

func userTruncate(db *sql.DB) {
	db.Exec("truncate user")
}

func userInsert(db *sql.DB) {
	tx, err := db.Begin()
	helper.UserPanicIfError(err)
	repository := repository.NewUserRepository()

	repository.Create(context.Background(), tx, domain.User{
		Email:    "zayn.malik@gmail.com",
		Username: "zaynmalik192",
		Password: "secret!@#",
		FullName: "Zayn Malik",
	})
	helper.CommitOrRollback(tx)
}

func userInsert2(db *sql.DB) {
	tx, err := db.Begin()
	helper.UserPanicIfError(err)
	repository := repository.NewUserRepository()

	repository.Create(context.Background(), tx, domain.User{
		Email:    "messi@psg.com",
		Username: "messi",
		Password: "secret!@#",
		FullName: "Lionel Messi",
	})
	helper.CommitOrRollback(tx)
}

func UserRestfulAPI(db *sql.DB) http.Handler {

	validate := validator.New()

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(db, userRepository, validate)
	userController := controller.NewUserController(*validate, userService)
	router := app.UserNewRouter(userController)
	middleware := middleware.NewUserMiddleware(router)
	return middleware
}

func TestUserGetDataSuccess(t *testing.T) {
	db := dbConnection()
	userTruncate(db)
	userInsert(db)
	userInsert2(db)
	router := UserRestfulAPI(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/users", nil)
	authKey := base64.StdEncoding.EncodeToString([]byte("noval:secret123"))
	request.Header.Add("Authorization", fmt.Sprintf("Basic %s", authKey))
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, int(response.StatusCode))

	var responseBody map[string]interface{}
	body, _ := io.ReadAll(response.Body)
	json.Unmarshal(body, &responseBody)
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])

	responseBodyArr := responseBody["data"].([]interface{})
	responseUser1 := responseBodyArr[0].(map[string]interface{})
	responseUser2 := responseBodyArr[1].(map[string]interface{})

	assert.Equal(t, "zayn.malik@gmail.com", responseUser1["email"])
	assert.Equal(t, "zaynmalik192", responseUser1["username"])
	assert.Equal(t, "Zayn Malik", responseUser1["full_name"])

	assert.Equal(t, "messi@psg.com", responseUser2["email"])
	assert.Equal(t, "messi", responseUser2["username"])
	assert.Equal(t, "Lionel Messi", responseUser2["full_name"])
}

func TestUserGetDataByIDSuccess(t *testing.T) {
	db := dbConnection()
	userTruncate(db)
	userInsert(db)
	router := UserRestfulAPI(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/users/1", nil)
	authKey := base64.StdEncoding.EncodeToString([]byte("noval:secret123"))
	request.Header.Add("Authorization", fmt.Sprintf("Basic %s", authKey))
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, int(response.StatusCode))

	var responseBody map[string]interface{}
	body, _ := io.ReadAll(recorder.Body)
	json.Unmarshal(body, &responseBody)
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, "zayn.malik@gmail.com", responseBody["data"].(map[string]interface{})["email"])
	assert.Equal(t, "zaynmalik192", responseBody["data"].(map[string]interface{})["username"])
	assert.Equal(t, "Zayn Malik", responseBody["data"].(map[string]interface{})["full_name"])
}

func TestUserGetDataByIDFailed(t *testing.T) {
	db := dbConnection()
	userTruncate(db)
	router := UserRestfulAPI(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/users/1", nil)
	authKey := base64.StdEncoding.EncodeToString([]byte("noval:secret123"))
	request.Header.Add("Authorization", fmt.Sprintf("Basic %s", authKey))
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, int(response.StatusCode))

	var responseBody map[string]interface{}
	body, _ := io.ReadAll(recorder.Body)
	json.Unmarshal(body, &responseBody)
	assert.Equal(t, 404, int(responseBody["code"].(float64)))
	assert.Equal(t, "NOT FOUND", responseBody["status"])
}

func TestUserCreateSuccess(t *testing.T) {
	db := dbConnection()
	userTruncate(db)
	router := UserRestfulAPI(db)

	user := web.UserCreateRequest{
		Email:    "cr7@real.madrid.com",
		Username: "cristiano.ronaldo",
		Password: "secret123",
		FullName: "Cristiano Ronaldo",
	}
	requestBodyByte, _ := json.Marshal(user)
	requestBody := io.NopCloser(bytes.NewBuffer(requestBodyByte))
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/users", requestBody)
	authKey := base64.StdEncoding.EncodeToString([]byte("noval:secret123"))
	request.Header.Add("Authorization", fmt.Sprintf("Basic %s", authKey))
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	var responseBody map[string]interface{}
	body, _ := io.ReadAll(response.Body)
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, "cr7@real.madrid.com", responseBody["data"].(map[string]interface{})["email"])
	assert.Equal(t, "cristiano.ronaldo", responseBody["data"].(map[string]interface{})["username"])
	assert.Equal(t, "Cristiano Ronaldo", responseBody["data"].(map[string]interface{})["full_name"])
}

func TestUserCreateFailed(t *testing.T) {
	db := dbConnection()
	userTruncate(db)
	router := UserRestfulAPI(db)

	user := web.UserCreateRequest{}
	requestBodyByte, _ := json.Marshal(user)
	requestBody := io.NopCloser(bytes.NewBuffer(requestBodyByte))
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/users", requestBody)
	authKey := base64.StdEncoding.EncodeToString([]byte("noval:secret123"))
	request.Header.Add("Authorization", fmt.Sprintf("Basic %s", authKey))
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	var responseBody map[string]interface{}
	body, _ := io.ReadAll(response.Body)
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responseBody["status"])
}

func TestUserUpdateSuccess(t *testing.T) {
	db := dbConnection()
	userTruncate(db)
	userInsert(db)

	router := UserRestfulAPI(db)
	user := web.UserUpdateRequest{
		Email:    "noval@gmail.com",
		Username: "novalwardhana",
		Password: "secret123",
		FullName: "Noval",
	}
	requestBodyByte, _ := json.Marshal(user)
	requestBody := io.NopCloser(bytes.NewBuffer(requestBodyByte))
	request := httptest.NewRequest(http.MethodPut, "http://localhost:8080/api/users/1", requestBody)
	authKey := base64.StdEncoding.EncodeToString([]byte("noval:secret123"))
	request.Header.Add("Authorization", fmt.Sprintf("Basic %s", authKey))
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, int(response.StatusCode))

	var responseBody map[string]interface{}
	body, _ := io.ReadAll(response.Body)
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, "noval@gmail.com", responseBody["data"].(map[string]interface{})["email"])
	assert.Equal(t, "novalwardhana", responseBody["data"].(map[string]interface{})["username"])
	assert.Equal(t, "Noval", responseBody["data"].(map[string]interface{})["full_name"])
}

func TestUserUpdateFailed(t *testing.T) {
	db := dbConnection()
	userTruncate(db)

	router := UserRestfulAPI(db)
	user := web.UserUpdateRequest{}
	requestBodyByte, _ := json.Marshal(user)
	requestBody := io.NopCloser(bytes.NewBuffer(requestBodyByte))
	request := httptest.NewRequest(http.MethodPut, "http://localhost:8080/api/users/1", requestBody)
	authKey := base64.StdEncoding.EncodeToString([]byte("noval:secret123"))
	request.Header.Add("Authorization", fmt.Sprintf("Basic %s", authKey))
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	var responseBody map[string]interface{}
	body, _ := io.ReadAll(response.Body)
	json.Unmarshal(body, &responseBody)
	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responseBody["status"])
}

func TestUserDeleteSuccess(t *testing.T) {
	db := dbConnection()
	userTruncate(db)
	userInsert(db)

	router := UserRestfulAPI(db)
	request := httptest.NewRequest(http.MethodDelete, "http://localhost:8080/api/users/1", nil)
	authKey := base64.StdEncoding.EncodeToString([]byte("noval:secret123"))
	request.Header.Add("Authorization", fmt.Sprintf("Basic %s", authKey))
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	var responseBody map[string]interface{}
	body, _ := io.ReadAll(response.Body)
	json.Unmarshal(body, &responseBody)
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
}

func TestUserDeleteFailed(t *testing.T) {
	db := dbConnection()
	userTruncate(db)

	router := UserRestfulAPI(db)
	request := httptest.NewRequest(http.MethodDelete, "http://localhost:8080/api/users/1", nil)
	authKey := base64.StdEncoding.EncodeToString([]byte("noval:secret123"))
	request.Header.Add("Authorization", fmt.Sprintf("Basic %s", authKey))
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)

	var responseBody map[string]interface{}
	body, _ := io.ReadAll(response.Body)
	json.Unmarshal(body, &responseBody)
	assert.Equal(t, 404, int(responseBody["code"].(float64)))
	assert.Equal(t, "NOT FOUND", responseBody["status"])
}

func TestUserNouAuthorized(t *testing.T) {
	db := dbConnection()
	userTruncate(db)
	userInsert(db)

	router := UserRestfulAPI(db)
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/users", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 401, response.StatusCode)

	var responseBody map[string]interface{}
	body, _ := io.ReadAll(response.Body)
	json.Unmarshal(body, &responseBody)
	assert.Equal(t, 401, int(responseBody["code"].(float64)))
	assert.Equal(t, "UNAUTHORIZED", responseBody["status"])
}
