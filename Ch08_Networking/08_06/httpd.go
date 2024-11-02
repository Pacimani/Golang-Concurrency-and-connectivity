// HTTP server example
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// healthHandler handles a simple health check, just responding with "OK".
func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")

}

// MathRequest is a request to MathHandler
type MathRequest struct {
	Op    string  `json:"operation"`
	Left  float64 `json:"number1"`
	Right float64 `json:"number2"`
}

// MathResponse is a response to MathRequest
type MathResponse struct {
	Error  string  `json:"error"`
	Result float64 `json:"result"`
}

// mathHandler returns result of calculation
func mathHandler(w http.ResponseWriter, r *http.Request) {

	// Decode and Validate
	defer r.Body.Close()

	dec := json.NewDecoder(r.Body)

	req := &MathRequest{}

	if err := dec.Decode(req); err != nil {
		log.Printf("Error: bad JSON -> %v\n", err)
		http.Error(w, "Bad JSON ", http.StatusBadRequest)
		return
	}

	if !strings.Contains("+-*/", req.Op) {
		log.Printf("Error: bad operation -> %q\n", req.Op)
		http.Error(w, "Bad Operation", http.StatusBadRequest)
		return
	}

	// work
	resp := &MathResponse{}

	switch req.Op {
	case "+":
		resp.Result = req.Left + req.Right
	case "-":
		resp.Result = req.Left - req.Right
	case "*":
		resp.Result = req.Left * req.Right
	case "/":
		if req.Right == 0 {
			resp.Error = "Cannot divide by zero"
		} else {
			resp.Result = req.Left / req.Right
		}

	default:
		resp.Error = "Unknown operation: " + req.Op
		return
	}

	// Encode result
	w.Header().Set("Content-Type", "application/json")
	if resp.Error != "" {
		w.WriteHeader(http.StatusBadRequest)
	}
	enc := json.NewEncoder(w)

	if err := enc.Encode(resp); err != nil {
		log.Printf("Error: bad JSON ->  %v - %s\n", resp, err)
		return
	}

	fmt.Fprintf(w, "%v", resp)

}

func main() {
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/math", mathHandler)

	addr := ":8080"
	log.Printf("Server is ready on %s\n", addr)

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Error: %v", err)
	}
}
