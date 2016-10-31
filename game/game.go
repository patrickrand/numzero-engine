package game

type Game struct {
	ID           int           `json:"id"`
	Achievements []Achievement `json:"achievements"`
}

type Achievement struct {
	ID    int    `json:"id"`
	Rules []Rule `json:"rules"`
}

type Rule struct {
	ID   int                    `json:"id"`
	Type string                 `json:"type"`
	Args map[string]interface{} `json:"args"`
}

type Fact struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}
