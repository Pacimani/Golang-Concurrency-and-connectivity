package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

// Request is a bank trnasaction
type Request struct {
	Login  string  `json:"user"`
	Type   string  `json:"type"`
	Amount float64 `json:"amount"`
}

var data = `{
	"user": "Scooge McDuck",
	"type": "withdraw",
	"amount": 123.4
}`

func main() {

	rdr := strings.NewReader(data) // Simulate a file/socket

	// decode request
	dec := json.NewDecoder(rdr)

	var req Request
	if err := dec.Decode(&req); err != nil {
		log.Fatalf("Error: can't decode: %s", err)
	}
	fmt.Printf("Contents:%+v\n", req)

	// create response
	prevBalance := 1_000_000.0 // loaded from database

	resp := map[string]interface{}{
		"ok":      true,
		"balance": prevBalance + req.Amount,
	}

	// encode response
	enc := json.NewEncoder(os.Stdout)
	if err := enc.Encode(resp); err != nil {
		log.Fatalf("Error: can't encode: %s", err)
	}
}
