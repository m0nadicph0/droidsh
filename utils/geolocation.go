package utils

import (
	"droidsh/constants"
	"encoding/json"
	"errors"
	"github.com/go-resty/resty/v2"
)

type Payload struct {
	Name      string  `json:"name"`
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
	Country   string  `json:"country"`
	State     string  `json:"state"`
}

func FetchLocation(appId string, place string) (Payload, error) {
	client := resty.New()
	resp, err := client.R().
		SetQueryParam("APPID", appId).
		SetQueryParam("q", place).
		Get(constants.OPENWM_URI)
	if err != nil {
		return Payload{}, err
	}
	locations := make([]Payload, 0)
	err = json.Unmarshal(resp.Body(), &locations)

	if err != nil {
		return Payload{}, err
	}

	if len(locations) == 0 {
		return Payload{}, errors.New("failed to fetch location")
	}
	return locations[0], err
}
