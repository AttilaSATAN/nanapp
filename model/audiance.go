package model

type Audiance struct {
	Languages []string `bson:"languages" json:"languages"`
	Genders   []string `bson:"genders" json:"genders"`
	AgeRange  [2]int   `bson:"ageRange" json:"age_range"`
	Locations []string `bson:"locations" json:"locations"`
	Interests []string `bson:"interests" json:"interests"`
}
