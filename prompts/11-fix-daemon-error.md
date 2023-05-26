# 11 Fix daemon error

## Prompt

Thanks. That's a lot better. I'm still getting an error when I run daemon mode: "Invalid number of arguments"
The server starts though and seems to function. Can you help fix the issue causing this error when I run daemon mode?

## Response

I apologize for the confusion. The "Invalid number of arguments" error is due to the check for the number of arguments in the rootCmd's Run function. It checks that there is exactly one argument (the password) when you run the command. However, when you run the command with the -d or --daemon flag, there are no additional arguments, and hence the error.

To fix this, we can add a check for the daemonMode in the rootCmd's Run function. If daemonMode is true, we do not need to check for arguments or perform the password check.

Here's the updated main.go:

```go
package main

import (
	"fmt"
	"net/http"
	"strconv"
	"unicode"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
)

// CheckPasswordStrength function as before

func handler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	password := vars["password"]
	steps := CheckPasswordStrength(password)
	fmt.Fprint(w, strconv.Itoa(steps))
}

func startServer() {
	r := mux.NewRouter()
	r.HandleFunc("/check/{password}", handler)
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
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
		startServer()
	}
}
```

With these changes, if you run the command in daemon mode with go run main.go -d or go run main.go --daemon, you should not see the "Invalid number of arguments" error. You can still run the command in CLI mode with go run main.go {password}.
