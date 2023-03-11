package main

import (
	"example/broker"
	"fmt"
)

type UserRegistered struct {
	Name string
}

func main() {
	b := broker.New()
	broker.Subscribe(b, WelcomeUser)
	broker.Subscribe(b, SendInvoice)

	broker.Publish(b, UserRegistered{"Bob"})
}

func WelcomeUser(event UserRegistered) {
	fmt.Printf("Hello, %v\n", event.Name)
}

func SendInvoice(event UserRegistered) {
	fmt.Printf("Sending invoice to %v\n", event.Name)
}
