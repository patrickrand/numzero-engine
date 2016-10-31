package game

type Player struct {
	ID       string                 `json:"id"`
	Games    map[string]Game        `json:"games"`
	Progress map[string]interface{} `json:"progress"` // map[Achievement.ID]interface{}
}
