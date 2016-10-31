package game

type Game struct {
	ID           int                 `json:"id"`
	Achievements map[int]Achievement `json:"achievements"`
}

type Achievement struct {
	ID     string          `json:"id"` // unique within a given game
	GameID string          `json:"game_id"`
	Rules  []Rule          `json:"rules"`
	Facts  map[string]Fact `json:"facts"` // populated via parsing Rules
}

func NewAchievement(id, gameID string, rules []Rule) *Achievement {
	facts := make(map[string]Fact)
	for _, r := range rules {
		facts[r.Fact.ID] = r.Fact
	}

	return &Achievement{
		ID:     id,
		GameID: gameID,
		Rules:  rules,
		Facts:  facts,
	}
}

type Rule struct {
	ID   int    `json:"id"`
	Type string `json:"type"` // i.e. "count"
	Fact Fact   `json:"fact"`
}

type Fact struct {
	ID     string                 `json:"id"`
	Args   map[string]interface{} `json:"args"` // i.e. "total": 8
	Points int                    `json:"-"`    // ignore for now
}
