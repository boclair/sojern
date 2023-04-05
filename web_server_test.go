package main

import (
	"io"
	"net/http"
	"strings"
	"testing"
	"time"
)

// These tests cannot run in parallel.  They all depend on port 8080.

func TestPingSuccess(t *testing.T) {
	fileChan := make(chan FileExists)
	srv := StartServer(fileChan, ":8080")
	serverWait()
	defer srv.Close()
	
	fileChan <- FILE_EXISTS

	resp, err := http.Get("http://localhost:8080/ping")
	if err != nil {
		t.Error("Error during GET: ", err)
	}

	if resp.StatusCode != 200 {
		t.Error("Unexpected status code:", resp.StatusCode)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error("Error reading BODY:", err)
	}

	if string(body) != "OK" {
		t.Error("Unexpected response:", body)
	}
}

func TestPingSuccessAndFail(t *testing.T) {
	fileChan := make(chan FileExists)
	srv := StartServer(fileChan, ":8081")
	serverWait()
	defer srv.Close()

	fileChan <- FILE_EXISTS
	resp, _ := http.Get("http://localhost:8081/ping")
	if resp.StatusCode != 200 {
		t.Error("Unexpected status code:", resp.StatusCode)
	}

	fileChan <- FILE_ABSENT
	resp, _ = http.Get("http://localhost:8081/ping")
	if resp.StatusCode != 503 {
		t.Error("Unexpected status code:", resp.StatusCode)
	}

	fileChan <- FILE_EXISTS
	resp, _ = http.Get("http://localhost:8081/ping")
	if resp.StatusCode != 200 {
		t.Error("Unexpected status code:", resp.StatusCode)
	}
}

func TestImg(t *testing.T) {
	fileChan := make(chan FileExists)
	srv := StartServer(fileChan, ":8082")
	serverWait()
	defer srv.Close()
	
	resp, err := http.Get("http://localhost:8082/img")
	if err != nil {
		t.Error("Error during GET: ", err)
	}

	if resp.StatusCode != 200 {
		t.Error("Unexpected status code:", resp.StatusCode)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error("Error reading BODY:", err)
	}

	if !strings.HasPrefix(string(body), "GIF89a") {
		t.Error("Unexpected response:", string(body))
	}
}


// Wait for the server to start
func serverWait() {
	time.Sleep(500 * time.Millisecond)
}
