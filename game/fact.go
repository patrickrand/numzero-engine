package game

// A Fact is the empirical data extracted from an event (e.g. tests-added: 2, profanity-used: 5, etc).
// Zero or more facts can be extracted from an event.
// Facts have a preconfigured points value that contribute towards the total score of the event.
type Fact struct {
	ID     string      `json:"id"`
	Type   string      `json:"type"` // i.e. "sum", meaning the contents of "value" are numbers, and should be added
	Value  interface{} `json:"value,omitempty"`
	Points float64     `json:"points"`
}
