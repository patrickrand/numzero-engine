package engine

import "github.com/patrickrand/numzero-engine/player"

// An Engine orchestrates the playing of games.
type Engine struct {
	Players map[string]player.Player
}

func New() *Engine {
	return &Engine{}
}

// Run your engine.
func (e *Engine) Run(players chan *player.Player, stop chan struct{}) error {

	return nil
}
