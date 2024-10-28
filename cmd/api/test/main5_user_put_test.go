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

	header := setToken(accessToken).ToArray()

	resp, err := sendRequest("PUT", endpoint+"/user", header, nil)
	if err != nil {
		t.Fatal(err)
	}

	_, err = verifyExpectedResponse(resp, expectedStatusCode, expectedMessage)
	if err != nil {
		t.Fatal(err)
	}
}

func TestPutUserUsernameNotExist(t *testing.T) {
	expectedStatusCode := http.StatusUnauthorized
	expectedMessage := "Body is not valid"

	requestBody := requestBody{}

	jsonString, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err)
	}

	header := setToken(accessToken).ToArray()

	resp, err := sendRequest("PUT", endpoint+"/user", header, bytes.NewBuffer(jsonString))
	if err != nil {
		t.Fatal(err)
	}

	_, err = verifyExpectedResponse(resp, expectedStatusCode, expectedMessage)
	if err != nil {
		t.Fatal(err)
	}
}

func TestPutUserTokenNotExist(t *testing.T) {
	expectedStatusCode := http.StatusUnauthorized
	expectedMessage := "Failed to authenticate"

	requestBody := requestBody{
		Username: updatedTestUser.Username,
		Email:    updatedTestUser.Email,
	}

	jsonString, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := sendRequest("PUT", endpoint+"/user", nil, bytes.NewBuffer(jsonString))
	if err != nil {
		t.Fatal(err)
	}

	_, err = verifyExpectedResponse(resp, expectedStatusCode, expectedMessage)
	if err != nil {
		t.Fatal(err)
	}
}

func TestPutUserTokenNotCorrect(t *testing.T) {
	expectedStatusCode := http.StatusUnauthorized
	expectedMessage := "Failed to authenticate"

	requestBody := requestBody{
		Username: updatedTestUser.Username,
		Email:    updatedTestUser.Email,
	}

	jsonString, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err)
	}

	header := setToken("incorrect token").ToArray()

	resp, err := sendRequest("PUT", endpoint+"/user", header, bytes.NewBuffer(jsonString))
	if err != nil {
		t.Fatal(err)
	}

	_, err = verifyExpectedResponse(resp, expectedStatusCode, expectedMessage)
	if err != nil {
		t.Fatal(err)
	}
}

func TestPutUserNameUpdate(t *testing.T) {
	expectedStatusCode := http.StatusOK
	expectedMessage := "User can be updated"

	expectedUser := &domain.User{
		Username: updatedTestUser.Username,
		Email:    testUser.Email,
	}

	requestBody := requestBody{
		Username: updatedTestUser.Username,
		Email:    testUser.Email,
	}

	jsonString, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err)
	}

	header := setToken(accessToken).ToArray()

	resp, err := sendRequest("PUT", endpoint+"/user", header, bytes.NewBuffer(jsonString))
	if err != nil {
		t.Fatal(err)
	}

	_, err = verifyExpectedResponse(resp, expectedStatusCode, expectedMessage)
	if err != nil {
		t.Fatal(err)
	}

	resp, err = sendRequest("GET", endpoint+"/user", header, bytes.NewBuffer(jsonString))
	if err != nil {
		t.Fatal(err)
	}

	response, err := verifyExpectedResponse(resp, expectedStatusCode, "User can be acquired")
	if err != nil {
		t.Fatal(err)
	}

	if response.User != *expectedUser {
		t.Fatal("Username is not updated")
	}
}

func TestPutUserMailUpdate(t *testing.T) {
	expectedStatusCode := http.StatusOK
	expectedMessage := "User can be updated"

	expectedUser := &domain.User{
		Username: updatedTestUser.Username,
		Email:    updatedTestUser.Email,
	}

	requestBody := requestBody{
		Username: updatedTestUser.Username,
		Email:    updatedTestUser.Email,
	}

	jsonString, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err)
	}

	header := setToken(accessToken).ToArray()

	resp, err := sendRequest("PUT", endpoint+"/user", header, bytes.NewBuffer(jsonString))
	if err != nil {
		t.Fatal(err)
	}

	_, err = verifyExpectedResponse(resp, expectedStatusCode, expectedMessage)
	if err != nil {
		t.Fatal(err)
	}

	resp, err = sendRequest("GET", endpoint+"/user", header, bytes.NewBuffer(jsonString))
	if err != nil {
		t.Fatal(err)
	}

	response, err := verifyExpectedResponse(resp, expectedStatusCode, "User can be acquired")
	if err != nil {
		t.Fatal(err)
	}

	if response.User != *expectedUser {
		t.Fatal("Email is not updated")
	}
}

func TestPutUserNameAndMailUpdate(t *testing.T) {
	expectedStatusCode := http.StatusOK
	expectedMessage := "User can be updated"

	expectedUser := &domain.User{
		Username: testUser.Username,
		Email:    testUser.Email,
	}

	requestBody := requestBody{
		Username: testUser.Username,
		Email:    testUser.Email,
	}

	jsonString, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err)
	}

	header := setToken(accessToken).ToArray()

	resp, err := sendRequest("PUT", endpoint+"/user", header, bytes.NewBuffer(jsonString))
	if err != nil {
		t.Fatal(err)
	}

	_, err = verifyExpectedResponse(resp, expectedStatusCode, expectedMessage)
	if err != nil {
		t.Fatal(err)
	}

	resp, err = sendRequest("GET", endpoint+"/user", header, bytes.NewBuffer(jsonString))
	if err != nil {
		t.Fatal(err)
	}

	response, err := verifyExpectedResponse(resp, expectedStatusCode, "User can be acquired")
	if err != nil {
		t.Fatal(err)
	}

	if response.User != *expectedUser {
		t.Fatal("Username And Email is not updated")
	}
}
