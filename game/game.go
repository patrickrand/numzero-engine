package game

type Game struct {
	ID           int                 `json:"id"`
	Achievements map[int]Achievement `json:"achievements"`
	Players      map[string]Player   `json:"players"`
}

type Player struct {
	ID       int                    `json:"id"`
	Progress map[string]interface{} `json:"progress"` // map[Achievement.ID]interface{}
}

type Achievement struct {
	ID    int          `json:"id"`
	Rules []Rule       `json:"rules"`
	Facts map[Fact]int `json:"facts"` // built via parsing Rules
}

type Rule struct {
	ID   int                    `json:"id"`
	Type string                 `json:"type"` // i.e. "count"
	Args map[string]interface{} `json:"args"` // map["fact"]:Fact, map["total"]:int
}

type Fact string
