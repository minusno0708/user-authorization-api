package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func TestDeleteUserBodyNotExist(t *testing.T) {
	expectedStatusCode := http.StatusBadRequest
	expectedMessage := "Body does not exist"

	resp, err := sendRequest("DELETE", endpoint+"/user", nil)
	if err != nil {
		t.Fatal(err)
	}

	_, err = verifyExpectedResponse(resp, expectedStatusCode, expectedMessage)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDeleteUserTokenNotExist(t *testing.T) {
	expectedStatusCode := http.StatusUnauthorized
	expectedMessage := "Body is not valid"

	requestBody := requestBody{}

	jsonString, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := sendRequest("DELETE", endpoint+"/user", bytes.NewBuffer(jsonString))
	if err != nil {
		t.Fatal(err)
	}

	_, err = verifyExpectedResponse(resp, expectedStatusCode, expectedMessage)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDeleteUserTokenNotCorrect(t *testing.T) {
	expectedStatusCode := http.StatusUnauthorized
	expectedMessage := "Failed to authenticate"

	requestBody := requestBody{
		TokenString: "incorrect token string",
	}

	jsonString, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := sendRequest("DELETE", endpoint+"/user", bytes.NewBuffer(jsonString))
	if err != nil {
		t.Fatal(err)
	}

	_, err = verifyExpectedResponse(resp, expectedStatusCode, expectedMessage)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDeleteUserSuccess(t *testing.T) {
	expectedStatusCode := http.StatusOK
	expectedMessage := "User can be deleted"

	requestBody := requestBody{
		TokenString: accessToken,
	}

	jsonString, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := sendRequest("DELETE", endpoint+"/user", bytes.NewBuffer(jsonString))
	if err != nil {
		t.Fatal(err)
	}

	_, err = verifyExpectedResponse(resp, expectedStatusCode, expectedMessage)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDeleteUserIsUserNotExist(t *testing.T) {
	expectedStatusCode := http.StatusUnauthorized
	expectedMessage := "Failed to authenticate"

	requestBody := requestBody{
		TokenString: accessToken,
	}

	jsonString, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := sendRequest("DELETE", endpoint+"/user", bytes.NewBuffer(jsonString))
	if err != nil {
		t.Fatal(err)
	}

	_, err = verifyExpectedResponse(resp, expectedStatusCode, expectedMessage)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUserCanBeDeleted(t *testing.T) {
	expectedStatusCode := http.StatusUnauthorized
	expectedMessage := "User ID or password is incorrect"

	requestBody := requestBody{
		UserID:   "testuser",
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

	_, err = verifyExpectedResponse(resp, expectedStatusCode, expectedMessage)
	if err != nil {
		t.Fatal(err)
	}
}
