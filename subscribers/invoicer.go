package subscribers

import (
	"example/events"
	"fmt"
	"reflect"
)

type Invoicer struct{}

func NewInvoicer() *Invoicer {
	return &Invoicer{}
}

func (i Invoicer) Handler() (reflect.Type, func(any)) {
	var e events.UserRegistered

	return reflect.TypeOf(e), func(event any) {
		i.SendInvoice(event.(events.UserRegistered))
	}
}

func (i Invoicer) SendInvoice(event events.UserRegistered) {
	fmt.Printf("Sending invoice to %v\n", event.Name)
}
