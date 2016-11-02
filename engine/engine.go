package engine

import (
	"os"

	"github.com/patrickrand/numzero-engine/event"
	"github.com/patrickrand/numzero-engine/player"
)

// An Engine orchestrates the playing of games.
type Engine struct {
	Players map[string]player.Player
	Events  chan event.Event
}

// New creates a new instance of a game engine.
func New() *Engine {
	return &Engine{
		Players: make(map[string]player.Player),
		Events:  make(chan event.Event),
	}
}

// Run the game engine.
func (e *Engine) Run() {
	for ev := range e.Events {
		switch ev := ev.(type) {
		case event.PlayerActivity:
		case event.AchievementAdded:
		case event.NewGame:
		case event.SystemError:
		case event.EngineShutdown:
			event.Log(ev)
			os.Exit(0)
		default:
		}
	}
}
