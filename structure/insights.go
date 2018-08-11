package structure

// Insights is holds general metrics for the a single platform of a campaign.
type Insights struct {
	Impressions      int32   `json:"impressions" bson:"impressions"`
	Clicks           int32   `json:"clicks" bson:"clicks"`
	WebsiteVisits    int32   `json:"website_visits" bson:"websiteVisits"`
	CostPerClick     float32 `json:"cost_per_click" bson:"costPerClick"`
	ClickThroughRate float32 `json:"click_through_rate" bson:"clickThrough_rate"`
	AdvancedKPI1     float32 `json:"advanced_kpi_1" bson:"advancedKPI1"`
}
