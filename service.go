package main

import (
	"fmt"
	"net/http"
)

type Service struct {
	greeter *Greeter
}

// NewService returns a new Service which will greet clients using the provided Greeter.
func NewService(g *Greeter) *Service {
	return &Service{greeter: g}
}

func (s *Service) Run() error {
	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		s := string(s.greeter.Greet())
		fmt.Fprintf(w, s)
	})

	return nil
}
