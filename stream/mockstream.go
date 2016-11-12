package stream

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/patrickrand/numzero-engine/engine"
	"github.com/patrickrand/numzero-engine/event"
)

type MockStream []engine.Event

func NewMockStream(src string) *MockStream {
	f, err := os.Open(src)
	if err != nil {
		log.Panic("[stream] ", err)
	}

	data := make([]map[string]interface{}, 0)
	if err := json.NewDecoder(f).Decode(&data); err != nil {
		log.Panic("[stream] ", err)
	}

	mockstream := make(MockStream, 0)
	for _, e := range data {
		t, ok := e["type"].(string)
		if !ok {
			log.Panic("[stream] key 'type' is missing, or unrecognized")
		}

		evt := event.Factory(t)
		if evt == nil {
			log.Panic("[stream] unrecognized event type: ", t)
		}

		mockstream = append(mockstream, evt)
	}

	return &mockstream
}

func (m *MockStream) Stream() chan engine.Event {
	stream := make(chan engine.Event)
	go func(mockstream MockStream, stream chan engine.Event) {
		var idx int
		n := len(mockstream)
		events := []engine.Event(mockstream)
		for range time.Tick(250 * time.Millisecond) {
			idx = rand.Intn(n)

			stream <- events[idx]
		}
	}(*m, stream)
	return stream
}
