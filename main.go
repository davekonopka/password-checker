package main

import (
	"fmt"
	"net/http"
	"strconv"
	"unicode"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var log = logrus.New()

var (
	requestCount = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "request_count",
		Help: "Number of password check requests.",
	})
)

func init() {
	// Metrics have to be registered to be exposed:
	prometheus.MustRegister(requestCount)

	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringP("loglevel", "l", "info", "Set the logging level (options: debug, info, warn, error, fatal, panic")
	viper.BindPFlag("loglevel", rootCmd.PersistentFlags().Lookup("loglevel"))
}

func initConfig() {
	logLevel, _ := logrus.ParseLevel(viper.GetString("loglevel"))
	log.SetLevel(logLevel)
}

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

func passwordHandler(w http.ResponseWriter, r *http.Request) {
	log.Debug("Received a request for password strength check")
	vars := mux.Vars(r)
	password := vars["password"]
	steps := CheckPasswordStrength(password)
	fmt.Fprint(w, strconv.Itoa(steps))

	requestCount.Inc()
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "OK")
}

func startServer() {
	r := mux.NewRouter()
	r.HandleFunc("/check/{password}", passwordHandler)
	r.HandleFunc("/healthcheck", healthCheckHandler)
	r.Handle("/metrics", promhttp.Handler())

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
