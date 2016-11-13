package stream

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/patrickrand/numzero-engine/event"
)

type MockStream []event.Event

func NewMockStream(src string) *MockStream {
	f, err := os.Open(src)
	if err != nil {
		log.Panic("[stream] ", err)
	}

	var mockstream MockStream
	if err := json.NewDecoder(f).Decode(&mockstream); err != nil {
		log.Panic("[stream] ", err)
	}

	return &mockstream
}

func (m *MockStream) Stream() chan *event.Event {
	stream := make(chan *event.Event)
	go func(mockstream MockStream, stream chan *event.Event) {
		events := []event.Event(mockstream)
		for range time.Tick(250 * time.Millisecond) {
			evt := events[rand.Intn(len(events))]
			evt.Timestamp = time.Now().UTC()
			stream <- &evt
		}
	}(*m, stream)
	return stream
}
