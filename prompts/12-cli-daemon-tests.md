# 12 Add CLI and daemon tests

## Prompt

Can you create unit tests for the cli and daemon portions of main.go?

## Response

For unit testing CLI applications in Go, you can use the os/exec package to run the application with different command line arguments and then verify the output. Here's an example of how you might do that:

```go
package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestCli(t *testing.T) {
	cmd := exec.Command("go", "run", "main.go", "password")
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedOutput := "0\n" // The password "password" is already strong
	if strings.Compare(string(output), expectedOutput) != 0 {
		t.Fatalf("Expected %v, got %v", expectedOutput, output)
	}
}
```

This test runs go run main.go password and verifies that the output is "0\n", which is the expected result for a strong password.

Testing the daemon part of your application is more challenging because it involves network interactions. One common approach is to start the server in a goroutine and then use an HTTP client to send requests and verify the responses. However, this is more of an integration test than a unit test.

```go
package main

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestDaemon(t *testing.T) {
	go main()

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

	expectedBody := "0\n"
	if strings.Compare(string(body), expectedBody) != 0 {
		t.Fatalf("Expected %v, got %v", expectedBody, body)
	}
}
```

This test starts the server, waits for a second, sends a GET request to http://localhost:8080/check/password, and then verifies that the response is "0\n".

Please note that these are very simple tests. A production-grade application would require more comprehensive testing, including error cases, edge cases, and performance tests. Also, consider that these tests have side effects and depend on external resources, so they are more like integration tests rather than pure unit tests. It might be a good idea to separate these tests from the unit tests and run them in a different test suite.
