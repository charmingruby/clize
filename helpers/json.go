package helpers

import "encoding/json"

func JSONDeserialize[T any](data []byte) (*T, error) {
	var item T

	if err := json.Unmarshal([]byte(data), &item); err != nil {
		return nil, err
	}

	return &item, nil
}

func JSONSerialize(v any) ([]byte, error) {

	data, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	return data, nil
}
