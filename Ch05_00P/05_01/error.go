package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/pkg/errors"
)

type Config struct {
}

func readConfig(path string) (*Config, error) {

	f, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open config file")
	}
	defer f.Close()

	cfg := &Config{}
	// parse file here

	return cfg, nil
}

func setUpLoading() {
	out, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	log.SetOutput(out)
}

func main() {

	start := time.Now()
	setUpLoading()
	cfg, err := readConfig("config.json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		log.Printf("Error: %v\n", err)
		fmt.Printf("Duration err: %v\n", time.Since(start))
		os.Exit(1)
	}
	fmt.Println(cfg)

	fmt.Printf("Duration: %v\n", time.Since(start))
}
