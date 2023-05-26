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
