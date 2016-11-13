package event

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"errors"

	"github.com/patrickrand/numzero-engine/storage"
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

	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	return nil
}

// Key returns the unique indentifer of the event.
func (evt *Event) Key() string {
	return fmt.Sprintf("%s:%s", evt.Type, evt.ID)
}

func Save(db *storage.Storage, evt *Event) error {
	if db == nil {
		return errors.New("cannot save event in nil database")
	}

	if evt == nil {
		return errors.New("cannot save nil event in database")
	}

	if err := db.Put(evt.Key(), evt); err != nil {
		return err
	}

	log.Print("[event] saved event to database: ", evt.Key())
	return nil
}

func Load(db *storage.Storage, key string) (*Event, error) {
	if db == nil {
		return nil, errors.New("cannot load event from nil database")
	}

	val, ok := db.Get(key)
	if !ok {
		return nil, errors.New("no event found in database for key: " + key)
	}

	evt, ok := val.(*Event)
	if !ok {
		return nil, errors.New("unrecognized event found in database for key: " + key)
	}

	log.Print("[event] loaded event from database: ", evt.Key())
	return evt, nil
}
