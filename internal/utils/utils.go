package utils

import (
	"encoding/json"
)

func ParseJsonBody(body []byte) (map[string]interface{}, error) {
	var data map[string]interface{}

	err := json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
