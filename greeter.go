package main

type Greeter struct {
	message Message
}

// NewGreeter returns a new Greeter which will respond with the given Message.
func NewGreeter(m Message) *Greeter {
	return &Greeter{message: m}
}

func (g *Greeter) Greet() Message { return g.message }
