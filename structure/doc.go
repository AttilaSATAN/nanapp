// Package structure is model data for rempresenting an ad campaign.
// The struct types of this package designed to represent hierarchically structure
// of a campaign data.
// This structure structs also holds the marshaling and unmarshaling tags of JSON and BSON.
// In abstract it reflects the `data.json` file as Go structs.
//
// Campaign
//  |-> Status<CampaignStatus>
//  |-> Platforms
//  |   |-> Facebook <Platform>
//  |   |-> Google <Platform>
//  |   |-> Instagram <Platform>
//
// Platform
//  |-> Audiance <Audiance>
//  |-> Creatives <Creatives>
//  |-> Insights <Insights>
package structure
