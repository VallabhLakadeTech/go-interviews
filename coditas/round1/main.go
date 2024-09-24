package main

import (
	"log"
	"net/http"
	"time"

	"github.com/VallabhLakadeTech/coditas/round1/controller"
	"github.com/gorilla/mux"
)

// Create a POST API in Go using the Gin framework. The API should accept a JSON payload with the following fields:
// name: string
// pan: string (PAN number)
// mobile: number (mobile number)
// email: string (email ID)
// The PAN number should follow the format of five letters, followed by four digits, followed by a letter (e.g., ABCDE1234F). The mobile number should be a 10-digit number and email should be a valid email address and validate the PAN number and mobile number using the validator.v10 package.
// Additionally, you need to implement middleware to log the API latency for each request
// Requirements:
// Create a Gin router with the POST endpoint.
// Implement middleware to log the API latency.
// Use the validator.v10 package to validate the PAN number and mobile number.
// If the validation fails, you should respond with an appropriate error message and statuscode
// If the validation passes, the endpoint should respond with a success message and status code.

func main() {

	router := mux.NewRouter()
	router.Use(LoggingMiddleware)
	router.HandleFunc("/pan", controller.SavePANDetails).Methods("POST")

	http.ListenAndServe(":8080", router)

}

func LoggingMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		log.Println("Started %s %s", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
		log.Println("Completed %s in %v", r.RequestURI, time.Since(startTime))
	})

}
