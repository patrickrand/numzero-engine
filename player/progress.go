package player

// Progress is the corresponding outgoing event, resulting from an incoming game event.
type Progress struct {
	ID            string      `json:"id"`
	EventID       string      `json:"event_id"` // the associated triggering game event
	PlayerID      string      `json:"player_id"`
	AchievementID string      `json:"achievement_id"`
	GameID        string      `json:"game_id"`
	Status        interface{} `json:"status"` // TBD
}
