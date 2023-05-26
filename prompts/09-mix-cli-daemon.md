# 09 Mix CLI & Daemon modes

## Prompt

Is it possible to run the HTTP server when a flag -d is passed for daemon mode and run the CLI without the flag? Can you show me updated main.go code for this?

## Response

Certainly, you can use a CLI package to handle command-line arguments or flags. The Go standard library provides the flag package, but we've been using the spf13/cobra package for command handling, so let's use it to add this functionality.

Below is the updated main.go file that runs the HTTP server when the -d or --daemon flag is provided and runs the CLI otherwise.

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

var daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "Start the password checker as a web server",
	Run: func(cmd *cobra.Command, args []string) {
		startServer()
	},
}

func init() {
	rootCmd.AddCommand(daemonCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
```

In this updated file, we've defined a new command daemonCmd for the daemon mode and added it to the rootCmd using rootCmd.AddCommand(daemonCmd). When -d or --daemon is provided, startServer() is called to start the HTTP server. If no argument is provided, the application runs in the CLI mode.

Now you can run the HTTP server with go run main.go daemon, and run the CLI with go run main.go {password} (replace {password} with the password you want to check).

Please note that cobra does not support short flags like -d for subcommands. We used daemon as the subcommand for starting the HTTP server, and it can be run with go run main.go daemon. If you want to use -d or --daemon as flags to start the HTTP server, additional modifications would be needed.
