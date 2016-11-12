package engine

import (
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

// EventFactory returns the abstracted event for the given type.
// Forgive me, Rob.
func EventFactory(eventType string) Event {
	switch eventType {
	case "new_game":
		return &NewGame{}
	case "player_activity":
		return &PlayerActivity{}
	case "achievement_added":
		return &AchievementAdded{}
	case "system_error":
		return &SystemError{}
	case "shutdown":
		return &Shutdown{}
	default:
		return nil
	}
}

// PlayerActivity is a real world or external activity a player was involved in that results in potential progress.
// The event should contain enough metadata to provide context around the activity that took place
// (e.g. Player1 had a pull request accepted by Linus Torvalds: (URL) - (Total Score)).
type PlayerActivity struct {
	*MockEvent // temporary interface embedding while developing

	EventID        string      `json:"id"`
	EventType      string      `json:"type"`
	EventTimestamp time.Time   `json:"timestamp"`
	PlayerID       string      `json:"player_id"`
	Facts          []game.Fact `json:"facts"`
}

// Shutdown represents a triggering event for shutting down a game engine.
type Shutdown struct {
	*MockEvent // temporary interface embedding while developing

	EventID        string    `json:"id"`
	EventType      string    `json:"type"`
	EventTimestamp time.Time `json:"timestamp"`
	Message        string    `json:"message"`
}

// NewGame signals that a new game should be registered with a game engine.
type NewGame struct {
	*MockEvent // temporary interface embedding while developing

	EventID        string     `json:"id"`
	EventType      string     `json:"type"`
	EventTimestamp time.Time  `json:"timestamp"`
	Game           *game.Game `json:"game"`
}

type MockEvent struct {
	Event
}

func (mock *MockEvent) ID() string {
	return "mock-id"
}
func (mock *MockEvent) Type() string {
	return "mock-type"
}

// AchievementAdded signals that a new achievement should be registered with a game engine.
type AchievementAdded struct {
	*MockEvent // temporary interface embedding while developing

	EventID          string            `json:"id"`
	EventType        string            `json:"type"`
	EventTimestamp   time.Time         `json:"timestamp"`
	AchievementAdded *game.Achievement `json:"achievement"`
}

// SystemError indicates that an upstream error occured in the game engine's flow.
type SystemError struct {
	*MockEvent // temporary interface embedding while developing

	EventID        string    `json:"id"`
	EventType      string    `json:"type"`
	EventTimestamp time.Time `json:"timestamp"`
	Error          string    `json:"error"`
}

// Log logs an event to output using a standardized format.
func Log(e Event) {
	if e == nil {
		log.Print("[engine] unable to log nil event")
		return
	}

	log.Printf("[%s] [%s]", e.Type(), e.ID())
}

// SetLogOutput sets the output writer that will be logged to.
func SetLogOutput(w io.Writer) {
	log.SetOutput(w)
}
