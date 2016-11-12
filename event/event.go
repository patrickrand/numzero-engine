package event

import (
	"encoding/json"
	"fmt"
	"time"
)

// Event events.
type Event struct {
	Type      string                 `json:"type"`
	ID        string                 `json:"id"`
	Timestamp time.Time              `json:"timestamp"`
	Payload   map[string]interface{} `json:"payload"`
}

// Bind attempts to convert event's payload into the given struct.
func (evt *Event) Bind(v interface{}) error {
	data, err := json.Marshal(evt.Payload)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, v); err != nil {
		return err
	}

	return nil
}

// String returns the string representation of the event.
func (evt *Event) String() string {
	return fmt.Sprintf("%s:%s", evt.Type, evt.ID)
}
