package rtc

import (
	"encoding/json"
	"net/url"
)

// ListRooms ...
func ListRooms(r *RTC, query url.Values) (*ListRoomsResponse, int, error) {
	respBody, status, err := r.Client.Query(ActionListRooms, query)
	if err != nil {
		return nil, status, err
	}

	output := new(ListRoomsResponse)
	err = json.Unmarshal(respBody, output)
	return output, status, err
}

// KickUser ...
func ListIndicators(r *RTC, req *ListIndicatorsRequest) (*ListIndicatorsResponse, int, error) {
	bts, err := json.Marshal(req)
	if err != nil {
		return nil, 0, err
	}

	respBody, status, err := r.Client.Json(ActionListIndicators, nil, string(bts))
	if err != nil {
		return nil, status, err
	}

	output := new(ListIndicatorsResponse)
	err = json.Unmarshal(respBody, output)
	return output, status, err
}
