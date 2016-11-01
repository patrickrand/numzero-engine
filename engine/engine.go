package engine

import "github.com/patrickrand/numzero-engine/player"

// An Engine orchestrates the playing of a game.
type Engine struct {
	Players map[string]player.Player
}
