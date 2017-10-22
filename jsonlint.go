package jsonlint

import (
	"encoding/json"

	"github.com/pkg/errors"
)

// Lint returns a formatted & linted JSON as a slice of bytes.
func Lint(jsonData []byte) ([]byte, error) {

	var data map[string]interface{}
	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		return nil, errors.Wrap(err, "JSON unmarshaling failed")
	}

	jsonBytes, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return nil, errors.Wrap(err, "JSON marshaling failed")
	}

	return jsonBytes, nil
}
