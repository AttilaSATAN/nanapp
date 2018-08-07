package model

import (
	"encoding/json"
	"time"
)

type Platform struct {
	Status          CampaignStatus `json:"status" bson:"status"`
	TotalBudget     int            `json:"total_budget" bson:"totalBudget"`
	RemainingBudget int            `json:"remaining_budget" bson:"remainingBudget"`
	StartData       time.Time      `json:"start_date" bson:"startDate"`
	EndData         time.Time      `json:"end_date" bson:"endDate"`
	Audiance        Audiance       `json:"target_audiance" bson:"targetAudiance"` //typo?
	Creatives       Creatives      `json:"creatives" bson:"Creatives"`
	Insights        Insights       `json:"insights" bson:"insights"`
}

func (p *Platform) UnmarshalJSON(b []byte) error {

	tp := &struct {
		Status          CampaignStatus `json:"status" bson:"status"`
		TotalBudget     int            `json:"total_budget" bson:"totalBudget"`
		RemainingBudget int            `json:"remaining_budget" bson:"remainingBudget"`
		StartData       int64          `json:"start_date" bson:"startDate"`
		EndData         int64          `json:"end_date" bson:"endDate"`
		Audiance        Audiance       `json:"target_audiance" bson:"targetAudiance"`
		Creatives       Creatives      `json:"creatives" bson:"Creatives"`
		Insights        Insights       `json:"insights" bson:"insights"`
	}{}

	err := json.Unmarshal(b, tp)
	if err != nil {
		return err
	}

	// It's ugly but other alternative was creating another Time type and adding marshaling behavior.
	// I just don't want to dive too much deep for the backend in this simple API.
	np := Platform{tp.Status, tp.TotalBudget, tp.RemainingBudget, time.Unix(tp.StartData/1000, 0), time.Unix(tp.EndData/1000, 0), tp.Audiance, tp.Creatives, tp.Insights}

	*p = np

	return nil
}
