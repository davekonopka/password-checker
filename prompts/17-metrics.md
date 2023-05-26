# 17 Add metrics endpoint

## Prompt

Can you help me add a metrics endpoint that exposes metrics values in a format prometheus scraper would understand? I want to track a running count of requests while the server is running and expose that metric as request_count.

## Response

Sure, I can help you with that.

For this purpose, you'll need to use the Prometheus client library for Go. You can download it using go get:

shell
Copy code
go get github.com/prometheus/client_golang/prometheus
go get github.com/prometheus/client_golang/prometheus/promhttp
Once you have the library, you can create a new prometheus.Counter metric and increase it every time a request is handled. Then, you can expose these metrics in a new /metrics endpoint that the Prometheus server can scrape.

Here is how you can add the request counter and the metrics endpoint:

```go
package main

import (
	"fmt"
	"net/http"
	"strconv"
	"unicode"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	requestCount = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "request_count",
		Help: "Number of password check requests.",
	})
)

func init() {
	// Metrics have to be registered to be exposed:
	prometheus.MustRegister(requestCount)
}

func CheckPasswordStrength(password string) int {
	// ... (The rest of your code here)
}

func passwordHandler(w http.ResponseWriter, r *http.Request) {
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
		// ... (The rest of your code here)
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

In this code, I've added a new requestCount counter and incremented it in the passwordHandler function. I've also registered the counter in the init function.

The new /metrics endpoint is served by the promhttp.Handler(), which automatically generates a metrics report in the format Prometheus understands. You can now start the server and access http://localhost:8080/metrics to see the Prometheus metrics, including the request_count metric.

Remember to import the Prometheus packages at the start of your file:

```go
import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)
```
