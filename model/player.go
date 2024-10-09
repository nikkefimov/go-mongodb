package model

type Player struct {
	ID       string `json:"id" bson:"_id"`
	Nickname string `json:"nickname" bson:"nickname"`
	Fraction string `json:"fraction" bson:"fraction"`
	Level    string `json:"level" bson:"level:"`
}
