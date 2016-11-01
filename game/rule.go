package game

// A Rule maps a fact to an objective, to be used within the context of an achievement.
type Rule struct {
	ID        int         `json:"id"`
	Fact      Fact        `json:"fact"`
	Objective interface{} `json:"objective"` // i.e. 50, assuming "fact.type" == "sum"
}
