package engine

import (
	"log"

	"github.com/patrickrand/numzero-engine/event"
)

type Handler interface {
	Handle(evt event.Event) error
}

type HandlerFunc func(event.Event) error

func (h HandlerFunc) Handle(evt event.Event) error {
	return h(evt)
}

func (eng *Engine) newGame(evt event.Event) error {

	return nil
}

func (eng *Engine) shutdown(evt event.Event) error {
	log.Print("[engine] mock shutting down...")
	return nil
}
