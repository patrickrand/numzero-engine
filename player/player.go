package player

import "github.com/patrickrand/numzero-engine/game"

// A Player is a participant in a game, with the goal of earning achievements.
type Player struct {
	ID string `json:"id"`
}

// Play the game.
func (p *Player) Play(g *game.Game, start, stop chan struct{}) chan *Progress {
	// ...
	return nil
}
