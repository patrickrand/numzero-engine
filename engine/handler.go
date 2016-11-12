package engine

import "github.com/patrickrand/numzero-engine/event"

type Handler interface {
	Handle(e event.Event) error
}

type HandlerFunc func(event.Event) error

func (h HandlerFunc) Handle(evt event.Event) error {
	return h(evt)
}

func shutdownHandler(evt event.Event) error {
	return nil
}
