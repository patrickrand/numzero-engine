package engine

import (
	"errors"
	"log"

	"github.com/patrickrand/numzero-engine/event"
	"github.com/patrickrand/numzero-engine/player"
)

var defaultHandlerFuncs = map[string]HandlerFunc{
	"engine.new_game":          nil,
	"engine.player_activity":   nil,
	"engine.achievement_added": nil,
	"engine.shutdown":          shutdownHandler,
	"engine.error":             nil,
}

// An Engine orchestrates the playing of games.
type Engine struct {
	Handlers map[string]Handler
	Players  map[string]player.Player
	Events   chan event.Event
}

// New creates a new instance of a game engine.
func New() *Engine {

	eng := &Engine{
		Players:  make(map[string]player.Player),
		Events:   make(chan event.Event),
		Handlers: make(map[string]Handler),
	}

	for eventType, handler := range defaultHandlerFuncs {
		eng.Handle(eventType, handler)
	}

	return eng
}

// Handle registers a mapping of event type to an event handler, for the given game engine.
func (eng *Engine) Handle(eventType string, h Handler) {
	eng.Handlers[eventType] = h
	log.Print("[engine] registered handler for event type:", eventType)
}

// Run the game engine.
func (eng *Engine) Run(events chan event.Event, errs chan error) {
	go func(events chan event.Event, errs chan error) {
		for evt := range events {
			log.Print("[engine] recieved event - ", event.String(evt))
			h, ok := eng.Handlers[evt.Type()]
			if !ok || h == nil {
				errs <- errors.New("no event handler registered for type=" + evt.Type())
				continue
			}

			if err := h.Handle(evt); err != nil {
				errs <- err
				continue
			}
		}
	}(events, errs)
}
