package structure

//Creatives holds meta data and human friendlies.
type Creatives struct {
	Header1     string `json:"header_1" bson:"header1"`
	Header2     string `json:"header_2" bson:"header2"`
	Description string `json:"description" bson:"description"`
	URL         string `json:"url" bson:"url"`
	Image       string `json:"image" bson:"image"`
}
