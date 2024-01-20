package main

import (
	"bytes"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"testing"

	"user-register-api/domain"
)

func TestPutUserParamsNotExist(t *testing.T) {
	expectedStatusCode := http.StatusNotFound
	expectedMessage := "404 page not found"

	resp, err := sendRequest("PUT", endpoint+"/user", nil)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != expectedStatusCode {
		t.Fatalf("Expected status code %v, got %v", expectedStatusCode, resp.StatusCode)
	}

	responseData, _ := ioutil.ReadAll(resp.Body)
	
	responseMessage := string(responseData)
  	if responseMessage != expectedMessage {
		t.Fatalf("Expected message %v, got %v", expectedMessage, responseMessage)
	}
}

func TestPutUserBodyNotExist(t *testing.T) {
	expectedStatusCode := http.StatusBadRequest
	expectedMessage := "Body does not exist"

	resp, err := sendRequest("PUT", endpoint+"/user/"+userID, nil)
	if err != nil {
		t.Fatal(err)
	}
	
	err = verifyExpectedResponse(resp, expectedStatusCode, expectedMessage)
	if err != nil {
		t.Fatal(err)
	}
}

func TestPutUserPasswordNotExist(t *testing.T) {
	expectedStatusCode := http.StatusUnauthorized
	expectedMessage := "Body is not valid"

	requestBody := &domain.User{
		Username: "testname",
	}

	jsonString, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := sendRequest("PUT", endpoint+"/user/"+userID, bytes.NewBuffer(jsonString))
	if err != nil {
		t.Fatal(err)
	}
	
	err = verifyExpectedResponse(resp, expectedStatusCode, expectedMessage)
	if err != nil {
		t.Fatal(err)
	}
}

func TestPutUserUsernameNotExist(t *testing.T) {
	expectedStatusCode := http.StatusUnauthorized
	expectedMessage := "Body is not valid"

	requestBody := &domain.User{
		Password: "testpass",
	}

	jsonString, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := sendRequest("PUT", endpoint+"/user/"+userID, bytes.NewBuffer(jsonString))
	if err != nil {
		t.Fatal(err)
	}
	
	err = verifyExpectedResponse(resp, expectedStatusCode, expectedMessage)
	if err != nil {
		t.Fatal(err)
	}
}

func TestPutUserUserNotFound(t *testing.T) {
	expectedStatusCode := http.StatusNotFound
	expectedMessage := "User not found"

	requestBody := &domain.User{
		Username: "testname",
		Password: "testpass",
	}

	jsonString, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := sendRequest("PUT", endpoint+"/user/"+"not_exist_user", bytes.NewBuffer(jsonString))
	if err != nil {
		t.Fatal(err)
	}
	
	err = verifyExpectedResponse(resp, expectedStatusCode, expectedMessage)
	if err != nil {
		t.Fatal(err)
	}
}

func TestPutUserPasswordNotCorrect(t *testing.T) {
	expectedStatusCode := http.StatusUnauthorized
	expectedMessage := "Password is incorrect"

	requestBody := &domain.User{
		Username: "testname",
		Password: "not_correct_pass",
	}

	jsonString, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := sendRequest("PUT", endpoint+"/user/"+userID, bytes.NewBuffer(jsonString))
	if err != nil {
		t.Fatal(err)
	}
	
	err = verifyExpectedResponse(resp, expectedStatusCode, expectedMessage)
	if err != nil {
		t.Fatal(err)
	}
}

func TestPutUserSuccess(t *testing.T) {
	expectedStatusCode := http.StatusOK
	expectedMessage := "User can be updated"
	expectedUserInfo := domain.User{
		UserID: "testuser",
		Username: "testname_updated",
	}

	requestBody := &domain.User{
		Username: "testname_updated",
		Password: "testpass",
	}

	jsonString, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := sendRequest("PUT", endpoint+"/user/"+userID, bytes.NewBuffer(jsonString))
	if err != nil {
		t.Fatal(err)
	}
	
	err = verifyExpectedResponse(resp, expectedStatusCode, expectedMessage)
	if err != nil {
		t.Fatal(err)
	}

	if response.User != expectedUserInfo {
		t.Fatalf("Expected user %v, got %v", expectedUserInfo, response.User)
	}
}
