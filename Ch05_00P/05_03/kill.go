package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
)

func killServer(pidFile string) error {

	file, err := os.Open(pidFile)

	if err != nil {
		return err
	}
	defer file.Close()

	var pid int

	if _, err := fmt.Fscanf(file, "%d", &pid); err != nil {
		return errors.Wrap(err, "Bad process id")
	}

	// simulate kill
	fmt.Printf("Killing server of process ID %d\n", pid)
	//fmt.Printf("Killing server of process ID %d\n", pid)

	if err := os.Remove(pidFile); err != nil {
		log.Printf("Warning: can't remove PID file: - %s\n", err)
	}
	return nil
}

func main() {

	err := killServer("server.pid")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
