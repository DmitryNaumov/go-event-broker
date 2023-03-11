package subscribers

import (
	"fmt"
	"reflect"

	"example/events"
)

type Welcomer struct{}

func NewWelcomer() *Welcomer {
	return &Welcomer{}
}

func (w Welcomer) WelcomeUser(event events.UserRegistered) {
	fmt.Printf("Hello, %v\n", event.Name)
}

func (w Welcomer) Handler() (reflect.Type, func(any)) {
	var e events.UserRegistered
	return reflect.TypeOf(e), func(event any) {
		w.WelcomeUser(event.(events.UserRegistered))
	}
}
