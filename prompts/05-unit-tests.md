# 05 Unit Tests

## Prompt

Could you create unit tests for the go code so far?

## Response

Sure, unit tests in Go are typically located in the same package as the code under test and are placed in a separate file with a _test.go suffix. Let's create a file main_test.go that contains tests for the CheckPasswordStrength function:

```go
package main

import "testing"

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
            steps:    36,
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
```

This code defines a series of test cases, each with a name, an input password, and the expected number of steps to make the password strong. It runs each test case using the t.Run function, and if the result of CheckPasswordStrength doesn't match the expected steps, it reports an error with t.Errorf.

To run the tests, you can use the go test command in your terminal:

```bash
go test
```

This command will run all test cases and report any failures. If all tests pass, it will just print ok.
