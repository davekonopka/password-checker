# 14 More test failures

## Prompt

Thanks. I made those changes but I'm still get errors when I run `go test`:

```
❯ go test
--- FAIL: TestCli (0.73s)
    main_test.go:21: Expected 0
        , got 2
Invalid number of arguments
--- FAIL: TestDaemon (1.02s)
    main_test.go:35: Expected no error, got Get "http://localhost:8080/check/password": dial tcp [::1]:8080: connect: connection refused
FAIL
exit status 1
FAIL	github.com/davekonopka/password-checker	2.341s
```

## Response

Apologies for the confusion and thank you for your patience. To correctly test both the command-line and HTTP server functionality, you would typically refactor your main() function so that it's more amenable to testing. However, as that may require a lot of changes, let's try another approach.

Firstly, please update the main.go file to use channels for better control of the HTTP server's lifecycle.

```go
package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

func startServer(stop chan os.Signal) {
	r := mux.NewRouter()
	r.HandleFunc("/check/{password}", handler)
	http.Handle("/", r)

	srv := &http.Server{Addr: ":8080"}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("ListenAndServe(): %v", err)
		}
	}()

	sig := <-stop
	log.Println("Received signal:", sig)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	srv.Shutdown(ctx)
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
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
		startServer(stop)
	}
}
```

Here, the startServer() function receives an os.Signal channel that listens for interrupt signals to gracefully stop the server.

Next, for the TestCli function in main_test.go, you need to be sure that the path to the Go executable is correctly configured in your environment variables, and the path to main.go in exec.Command() is correct.

```go
package main

import (
	"net/http"
	"os/exec"
	"strings"
	"testing"
	"time"
)

func TestCli(t *testing.T) {
	cmd := exec.Command("go", "run", "./main.go", "password") // Be sure this path is correct
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedOutput := "0\n"
	if strings.TrimSpace(string(output)) != expectedOutput {
		t.Fatalf("Expected %v, got %v", expectedOutput, string(output))
	}
}

func TestDaemon(t *testing.T) {
	go func() {
		cmd := exec.Command("go", "run", "./main.go", "-d") // Be sure this path is correct
		err := cmd.Start()
		if err != nil {
			t.Fatal(err)
		}
	}()

	// Give the server a moment to start
	time.Sleep(1 * time.Second)

	resp, err := http.Get("http://localhost:8080/check/password")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedBody := "0"
	if strings.TrimSpace(string(body)) != expectedBody {
		t.Fatalf("Expected %v, got %v", expectedBody, string(body))
	}
}
```

The above changes should help you test the command-line and HTTP server functionality of your application. Note that the test TestDaemon starts the server as a goroutine and sends a request to the server. After getting the response, the test should ideally also stop the server to clean up, but since the server is started in a separate process with exec.Command, it's a bit more complicated. It's also worth mentioning that running servers during unit tests is not a typical practice. Usually, such tests would be considered integration tests. Unit tests should ideally test the functions in isolation.
