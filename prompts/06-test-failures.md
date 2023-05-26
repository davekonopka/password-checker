# 06 Test Failures

## Prompt

When I run the unit tests I'm getting failures. Can you check the test output and the code you provided and give me direction on how to fix things?

Here is the test output:

```
--- FAIL: TestCheckPasswordStrength (0.00s)
    --- FAIL: TestCheckPasswordStrength/Test_1:_Minimum_Length (0.00s)
        main_test.go:42: Expected 5 steps, but got 7 steps
    --- FAIL: TestCheckPasswordStrength/Test_5:_Exceeding_Maximum_Length (0.00s)
        main_test.go:42: Expected 36 steps, but got 34 steps
FAIL
exit status 1
FAIL	github.com/davekonopka/password-checker	0.186s
```
## Response

Looking at the output, it seems like the failures are related to password length checking logic. Specifically, the failures happen when the password length is less than 6 or more than 20.

The problem appears to be that while checking for password length, we also need to check whether the increment we do for missing character types (lowercase, uppercase, and digit) have already been accounted for in our password length checks.

Here's the updated CheckPasswordStrength function:

```go
package main

import (
    "unicode"
)

func CheckPasswordStrength(password string) int {
    steps := 0

    var uppercase, lowercase, digit bool

    // convert string to rune array for easy manipulation
    runes := []rune(password)

    // Check length
    if len(runes) < 6 {
        steps += 6 - len(runes)
    } else if len(runes) > 20 {
        steps += len(runes) - 20
    }

    // Iterate over password
    for i, char := range runes {
        // Check for required character types
        switch {
        case unicode.IsLower(char):
            lowercase = true
        case unicode.IsUpper(char):
            uppercase = true
        case unicode.IsDigit(char):
            digit = true
        }

        // Check for three repeating characters
        if i >= 2 && runes[i] == runes[i-1] && runes[i] == runes[i-2] {
            steps++
        }
    }

    // Adjust for cases when length is less than 6
    if len(runes) < 6 {
        missing := 0
        if !lowercase {
            missing++
        }
        if !uppercase {
            missing++
        }
        if !digit {
            missing++
        }
        if missing > steps {
            steps = missing
        }
    } else {
        // Check if all character types are present
        if !lowercase {
            steps++
        }
        if !uppercase {
            steps++
        }
        if !digit {
            steps++
        }
    }

    return steps
}
```

This version of the function adjusts the steps for missing character types when the length is less than 6. If the number of missing types is greater than the steps needed to reach the minimum length, it sets steps equal to the number of missing types. This is because each addition to reach the minimum length can also be used to fulfill a missing type requirement. For lengths greater than or equal to 6, it increments the step for each missing type as before.

Please run the tests again with this update
