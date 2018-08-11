package structure

import (
	"reflect"
	"testing"
)

func TestCampaignStatus_String(t *testing.T) {
	tests := []struct {
		name    string
		cs      CampaignStatus
		want    string
		wantErr bool
	}{
		{name: "int 0 as a value", cs: CampaignStatus(0), want: "Scheduled", wantErr: false},
		{name: "int 1 as a value", cs: CampaignStatus(1), want: "Delivering", wantErr: false},
		{name: "int 2 as a value", cs: CampaignStatus(2), want: "Ended", wantErr: false},
		{name: "int 3 as a value", cs: CampaignStatus(3), want: "", wantErr: true}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.cs.String()
			if (err != nil) != tt.wantErr {
				t.Errorf("CampaignStatus.String() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CampaignStatus.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCampaignStatus_MarshalJSON(t *testing.T) {

	tests := []struct {
		name    string
		cs      CampaignStatus
		want    []byte
		wantErr bool
	}{
		{name: "int 0 as a value", cs: CampaignStatus(0), want: []byte("\"Scheduled\""), wantErr: false},
		{name: "int 1 as a value", cs: CampaignStatus(1), want: []byte("\"Delivering\""), wantErr: false},
		{name: "int 2 as a value", cs: CampaignStatus(2), want: []byte("\"Ended\""), wantErr: false},
		{name: "int 3 as a value", cs: CampaignStatus(3), want: []byte(nil), wantErr: true}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.cs.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("CampaignStatus.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CampaignStatus.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCampaignStatus_UnmarshalJSON(t *testing.T) {
	csv := CampaignStatus(0)

	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		cs      *CampaignStatus
		args    args
		wantErr bool
	}{
		{name: "\"Scheduled\"", cs: &csv, args: args{b: []byte("\"Scheduled\"")}, wantErr: false},
		{name: "\"Delivering\"", cs: &csv, args: args{b: []byte("\"Delivering\"")}, wantErr: false},
		{name: "\"Ended\"", cs: &csv, args: args{b: []byte("\"Ended\"")}, wantErr: false},
		{name: "\"NonExisting\"", cs: &csv, args: args{b: []byte("\"NonExisting\"")}, wantErr: true},
		{name: "Empty", cs: &csv, args: args{b: []byte("\"\"")}, wantErr: true}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.cs.UnmarshalJSON(tt.args.b); (err != nil) != tt.wantErr {
				t.Errorf("CampaignStatus.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
