package stream

import "github.com/patrickrand/numzero-engine/engine"

type Streamer interface {
	Stream() chan engine.Event
}
