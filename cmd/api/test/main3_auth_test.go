package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
	"user-register-api/domain"
)

func TestSigninBodyNotExist(t *testing.T) {
	expectedStatusCode := http.StatusBadRequest
	expectedMessage := "Body does not exist"

	resp, err := sendRequest("POST", endpoint+"/signin", nil)
	if err != nil {
		t.Fatal(err)
	}

	err = verifyExpectedResponse(resp, expectedStatusCode, expectedMessage, nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSigninUserIDNotExist(t *testing.T) {
	expectedStatusCode := http.StatusUnauthorized
	expectedMessage := "Body is not valid"

	requestBody := &domain.User{
		Password: "testpass",
	}

	jsonString, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := sendRequest("POST", endpoint+"/signin", bytes.NewBuffer(jsonString))
	if err != nil {
		t.Fatal(err)
	}

	err = verifyExpectedResponse(resp, expectedStatusCode, expectedMessage, nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSigninPasswordNotExist(t *testing.T) {
	expectedStatusCode := http.StatusUnauthorized
	expectedMessage := "Body is not valid"

	requestBody := &domain.User{
		UserID: "testuser",
	}

	jsonString, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := sendRequest("POST", endpoint+"/signin", bytes.NewBuffer(jsonString))
	if err != nil {
		t.Fatal(err)
	}

	err = verifyExpectedResponse(resp, expectedStatusCode, expectedMessage, nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSignoutBodyNotExist(t *testing.T) {
	expectedStatusCode := http.StatusBadRequest
	expectedMessage := "Body does not exist"

	resp, err := sendRequest("DELETE", endpoint+"/signout", nil)
	if err != nil {
		t.Fatal(err)
	}

	err = verifyExpectedResponse(resp, expectedStatusCode, expectedMessage, nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSignoutTokenNotExist(t *testing.T) {
	expectedStatusCode := http.StatusUnauthorized
	expectedMessage := "Body is not valid"

	requestBody := &domain.User{}

	jsonString, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := sendRequest("DELETE", endpoint+"/signout", bytes.NewBuffer(jsonString))
	if err != nil {
		t.Fatal(err)
	}

	err = verifyExpectedResponse(resp, expectedStatusCode, expectedMessage, nil)
	if err != nil {
		t.Fatal(err)
	}
}
