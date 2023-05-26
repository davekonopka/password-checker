package main

import (
	"io/ioutil"
	"net/http"
	"os/exec"
	"strings"
	"testing"
	"time"
)

func TestCli(t *testing.T) {
	cmd := exec.Command("go", "run", "main.go", "password")
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedOutput := "0\n" // The password "password" is already strong
	if strings.TrimSpace(string(output)) != expectedOutput {
		t.Fatalf("Expected %v, got %v", expectedOutput, string(output))
	}
}

func TestDaemon(t *testing.T) {
	go func() {
		main()
	}()

	// Give the server a moment to start
	time.Sleep(1 * time.Second)

	resp, err := http.Get("http://localhost:8080/check/password")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedBody := "0"
	if strings.TrimSpace(string(body)) != expectedBody {
		t.Fatalf("Expected %v, got %v", expectedBody, string(body))
	}
}

func TestCheckPasswordStrength(t *testing.T) {
	tests := []struct {
		name     string
		password string
		steps    int
	}{
		{
			name:     "Test 1: Minimum Length",
			password: "a",
			steps:    5,
		},
		{
			name:     "Test 2: Missing Upper Case and Digit",
			password: "aA1",
			steps:    3,
		},
		{
			name:     "Test 3: Already Strong",
			password: "1337C0d3",
			steps:    0,
		},
		{
			name:     "Test 4: Repeating Characters",
			password: "aaaBBB1",
			steps:    2,
		},
		{
			name:     "Test 5: Exceeding Maximum Length",
			password: "abcABC123abcABC123abcABC123abcABC123abcABC123abcABC123",
			steps:    34,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			steps := CheckPasswordStrength(test.password)
			if steps != test.steps {
				t.Errorf("Expected %d steps, but got %d steps", test.steps, steps)
			}
		})
	}
}
