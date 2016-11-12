package engine

import "github.com/patrickrand/numzero-engine/event"

type Listener interface {
	Listen(events <-chan event.Event) chan<- error
}
