package main

import (
	"net/http"
	"testing"

	"user-register-api/domain"
)

func TestGetUserTokenNotExist(t *testing.T) {
	expectedStatusCode := http.StatusUnauthorized
	expectedMessage := "Failed to authenticate"

	resp, err := sendRequest("GET", endpoint+"/user", nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	_, err = verifyExpectedResponse(resp, expectedStatusCode, expectedMessage)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetUserTokenNotCorrect(t *testing.T) {
	expectedStatusCode := http.StatusUnauthorized
	expectedMessage := "Failed to authenticate"

	header := setToken("incorrect token").ToArray()

	resp, err := sendRequest("GET", endpoint+"/user", header, nil)
	if err != nil {
		t.Fatal(err)
	}

	_, err = verifyExpectedResponse(resp, expectedStatusCode, expectedMessage)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetUserSuccess(t *testing.T) {
	expectedStatusCode := http.StatusOK
	expectedMessage := "User can be acquired"
	expectedUser := &domain.User{
		UserID:   "testuser",
		Username: "testname",
	}

	header := setToken(accessToken).ToArray()

	resp, err := sendRequest("GET", endpoint+"/user", header, nil)
	if err != nil {
		t.Fatal(err)
	}

	response, err := verifyExpectedResponse(resp, expectedStatusCode, expectedMessage)
	if err != nil {
		t.Fatal(err)
	}

	if response.User != *expectedUser {
		t.Fatal("User does not match")
	}
}
