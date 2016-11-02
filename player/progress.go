package player

type Progress struct {
	ID            string `json:"id"`
	PlayerID      string `json:"player_id"`
	AchievementID string `json:"achievement_id"`
	GameID        string `json:"game_id"`
}
