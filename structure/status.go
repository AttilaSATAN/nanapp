package structure

// Status is the current status representation of a campaign or a single platform of a campaign.
// There is tree possible statuses. So status implemented on top of a custom Enum implementation with lookup tables.

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/mongodb/mongo-go-driver/bson"
)

const (
	//Scheduled is a campaign status that represents the campaign that going to be delivered in a future date.
	Scheduled CampaignStatus = 0
	//Delivering is a campaign status that represents currently active and in-delivery campaigns.
	Delivering CampaignStatus = 1
	//Ended is a campaign status that represents ended, not in-delivery campaigns.
	Ended CampaignStatus = 2
	//TODO: Resarch the possibility of using `iota`
)

var lookupCampaignStatuses map[string]int = map[string]int{
	"Scheduled":  0,
	"Delivering": 1,
	"Ended":      2}

var reverseLookupCampaignStatuses [3]string = [3]string{
	"Scheduled",
	"Delivering",
	"Ended"}

//CampaignStatus is an Enum type for status.
type CampaignStatus int

// String is string representation of CampaignStatus
func (cs CampaignStatus) String() (string, error) {
	if cs > 2 {
		return "", errors.New("value is out of the range of the CampaignStatus")
	}
	return reverseLookupCampaignStatuses[cs], nil
}

func (cs *CampaignStatus) UnmarshalBSON(b []byte) error {

	fmt.Println(b, string(b))
	*cs = 1
	return nil
}

// MarshalJSON is Marshaler implementation
func (cs CampaignStatus) MarshalJSON() ([]byte, error) {

	str, err := cs.String()

	if err != nil {
		return nil, err
	}
	return json.Marshal(str)
}

//UnmarshalJSON is Unmarshaler implementation
func (cs *CampaignStatus) UnmarshalJSON(b []byte) error {

	s := string(b)

	index, t := lookupCampaignStatuses[strings.Replace(s, "\"", "", -1)]
	tr(*cs)
	if t { // If JSON value does not exists in lookup table out of range error must be thrown
		*cs = CampaignStatus(index)
	} else {
		return errors.New("value is out of the range of the CampaignStatus")
	}
	return nil
}

func tr(v interface{}) {

	switch v.(type) {
	case bson.Unmarshaler:
		log.Println("YEEEEEEEEEEEEEEEEEEESSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSS")
	default:
		log.Println("FFFFUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUCCCCCCCCCCCCCCCCCCCCCCCCCCCKKKKKKKKKKKKKKKK")
	}
}
