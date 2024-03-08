package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
	"user-register-api/domain"
)

func TestPutUserBodyNotExist(t *testing.T) {
	expectedStatusCode := http.StatusBadRequest
	expectedMessage := "Body does not exist"

	resp, err := sendRequest("PUT", endpoint+"/user", nil)
	if err != nil {
		t.Fatal(err)
	}

	err = verifyExpectedResponse(resp, expectedStatusCode, expectedMessage, nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestPutUserTokenNotExist(t *testing.T) {
	expectedStatusCode := http.StatusUnauthorized
	expectedMessage := "Body is not valid"

	requestBody := requestBody{
		Username: "testname",
	}

	jsonString, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := sendRequest("PUT", endpoint+"/user", bytes.NewBuffer(jsonString))
	if err != nil {
		t.Fatal(err)
	}

	err = verifyExpectedResponse(resp, expectedStatusCode, expectedMessage, nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestPutUserUsernameNotExist(t *testing.T) {
	expectedStatusCode := http.StatusUnauthorized
	expectedMessage := "Body is not valid"

	requestBody := requestBody{
		TokenString: accessToken,
	}

	jsonString, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := sendRequest("PUT", endpoint+"/user", bytes.NewBuffer(jsonString))
	if err != nil {
		t.Fatal(err)
	}

	err = verifyExpectedResponse(resp, expectedStatusCode, expectedMessage, nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestPutUserTokenNotCorrect(t *testing.T) {
	expectedStatusCode := http.StatusUnauthorized
	expectedMessage := "Failed to authenticate"

	requestBody := requestBody{
		Username:    "testname",
		TokenString: "incorrect token string",
	}

	jsonString, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := sendRequest("PUT", endpoint+"/user", bytes.NewBuffer(jsonString))
	if err != nil {
		t.Fatal(err)
	}

	err = verifyExpectedResponse(resp, expectedStatusCode, expectedMessage, nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestPutUserSuccess(t *testing.T) {
	expectedStatusCode := http.StatusOK
	expectedMessage := "User can be updated"

	expectedUser := &domain.User{
		UserID:   "testuser",
		Username: "testname_updated",
	}

	requestBody := requestBody{
		Username:    "testname_updated",
		TokenString: accessToken,
	}

	jsonString, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := sendRequest("PUT", endpoint+"/user", bytes.NewBuffer(jsonString))
	if err != nil {
		t.Fatal(err)
	}

	err = verifyExpectedResponse(resp, expectedStatusCode, expectedMessage, nil)
	if err != nil {
		t.Fatal(err)
	}

	resp, err = sendRequest("GET", endpoint+"/user", bytes.NewBuffer(jsonString))
	if err != nil {
		t.Fatal(err)
	}

	err = verifyExpectedResponse(resp, expectedStatusCode, "User can be acquired", expectedUser)
	if err != nil {
		t.Fatal(err)
	}
}
