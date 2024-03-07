package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
	"user-register-api/domain"
)

var accessToken string

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

func TestSigninUserNotExist(t *testing.T) {
	expectedStatusCode := http.StatusUnauthorized
	expectedMessage := "User ID or password is incorrect"

	requestBody := &domain.User{
		UserID:   "testuser_not_exist",
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

func TestSigninPasswordNotCorrect(t *testing.T) {
	expectedStatusCode := http.StatusUnauthorized
	expectedMessage := "User ID or password is incorrect"

	requestBody := &domain.User{
		UserID:   "testuser",
		Password: "testpass_not_correct",
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

func TestSigninSuccess(t *testing.T) {
	expectedStatusCode := http.StatusCreated
	expectedMessage := "Token can be acquired"

	requestBody := &domain.User{
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

	var response struct {
		Message     string `json:"message"`
		TokenString string `json:"token"`
	}

	if resp.StatusCode != expectedStatusCode {
		t.Fatalf("Expected status code %v, got %v", expectedStatusCode, resp.StatusCode)
	}

	responseData, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(responseData, &response)
	if err != nil {
		t.Fatal(err)
	}

	if response.Message != expectedMessage {
		t.Fatalf("Expected message %v, got %v", expectedMessage, response.Message)
	}

	if response.TokenString == "" {
		t.Fatal("Token is empty")
	}

	accessToken = response.TokenString
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

	requestBody := &domain.Token{}

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

func TestSignoutIncorrectToken(t *testing.T) {
	expectedStatusCode := http.StatusUnauthorized
	expectedMessage := "Failed to authenticate"

	requestBody := &domain.Token{
		TokenString: "incorrect_token",
	}
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

func TestSignoutSuccess(t *testing.T) {
	expectedStatusCode := http.StatusOK
	expectedMessage := "Token can be deleted"

	requestBody := &domain.Token{
		TokenString: accessToken,
	}

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

func TestGetAccessToken(t *testing.T) {
	requestBody := &domain.User{
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

	var response struct {
		Message     string `json:"message"`
		TokenString string `json:"token"`
	}

	responseData, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(responseData, &response)
	if err != nil {
		t.Fatal(err)
	}

	if response.TokenString == "" {
		t.Fatal("Token gets failed")
	}

	accessToken = response.TokenString
}
