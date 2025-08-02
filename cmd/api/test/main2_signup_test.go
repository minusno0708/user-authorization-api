package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func TestSignupBodyNotExist(t *testing.T) {
	expectedStatusCode := http.StatusBadRequest
	expectedMessage := "Body does not exist"

	resp, err := sendRequest("POST", endpoint+"/signup", nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	_, err = verifyExpectedResponse(resp, expectedStatusCode, expectedMessage)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSignupUsernameNotExist(t *testing.T) {
	expectedStatusCode := http.StatusUnauthorized
	expectedMessage := "Body is not valid"

	requestBody := requestBody{
		Email:    testUser.Email,
		Password: testUser.Password,
	}

	jsonString, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := sendRequest("POST", endpoint+"/signup", nil, bytes.NewBuffer(jsonString))
	if err != nil {
		t.Fatal(err)
	}

	_, err = verifyExpectedResponse(resp, expectedStatusCode, expectedMessage)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSignupEmailNotExist(t *testing.T) {
	expectedStatusCode := http.StatusUnauthorized
	expectedMessage := "Body is not valid"

	requestBody := requestBody{
		Username: testUser.Username,
		Password: testUser.Password,
	}

	jsonString, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := sendRequest("POST", endpoint+"/signup", nil, bytes.NewBuffer(jsonString))
	if err != nil {
		t.Fatal(err)
	}

	_, err = verifyExpectedResponse(resp, expectedStatusCode, expectedMessage)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSignupPasswordNotExist(t *testing.T) {
	expectedStatusCode := http.StatusUnauthorized
	expectedMessage := "Body is not valid"

	requestBody := requestBody{
		Username: testUser.Username,
		Email:    testUser.Email,
	}

	jsonString, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := sendRequest("POST", endpoint+"/signup", nil, bytes.NewBuffer(jsonString))
	if err != nil {
		t.Fatal(err)
	}

	_, err = verifyExpectedResponse(resp, expectedStatusCode, expectedMessage)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSignupSuccess(t *testing.T) {
	expectedStatusCode := http.StatusCreated
	expectedMessage := "User created successfully"

	requestBody := requestBody{
		Username: testUser.Username,
		Email:    testUser.Email,
		Password: testUser.Password,
	}

	jsonString, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := sendRequest("POST", endpoint+"/signup", nil, bytes.NewBuffer(jsonString))
	if err != nil {
		t.Fatal(err)
	}

	_, err = verifyExpectedResponse(resp, expectedStatusCode, expectedMessage)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSignupUserConflict(t *testing.T) {
	expectedStatusCode := http.StatusConflict
	expectedMessage := "User already exists"

	requestBody := requestBody{
		Username: testUser.Username,
		Email:    testUser.Email,
		Password: testUser.Password,
	}

	jsonString, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := sendRequest("POST", endpoint+"/signup", nil, bytes.NewBuffer(jsonString))
	if err != nil {
		t.Fatal(err)
	}

	_, err = verifyExpectedResponse(resp, expectedStatusCode, expectedMessage)
	if err != nil {
		t.Fatal(err)
	}
}
