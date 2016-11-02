package player

import "github.com/patrickrand/numzero-engine/game"

// A Player is a participant in a game, with the goal of earning achievements.
type Player struct {
	ID          string                         `json:"id"`
	ProgressMap map[string]map[string]Progress `json:"progress"` // map[game_id][achievment_id]Progress{}
}

// New returns a new instance of a player.
func New(id string) *Player {
	return &Player{
		ID:          id,
		ProgressMap: make(map[string]map[string]Progress),
	}
}

// Play the game.
func (p *Player) Play(g *game.Game, stop chan struct{}) chan *Progress {
	// ...
	return nil
}

// Save saves the given instance of progress for a player.
func (p *Player) Save(progress Progress) error {
	return nil
}
