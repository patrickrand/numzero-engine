package engine

import (
	"errors"
	"log"

	"github.com/patrickrand/numzero-engine/event"
	"github.com/patrickrand/numzero-engine/game"
	"github.com/patrickrand/numzero-engine/player"
)

var defaultHandlerFuncs = map[string]HandlerFunc{
	"engine.new_game":          nil,
	"engine.player_activity":   nil,
	"engine.achievement_added": nil,
	"engine.shutdown":          nil,
	"engine.error":             nil,
}

// An Engine orchestrates the playing of games.
type Engine struct {
	Handlers map[string]Handler
	Games    map[string]game.Game
	Players  map[string]player.Player
	Events   chan event.Event
}

// New creates a new instance of a game engine.
func New() *Engine {

	eng := &Engine{
		Games:    make(map[string]game.Game),
		Players:  make(map[string]player.Player),
		Events:   make(chan event.Event),
		Handlers: make(map[string]Handler),
	}

	eng.registerDefaultHandlers()
	return eng
}

func (eng *Engine) registerDefaultHandlers() {
	eng.Handle("engine.new_game", HandlerFunc(eng.newGame))
	eng.Handle("engine.shutdown", HandlerFunc(eng.shutdown))
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
			log.Printf("[engine] recieved event { %s %s %v }", evt.Type, evt.ID, evt.Payload)

			h, ok := eng.Handlers[evt.Type]
			if !ok {
				errs <- errors.New("no handler registered for event type: " + evt.Type)
				continue
			}

			if err := h.Handle(evt); err != nil {
				errs <- err
				continue
			}
		}
	}(events, errs)
}
