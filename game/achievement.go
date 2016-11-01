package game

// An Achievement is a named accomplishment with a set of conditions (rules) required for the achievement to be earned.
// These are challenges designed to entice players to try to meet the conditions to get the achievement.
type Achievement struct {
	// ID is unique identifier of an achievment, within a given game.
	ID     string `json:"id"`
	GameID string `json:"game_id"`
	Rules  []Rule `json:"rules"`
}

// Facts returns the global map of facts that are associated with an achievement's entire rule set.
func (a Achievement) Facts() map[string]Fact {
	facts := make(map[string]Fact)
	for _, r := range a.Rules {
		facts[r.Fact.ID] = r.Fact
	}
	return facts
}
