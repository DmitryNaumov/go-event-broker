package main

import (
	"example/broker"
	"example/events"
	"example/subscribers"

	"go.uber.org/fx"
)

func main() {

	fx.New(
		fx.Provide(
			broker.New,
			asSubscriber(subscribers.NewInvoicer),
			asSubscriber(subscribers.NewWelcomer),
		),
		fx.Invoke(fx.Annotate(registerSubscribers, fx.ParamTags(`group:"handlers"`))),
		fx.Invoke(run),
	).Run()
}

func registerSubscribers(list []broker.Subscriber, b *broker.EventBroker) {
	for _, v := range list {
		b.Subscribe(v)
	}
}

func run(b *broker.EventBroker) {
	broker.Publish(b, events.UserRegistered{"Bob"})
}

func asSubscriber(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(broker.Subscriber)),
		fx.ResultTags(`group:"handlers"`),
	)
}
