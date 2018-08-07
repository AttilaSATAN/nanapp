package model

import "github.com/mongodb/mongo-go-driver/bson/objectid"

type Campaign struct {
	ID          *objectid.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UID         int32              `json:"id" bson:"uid"`
	Name        string             `json:"name" bson:"name"`
	Goal        string             `json:"goal" bson:"goal"`
	TotalBudget int                `json:"total_budget" bson:"total_budget"`
	Status      CampaignStatus     `json:"status" bson:"status"`
	Platforms   struct {
		Facebook  *Platform `json:"facebook" bson:"facebook"`
		Google    *Platform `json:"google" bson:"google"`
		Instagram *Platform `json:"instagram" bson:"instagram"`
	} `json:"platforms" bson:"platforms"`
}
