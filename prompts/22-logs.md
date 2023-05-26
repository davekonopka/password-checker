# 21 Logs

## Prompt

Can you help me add stdout logging to the password-checker web service? I'd like to log a message whenever a request is made. I'd like the message to be log level debug and offer a way to run the daemon with an optional log level setting to only show messages at that level or higher.

## Response

To implement this feature in Go, we can use a popular logging library such as [logrus](https://github.com/sirupsen/logrus). Logrus is a structured logger for Go, completely API compatible with the standard library logger. Here is how you can implement logrus in your code:

Firstly, install the library by running:

```bash
go get github.com/sirupsen/logrus
```

Then, adjust your main.go as follows:

```go
package main

import (
	"fmt"
	"net/http"
	"strconv"
	"unicode"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var log = logrus.New()

// CheckPasswordStrength ...
func CheckPasswordStrength(password string) int {
    // ...
}

func handler(w http.ResponseWriter, r *http.Request) {
    // ...
    log.Debug("Received a request for password strength check")
    // ...
}

func startServer() {
    // ...
}

var rootCmd = &cobra.Command{
    // ...
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringP("loglevel", "l", "info", "Set the logging level (options: debug, info, warn, error, fatal, panic")
	viper.BindPFlag("loglevel", rootCmd.PersistentFlags().Lookup("loglevel"))
}

func initConfig() {
	logLevel, _ := logrus.ParseLevel(viper.GetString("loglevel"))
	log.SetLevel(logLevel)
}

var daemonMode bool

func main() {
    // ...
}
```

In this code, I added an optional flag for log level that accepts the following options: debug, info, warn, error, fatal, panic. Depending on the log level set, only logs of that level or higher will be shown. By default, if no log level is set, it will show info and above.

You can now use the command `go run main.go -d --loglevel=debug` to run the daemon with debug log level. The log message will be printed each time a request is made.
