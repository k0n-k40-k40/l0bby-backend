package utils

import (
	"encoding/json"
)

const (
	DatabaseSource = "root:l0bby@tcp(172.17.0.2:3306)/l0bby"
)

func ParseJsonBody(body []byte) (map[string]interface{}, error) {
	var data map[string]interface{}

	err := json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
