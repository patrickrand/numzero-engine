package player

import "github.com/patrickrand/numzero-engine/game"

// A Player is a participant in a game who can earn achievements.
type Player struct {
	ID       string                 `json:"id"`
	Games    map[string]game.Game   `json:"games"`
	Progress map[string]interface{} `json:"progress"` // map[Achievement.ID]interface{}
}
