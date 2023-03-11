package broker

import "reflect"

type EventBroker struct {
	subscribers map[reflect.Type]*list
}

type list struct {
	v []func(any)
}

func (l *list) append(f func(any)) {
	l.v = append(l.v, f)
}

func New() *EventBroker {
	return &EventBroker{subscribers: make(map[reflect.Type]*list)}
}

func Subscribe[T any](broker *EventBroker, handler func(event T)) {
	key := getKey[T]()

	l, ok := broker.subscribers[key]
	if !ok {
		l = &list{}
		broker.subscribers[key] = l
	}
	l.append(adapter(handler))
}

func Publish[T any](broker *EventBroker, event T) {
	key := getKey[T]()

	l, ok := broker.subscribers[key]
	if !ok {
		return
	}

	for _, handler := range l.v {
		handler(event)
	}
}

func getKey[T any]() reflect.Type {
	var tArr [0]T
	t := reflect.TypeOf(tArr).Elem()

	return t
}

func adapter[T any](handler func(T)) func(any) {
	return func(arg any) {
		handler(arg.(T))
	}
}
