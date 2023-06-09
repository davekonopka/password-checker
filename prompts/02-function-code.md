# 02 Function Code

## Prompt

Could you generate the code for step 1?

## Response

Sure, here's the code for CheckPasswordStrength:

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

    return steps
}
```

This function works by iterating over the password and checking each character. It uses the unicode package to check the character type. If it finds three identical characters in a row, it increments the step count. It also keeps track of whether it's seen a lowercase letter, an uppercase letter, and a digit, and increments the step count for each type that's missing. Finally, it checks the password's length and adjusts the step count if necessary.

Please note that this function only covers the basic requirements. It might not cover some edge cases and may not be optimal. I'd recommend testing it thoroughly and making adjustments as necessary.
