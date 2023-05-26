# 15 Revert, remove CLI, daemon tests

After a few attempts at getting CLI and daemon tests without success I decided to bail out and revert those changes.

## Prompt

I wasn't able to get the TestCli and TestDaemon tests succeeding. I removed the tests and reverted main.go to the way it was before our last two changes trying to add and fix the tests.

## Response

I understand your frustration and I apologize for the issues you've been facing. The primary challenge here is that we're trying to write unit tests for the CLI and daemon parts of the application which is more complex due to the nature of these interfaces, especially when trying to test them in the same process.

Here's how you can proceed:

Integration Tests: Instead of trying to test these parts in a unit test, consider writing integration tests. These tests would involve actually running the command as a separate process and checking the output, and starting the server as a separate process and sending HTTP requests.

Refactor to Make Code More Testable: This involves abstracting away the command execution and HTTP handling parts to separate functions that can be more easily tested. The CLI part would involve refactoring the code to separate the logic of argument parsing and handling into a function that takes in the arguments and returns the output or error, which can be easily unit tested.

As for the HTTP server part, the handler function can be tested by creating a request and a ResponseRecorder (from net/http/httptest), and then checking the response. But testing the server start and stop part is more tricky, as it involves networking and is best left for integration tests.

However, if you'd like to proceed with integration tests, I can provide assistance for that as well. Here's a simple example of how you could test the CLI part:

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

	expectedOutput := "0\n"
	if strings.TrimSpace(string(output)) != expectedOutput {
		t.Fatalf("Expected %v, got %v", expectedOutput, string(output))
	}
}
```

You can replace "password" with the password you want to test, and "0\n" with the expected output.

Please note that this kind of test is best used for integration tests and not unit tests, and should be used sparingly as these kinds of tests are usually slower and can be affected by external factors (like the actual command-line environment). This is just an example to give you an idea, and the actual tests should be more comprehensive and robust.

Let me know how you'd like to proceed or if you need help with other aspects of your application.
