# 08 Add HTTP server

## Prompt

Thanks. Things are working well now. I would like to add a REST HTTP web server that takes in a password as a parameter, runs CheckPasswordStrength, and returns the return value. 

## Response

Alright, to expose the CheckPasswordStrength function over HTTP, you can use the net/http package in the Go standard library to create a web server. For a simple server, you can use the http.HandleFunc function to register a function to be called whenever a certain path is accessed.

Here is a simple server that provides the functionality you described:

```go
package main

import (
	"fmt"
	"net/http"
	"strconv"
	"unicode"

	"github.com/gorilla/mux"
)

// CheckPasswordStrength function as before

func handler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	password := vars["password"]
	steps := CheckPasswordStrength(password)
	fmt.Fprint(w, strconv.Itoa(steps))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/check/{password}", handler)
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
```

Here, we're using the Gorilla Mux router for handling the dynamic URL. You can install it with go get -u github.com/gorilla/mux.

You can start the server with go run main.go, and then access it at http://localhost:8080/check/{password} (replace {password} with the password you want to check).

Please note that sending passwords as URL parameters is generally not a good practice as URLs are often logged by systems, and the passwords could end up in plain text in some logs. Instead, it's recommended to use HTTPS and send sensitive information in a POST request body or headers, which are encrypted.

This server is also not validating inputs, and it will crash if the URL does not match the expected format. In a production application, you would want to add error checking and handle edge cases properly.
