package main

import (
	"fmt"
	"os"
)

func main() {
	s, err := InitializeService()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to initialize service: %v.", err)
		os.Exit(1)
	}

	if err := s.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Unable to run service: %v.", err)
		os.Exit(1)
	}
}
