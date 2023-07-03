package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetApplicationDetails(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ApplicationDetails)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expectedVersion, err := ioutil.ReadFile("./VERSION")
	if err != nil {
		expectedVersion = []byte("No version set")
	}

	expected := `<!DOCTYPE html>
	<html>
	<body style="background-color:white;">
	<h2 style="color:black;">Deployed Version: ` + string(expectedVersion) + ` </h2>
	<h2 style="color:black;">Environment: default</h2>
	<h3 style="color:black;">List of all environment variables</h3>`

	if !strings.HasPrefix(rr.Body.String(), expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
