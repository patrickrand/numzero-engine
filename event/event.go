package event

import (
	"encoding/json"
	"io"
	"log"
	"time"

	"github.com/patrickrand/numzero-engine/game"
)

type Event interface {
	ID() string
	Type() string
	Timestamp() time.Time
	Payload() interface{}
}

// PlayerActivity is a real world or external activity a player was involved in that results in potential progress.
// The event should contain enough metadata to provide context around the activity that took place
// (e.g. Player1 had a pull request accepted by Linus Torvalds: (URL) - (Total Score)).
type PlayerActivity struct {
	Event // temporary interface embedding while developing

	EventID        string      `json:"id"`
	EventTimestamp time.Time   `json:"timestamp"`
	PlayerID       string      `json:"player_id"`
	Facts          []game.Fact `json:"facts"`
}

// EngineShutdown represents a triggering event for shutting down a game engine.
type EngineShutdown struct {
	Event // temporary interface embedding while developing

	EventID        string    `json:"id"`
	EventTimestamp time.Time `json:"timestamp"`
	Message        string    `json:"message"`
}

// NewGame signals that a new game should be registered with a game engine.
type NewGame struct {
	Event // temporary interface embedding while developing

	EventID        string     `json:"id"`
	EventTimestamp time.Time  `json:"timestamp"`
	Game           *game.Game `json:"game"`
}

// AchievementAdded signals that a new achievement should be registered with a game engine.
type AchievementAdded struct {
	Event // temporary interface embedding while developing

	EventID          string            `json:"id"`
	EventTimestamp   time.Time         `json:"timestamp"`
	AchievementAdded *game.Achievement `json:"achievement"`
}

// SystemError indicates that an upstream error occured in the game engine's flow.
type SystemError struct {
	Event // temporary interface embedding while developing

	EventID        string    `json:"id"`
	EventTimestamp time.Time `json:"timestamp"`
	Error          string    `json:"error"`
}

// Log logs an event to output using a standardized format.
func Log(e Event) {
	if e == nil {
		log.Print("[ERROR] unable to log nil event")
		return
	}

	data, err := json.Marshal(e.Payload())
	if err != nil {
		log.Printf("[ERROR] failed to log event: %#v", e)
		return
	}

	log.Printf("[%s] [%s] %s", e.Type(), e.ID(), string(data))
}

// SetLogOutput sets the output writer that will be logged to.
func SetLogOutput(w io.Writer) {
	log.SetOutput(w)
}
