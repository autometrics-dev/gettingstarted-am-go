package main

import (
    "encoding/json"
    "net/http"
    "time"
	"github.com/autometrics-dev/autometrics-go/prometheus/autometrics"

)

//go:generate autometrics

func main() {

	autometrics.Init(
			nil,
			autometrics.DefBuckets,
			autometrics.BuildInfo{Version: "0.4.0", Commit: "anySHA", Branch: ""},
		)
    http.HandleFunc("/json", jsonHandler)
    http.HandleFunc("/error", errorHandler)
    http.HandleFunc("/slow", slowHandler)

    http.ListenAndServe(":8080", nil)
}

//autometrics:inst --slo "API" --latency-target 99 --latency-ms 100
func jsonHandler(w http.ResponseWriter, r *http.Request) {
    data := map[string]string{"message": "Hello, World!"}
    json.NewEncoder(w).Encode(data)
}

//autometrics:inst --slo "API" --latency-target 99 --latency-ms 100
func errorHandler(w http.ResponseWriter, r *http.Request) {
    http.Error(w, "Oops! Something went wrong.", http.StatusInternalServerError)
}

//autometrics:inst --slo "API" --latency-target 99 --latency-ms 100
func slowHandler(w http.ResponseWriter, r *http.Request) {
    time.Sleep(5 * time.Second)
    data := map[string]string{"message": "This response took 5 seconds to generate."}
    json.NewEncoder(w).Encode(data)
}
