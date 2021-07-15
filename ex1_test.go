package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestVersion(t *testing.T) {
	// Create a HTTP request
	req, err := http.NewRequest(http.MethodGet, "/version", nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()

	version(w, req)
	res := w.Result()
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error getting version %v", err)
	}

	if string(data) != VERSION {
		t.Errorf("Invalid version returned %v", string(data))
	}
}

func TestDuration(t *testing.T) {
	// Create a HTTP request
	req, err := http.NewRequest(http.MethodGet, "/duration", nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()

	duration(w, req)
	res := w.Result()
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error getting duration %v", err)
	}

	// Response should end in 's' - seconds
	if data[len(data)-1] != []byte("s")[0] {
		t.Errorf("Invalid duration returned %v", string(data))
	}
}
