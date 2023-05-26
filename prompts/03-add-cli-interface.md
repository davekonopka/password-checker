# 03 Add CLI Interface

## Prompt

Could you connect this function to a command line interface using the spf13/cobra Go library in the Go app you wrote?

## Response

Sure, I can provide an example of how to implement a command line interface using the spf13/cobra library.

To use the library, you will first need to install it with the following command:

```bash
go get -u github.com/spf13/cobra/cobra
```

Then you can create a new command that uses the CheckPasswordStrength function. Here is an example:

```go
package main

import (
    "fmt"
    "os"
    "strings"
    "unicode"

    "github.com/spf13/cobra"
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

var rootCmd = &cobra.Command{
    Use:   "password_checker",
    Short: "Password strength checker",
    Long:  `This application checks the strength of a password.`,
    Run: func(cmd *cobra.Command, args []string) {
        password := strings.Join(args, " ")
        steps := CheckPasswordStrength(password)
        fmt.Printf("Steps required to make the password strong: %d\n", steps)
    },
}

func main() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
```

This script will create a command line application that accepts a password as an argument. The Run function inside the rootCmd will call CheckPasswordStrength and print the number of steps required to make the password strong.

To run the application, you can simply use the following command (replace "password" with the password you want to check):

```bash
go run main.go password
```
