package main

import (
	"log"

	"github.com/patrickrand/numzero-engine/engine"
	"github.com/patrickrand/numzero-engine/event"
	"github.com/patrickrand/numzero-engine/storage"
	"github.com/patrickrand/numzero-engine/stream"
)

var events = make(chan *event.Event)

func main() {
	incoming := stream.NewMockStream("stream/incoming.json")
	go func(st stream.Streamer) {
		for evt := range st.Stream() {
			events <- evt
		}
		close(events)
	}(incoming)

	db := storage.New()

	errs := make(chan error)
	go func(events chan *event.Event, errs chan error) {
		engine.New(db).Run(events, errs)
	}(events, errs)

	for err := range errs {
		log.Print("[main] ", err)
	}
}
