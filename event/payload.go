package event

import "github.com/patrickrand/numzero-engine/game"

// PlayerActivity is a real world or external activity a player was involved in that results in potential progress.
// The event should contain enough metadata to provide context around the activity that took place
// (e.g. Player1 had a pull request accepted by Linus Torvalds: (URL) - (Total Score)).
type PlayerActivity struct {
	PlayerID string      `json:"player_id"`
	Facts    []game.Fact `json:"facts"`
}

// Shutdown represents a triggering event for shutting down a game engine.
type Shutdown struct {
	Message string `json:"message"`
}

// NewGame signals that a new game should be registered with a game engine.
type NewGame struct {
	Game *game.Game `json:"game"`
}

// AchievementAdded signals that a new achievement should be registered with a game engine.
type AchievementAdded struct {
	AchievementAdded *game.Achievement `json:"achievement"`
}

// SystemError indicates that an upstream error occured in the game engine's flow.
type SystemError struct {
	Error string `json:"error"`
}
