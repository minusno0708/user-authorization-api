package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"user-register-api/domain"
)

func TestSignupBodyNotExist(t *testing.T) {
	expectedStatusCode := http.StatusBadRequest
	expectedMessage := "Body does not exist"

	resp, err := sendRequest("POST", endpoint+"/signup", nil)
	if err != nil {
		t.Fatal(err)
	}

	err = verifyExpectedResponse(resp, expectedStatusCode, expectedMessage, nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSignupUserIDNotExist(t *testing.T) {
	expectedStatusCode := http.StatusUnauthorized
	expectedMessage := "Body is not valid"

	requestBody := requestBody{
		Username: "testuser",
		Password: "testpass",
	}

	jsonString, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := sendRequest("POST", endpoint+"/signup", bytes.NewBuffer(jsonString))
	if err != nil {
		t.Fatal(err)
	}

	err = verifyExpectedResponse(resp, expectedStatusCode, expectedMessage, nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSignupPasswordNotExist(t *testing.T) {
	expectedStatusCode := http.StatusUnauthorized
	expectedMessage := "Body is not valid"

	requestBody := requestBody{
		UserID:   "testuser",
		Username: "testuser",
	}

	jsonString, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := sendRequest("POST", endpoint+"/signup", bytes.NewBuffer(jsonString))
	if err != nil {
		t.Fatal(err)
	}

	err = verifyExpectedResponse(resp, expectedStatusCode, expectedMessage, nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSignupSuccessUsernameExist(t *testing.T) {
	expectedStatusCode := http.StatusCreated
	expectedMessage := "User created successfully"

	requestBody := requestBody{
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

	resp, err := sendRequest("POST", endpoint+"/signup", bytes.NewBuffer(jsonString))
	if err != nil {
		t.Fatal(err)
	}

	err = verifyExpectedResponse(resp, expectedStatusCode, expectedMessage, expectedUser)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSignupSuccessUsernameNotExist(t *testing.T) {
	expectedStatusCode := http.StatusCreated
	expectedMessage := "User created successfully"

	requestBody := requestBody{
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

	resp, err := sendRequest("POST", endpoint+"/signup", bytes.NewBuffer(jsonString))
	if err != nil {
		t.Fatal(err)
	}

	err = verifyExpectedResponse(resp, expectedStatusCode, expectedMessage, expectedUser)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSignupUserConflict(t *testing.T) {
	expectedStatusCode := http.StatusConflict
	expectedMessage := "User already exists"

	requestBody := requestBody{
		UserID:   "testuser",
		Username: "testname",
		Password: "testpass",
	}

	jsonString, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := sendRequest("POST", endpoint+"/signup", bytes.NewBuffer(jsonString))
	if err != nil {
		t.Fatal(err)
	}

	err = verifyExpectedResponse(resp, expectedStatusCode, expectedMessage, nil)
	if err != nil {
		t.Fatal(err)
	}
}
