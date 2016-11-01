package player

import "github.com/patrickrand/numzero-engine/game"

// An Event is a real world or external activity a player was involved in that results in potential progress.
// The event should contain enough metadata to provide context around the activity that took place
// (e.g. Player1 had a pull request accepted by Linux Torvalds: (URL) - (Total Score)).
type Event struct {
	ID       string      `json:"id"`
	PlayerID string      `json:"player_id"`
	Facts    []game.Fact `json:"facts"`
}
