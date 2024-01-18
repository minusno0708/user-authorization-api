package main

import (
	"bytes"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"testing"

	"user-register-api/domain"
)

var userID = "testuser"

func TestGetUserParamsNotExist(t *testing.T) {
	expectedStatusCode := http.StatusNotFound
	expectedMessage := "404 page not found"

	req, err := http.NewRequest("GET", endpoint+"/user", nil)
	if err != nil {
		t.Fatal(err)
	}
	client := &http.Client{}

	resp, err := client.Do(req)
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

func TestGetUserBodyNotExist(t *testing.T) {
	expectedStatusCode := http.StatusBadRequest
	expectedMessage := "Body does not exist"

	req, err := http.NewRequest("GET", endpoint+"/user/"+userID, nil)
	if err != nil {
		t.Fatal(err)
	}
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
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
}

func TestGetUserPasswordNotExist(t *testing.T) {
	expectedStatusCode := http.StatusUnauthorized
	expectedMessage := "Body is not valid"

	requestBody := &domain.User{}

	jsonString, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("GET", endpoint+"/user/"+userID, bytes.NewBuffer(jsonString))
	if err != nil {
		t.Fatal(err)
	}
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
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
}

func TestGetUserUserNotFound(t *testing.T) {
	expectedStatusCode := http.StatusNotFound
	expectedMessage := "User not found"

	requestBody := &domain.User{
		Password: "testpass",
	}

	jsonString, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("GET", endpoint+"/user/"+"not_exist_user", bytes.NewBuffer(jsonString))
	if err != nil {
		t.Fatal(err)
	}
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
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
}

func TestGetUserSuccess(t *testing.T) {
	expectedStatusCode := http.StatusOK
	expectedMessage := "User can be acquired"
	expectedUserInfo := domain.User{
		UserID: "testuser",
		Username: "testname",
	}

	requestBody := &domain.User{
		Password: "testpass",
	}

	jsonString, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("GET", endpoint+"/user/"+userID, bytes.NewBuffer(jsonString))
	if err != nil {
		t.Fatal(err)
	}
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
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

	if response.User.UserID != expectedUserInfo.UserID {
		t.Fatalf("Expected user id %v, got %v", expectedUserInfo.UserID, response.User.UserID)
	}
	if response.User.Username != expectedUserInfo.Username {
		t.Fatalf("Expected username %v, got %v", expectedUserInfo.Username, response.User.Username)
	}
}
