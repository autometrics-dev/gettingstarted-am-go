package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/autometrics-dev/autometrics-go/prometheus/autometrics"
	"github.com/autometrics-dev/autometrics-go/prometheus/midhttp"
)

//go:generate autometrics

func main() {
	autometrics.Init(
		nil,
		autometrics.DefBuckets,
		autometrics.BuildInfo{Version: "0.4.0", Commit: "anySHA", Branch: ""},
	)

	http.Handle("/json", midhttp.Autometrics(
		http.HandlerFunc(jsonHandler),
		autometrics.WithSloName("API"),
		autometrics.WithAlertLatency(100*time.Millisecond, 0.99),
	))
	http.Handle("/error", midhttp.Autometrics(
		http.HandlerFunc(errorHandler),
		autometrics.WithSloName("API"),
		autometrics.WithAlertLatency(100*time.Millisecond, 0.99),
	))
	http.Handle("/slow", midhttp.Autometrics(
		http.HandlerFunc(slowHandler),
		autometrics.WithSloName("API"),
		autometrics.WithAlertLatency(100*time.Millisecond, 0.99),
	))

	http.ListenAndServe(":8080", nil)
}

// jsonHandler is the handler function for the '/json' endpoint.
func jsonHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{"message": "Hello, World!"}
	json.NewEncoder(w).Encode(data)
}

// errorHandler is the handler function for the '/error' endpoint.
func errorHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Oops! Something went wrong.", http.StatusInternalServerError)
}

// slowHandler is the handler function for the '/slow' endpoint.
func slowHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(5 * time.Second)
	data := map[string]string{"message": "This response took 5 seconds to generate."}
	json.NewEncoder(w).Encode(data)
}
