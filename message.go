package main

type Message string

// NewMessage returns a new Message.
func NewMessage() Message {
	return Message("Hello, World")
}
