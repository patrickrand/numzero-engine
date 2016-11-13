package engine

import (
	"errors"
	"log"

	"github.com/patrickrand/numzero-engine/event"
	"github.com/patrickrand/numzero-engine/game"
	"github.com/patrickrand/numzero-engine/player"
	"github.com/patrickrand/numzero-engine/storage"
)

// An Engine orchestrates the playing of games.
type Engine struct {
	Handlers map[string]Handler
	Games    map[string]game.Game
	Players  map[string]player.Player
	Events   chan event.Event
	*storage.Storage
}

// New creates a new instance of a game engine.
func New(db *storage.Storage) *Engine {
	eng := &Engine{
		Games:    make(map[string]game.Game),
		Players:  make(map[string]player.Player),
		Events:   make(chan event.Event),
		Handlers: make(map[string]Handler),
		Storage:  db,
	}

	eng.registerDefaultHandlers()
	return eng
}

func (eng *Engine) registerDefaultHandlers() {
	eng.Handle("engine.new_game", HandlerFunc(eng.newGame))
	eng.Handle("engine.player_activity", HandlerFunc(eng.playerActivity))
	eng.Handle("engine.shutdown", HandlerFunc(eng.shutdown))
}

// Handle registers a mapping of event type to an event handler, for the given game engine.
func (eng *Engine) Handle(eventType string, h Handler) {
	eng.Handlers[eventType] = h
	log.Print("[engine] registered handler for event type:", eventType)
}

// Run the game engine.
func (eng *Engine) Run(events chan *event.Event, errs chan error) {
	go func(events chan *event.Event, errs chan error) {
		for evt := range events {
			if evt == nil {
				log.Print("[engine] recieved nil event, ignoring...")
				continue
			}

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

func (eng *Engine) newGame(evt *event.Event) error {
	if eng.Games == nil {
		return errors.New("game map cannot be nil")
	}

	var gm game.Game
	if err := evt.Bind(&gm); err != nil {
		return err
	}

	if _, ok := eng.Games[gm.ID]; ok {
		return errors.New("game already exists with ID: " + gm.ID)
	}

	eng.Games[gm.ID] = gm
	event.Save(eng.Storage, evt)
	return nil
}

func (eng *Engine) playerActivity(evt *event.Event) error {
	var p player.Player
	if err := evt.Bind(&p); err != nil {
		return err
	}
	event.Save(eng.Storage, evt)
	return nil
}
func (eng *Engine) shutdown(evt *event.Event) error {
	return nil
}

func (eng *Engine) systemError(evt *event.Event) error {
	return errors.New(evt.Key())
}
