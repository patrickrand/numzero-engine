package stream

import "github.com/patrickrand/numzero-engine/event"

type Streamer interface {
	Stream() chan event.Event
}
