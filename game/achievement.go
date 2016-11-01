package game

// An Achievement is a named accomplishment with a set of conditions required for the achievement to be earned.
// These are challenges designed to entice players to try to meet the conditions to get the achievement.
type Achievement struct {
	// ID is unique identifier of an achievment, within a given game.
	ID     string `json:"id"`
	GameID string `json:"game_id"`
	Rules  []Rule `json:"rules"`
}
