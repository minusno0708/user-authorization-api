package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"user-register-api/domain"
)

func TestCreateBodyNotExist(t *testing.T) {
	expectedStatusCode := http.StatusBadRequest
	expectedMessage := "Body does not exist"

	resp, err := sendRequest("POST", endpoint+"/create", nil)
	if err != nil {
		t.Fatal(err)
	}

	err = verifyExpectedResponse(resp, expectedStatusCode, expectedMessage, nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCreateUserIDNotExist(t *testing.T) {
	expectedStatusCode := http.StatusUnauthorized
	expectedMessage := "Body is not valid"

	requestBody := &domain.User{
		Username: "testuser",
		Password: "testpass",
	}

	jsonString, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := sendRequest("POST", endpoint+"/create", bytes.NewBuffer(jsonString))
	if err != nil {
		t.Fatal(err)
	}

	err = verifyExpectedResponse(resp, expectedStatusCode, expectedMessage, nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCreatePasswordNotExist(t *testing.T) {
	expectedStatusCode := http.StatusUnauthorized
	expectedMessage := "Body is not valid"

	requestBody := &domain.User{
		UserID:   "testuser",
		Username: "testuser",
	}

	jsonString, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := sendRequest("POST", endpoint+"/create", bytes.NewBuffer(jsonString))
	if err != nil {
		t.Fatal(err)
	}

	err = verifyExpectedResponse(resp, expectedStatusCode, expectedMessage, nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCreateSuccessUsernameExist(t *testing.T) {
	expectedStatusCode := http.StatusCreated
	expectedMessage := "User created successfully"

	requestBody := &domain.User{
		UserID:   "testuser",
		Username: "testname",
		Password: "testpass",
	}

	expectedUser := &domain.User{
		UserID:   "testuser",
		Username: "testname",
	}

	jsonString, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := sendRequest("POST", endpoint+"/create", bytes.NewBuffer(jsonString))
	if err != nil {
		t.Fatal(err)
	}

	err = verifyExpectedResponse(resp, expectedStatusCode, expectedMessage, expectedUser)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCreateSuccessUsernameNotExist(t *testing.T) {
	expectedStatusCode := http.StatusCreated
	expectedMessage := "User created successfully"

	requestBody := &domain.User{
		UserID:   "testuser_name_not_exist",
		Password: "testpass",
	}

	expectedUser := &domain.User{
		UserID:   "testuser_name_not_exist",
		Username: "testuser_name_not_exist",
	}

	jsonString, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := sendRequest("POST", endpoint+"/create", bytes.NewBuffer(jsonString))
	if err != nil {
		t.Fatal(err)
	}

	err = verifyExpectedResponse(resp, expectedStatusCode, expectedMessage, expectedUser)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCreateUserConflict(t *testing.T) {
	expectedStatusCode := http.StatusConflict
	expectedMessage := "User already exists"

	requestBody := &domain.User{
		UserID:   "testuser",
		Username: "testname",
		Password: "testpass",
	}

	jsonString, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := sendRequest("POST", endpoint+"/create", bytes.NewBuffer(jsonString))
	if err != nil {
		t.Fatal(err)
	}

	err = verifyExpectedResponse(resp, expectedStatusCode, expectedMessage, nil)
	if err != nil {
		t.Fatal(err)
	}
}
