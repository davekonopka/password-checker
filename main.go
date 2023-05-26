package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"unicode"

	"github.com/gorilla/mux"
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

func handler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	password := vars["password"]
	steps := CheckPasswordStrength(password)
	fmt.Fprint(w, strconv.Itoa(steps))
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		r := mux.NewRouter()
		r.HandleFunc("/check/{password}", handler)
		http.Handle("/", r)
		http.ListenAndServe(":8080", nil)
	}
}
