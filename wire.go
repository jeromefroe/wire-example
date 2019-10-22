//+build wireinject

package main

import "github.com/google/wire"

// InitializeService initializes a new Service.
func InitializeService() (*Service, error) {
	wire.Build(NewService, NewGreeter, NewMessage)
	return &Service{}, nil
}
