package engine

import (
	"errors"

	"github.com/patrickrand/numzero-engine/event"
)

type Handler interface {
	Handle(evt *event.Event) error
}

type HandlerFunc func(*event.Event) error

func (h HandlerFunc) Handle(evt *event.Event) error {
	if evt == nil {
		return errors.New("unable to handle nil event")
	}
	return h(evt)
}
