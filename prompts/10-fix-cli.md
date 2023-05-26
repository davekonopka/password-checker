# 10 Fix CLI

## Prompt 

I'm getting an error when I try to run CLI mode with the latest changes. Can you help me fix it so CLI mode works to take a password, run the password check, and return the result? The HTTP daemon mode works as expected.

```
❯ go run main.go sfkjsdlfkj
Error: unknown command "sfkjsdlfkj" for "password-checker"
Run 'password-checker --help' for usage.
unknown command "sfkjsdlfkj" for "password-checker"
```

## Response

I apologize for the confusion, and thank you for pointing out the issue. The error is due to misunderstanding of the command structure in the cobra package.

A more proper way to deal with this case would be to use flags for the daemon mode. This would keep the rootCmd to handle single argument as password in CLI mode.

Let's correct the implementation:

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
		if len(args) != 1 {
			fmt.Println("Invalid number of arguments")
			return
		}
		password := args[0]
		steps := CheckPasswordStrength(password)
		fmt.Println(steps)
	},
}

func main() {
	var daemonMode bool

	rootCmd.PersistentFlags().BoolVarP(&daemonMode, "daemon", "d", false, "Start the password checker as a web server")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}

	if daemonMode {
		startServer()
	}
}
```

In this updated version, we define a boolean flag daemon for the rootCmd. When the -d or --daemon flag is provided, startServer() is called to start the HTTP server. If no flag is provided, the application runs in the CLI mode.

You can now run the HTTP server with go run main.go -d or go run main.go --daemon, and run the CLI with go run main.go {password} (replace {password} with the password you want to check).
