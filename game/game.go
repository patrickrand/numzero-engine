package game

// A Game is the set of achievements that are defined to incentivize a specific behavior or outcome.
// Although not a requirement, a game is generally associated with a specific practice, topic, or video game.
type Game struct {
	ID           string        `json:"id"`
	Achievements []Achievement `json:"achievements"`
}

// New returns a new instance of a game.
func New(id string) *Game {
	return &Game{ID: id}
}

// AddAchievement registers a new achievement for a given game, if an achievement
// with the same ID does not already exist for that game.
func (g *Game) AddAchievement(id string, rules []Rule) []Achievement {
	for _, a := range g.Achievements {
		if id == a.ID {
			return g.Achievements
		}
	}

	g.Achievements = append(g.Achievements, Achievement{
		ID:     id,
		GameID: g.ID,
		Rules:  rules,
	})

	return g.Achievements
}
