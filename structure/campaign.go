package structure

import (
	"encoding/json"

	"github.com/mongodb/mongo-go-driver/bson/objectid"
)

// Campaign is the entry point to structure data hierarchy.
type Campaign struct {
	ID          *objectid.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UID         int32              `json:"id" bson:"uid"`
	Name        string             `json:"name" bson:"name"`
	Goal        string             `json:"goal" bson:"goal"`
	TotalBudget int                `json:"total_budget" bson:"total_budget"`
	Status      string             `json:"status" bson:"status"`
	Platforms   struct {
		Facebook  *Platform `json:"facebook" bson:"facebook"`
		Google    *Platform `json:"google" bson:"google"`
		Instagram *Platform `json:"instagram" bson:"instagram"`
	} `json:"platforms" bson:"platforms"`
}

// We need a custom Marshaller just because `github.com/mongodb/mongo-go-driver/bson/objectID` has
// a unmarshaller implementation but not a marshaller for human readable ID. For them an ObjectID
// is a [12]byte.

//MarshalJSON is Marshaller implementation
func (c *Campaign) MarshalJSON() ([]byte, error) {
	c2 := &struct {
		ID          string `json:"_id,omitempty" bson:"_id,omitempty"`
		UID         int32  `json:"id" bson:"uid"`
		Name        string `json:"name" bson:"name"`
		Goal        string `json:"goal" bson:"goal"`
		TotalBudget int    `json:"total_budget" bson:"total_budget"`
		Status      string `json:"status" bson:"status"`
		Platforms   struct {
			Facebook  *Platform `json:"facebook" bson:"facebook"`
			Google    *Platform `json:"google" bson:"google"`
			Instagram *Platform `json:"instagram" bson:"instagram"`
		} `json:"platforms" bson:"platforms"`
	}{}

	c2.ID = c.ID.Hex()
	c2.UID = c.UID
	c2.Name = c.Name
	c2.Goal = c.Goal
	c2.TotalBudget = c.TotalBudget
	c2.Status = c.Status
	c2.Platforms = c.Platforms

	return json.Marshal(c2)
}
