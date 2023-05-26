package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
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

func handler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	password := vars["password"]
	steps := CheckPasswordStrength(password)
	fmt.Fprint(w, strconv.Itoa(steps))
}

func startServer(stop chan bool) {
	r := mux.NewRouter()
	r.HandleFunc("/check/{password}", handler)
	http.Handle("/", r)

	srv := &http.Server{Addr: ":8080"}

	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			// unexpected error. port in use?
			log.Fatalf("ListenAndServe(): %v", err)
		}
	}()

	// returning in a non-blocking way
	go func() {
		<-stop
		if err := srv.Shutdown(context.Background()); err != nil {
			panic(err)
		}
	}()
}

var rootCmd = &cobra.Command{
	Use:   "password-checker",
	Short: "Password checker is a Go application to check password strength",
	Run: func(cmd *cobra.Command, args []string) {
		if daemonMode {
			return
		}
		if len(args) != 1 {
			fmt.Println("Invalid number of arguments")
			return
		}
		password := args[0]
		steps := CheckPasswordStrength(password)
		fmt.Println(steps)
	},
}

var daemonMode bool

func main() {
	rootCmd.PersistentFlags().BoolVarP(&daemonMode, "daemon", "d", false, "Start the password checker as a web server")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}

	if daemonMode {
		stop := make(chan bool)
		go startServer(stop)

		// Block main() until it's told to stop.
		<-stop
	}
}
