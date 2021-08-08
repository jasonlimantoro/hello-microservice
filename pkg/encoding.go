package pkg

import (
	"encoding/json"
	"io"
)

func DecodeJSON(body io.Reader, result interface{}) error {
	decoder := json.NewDecoder(body)
	decoder.DisallowUnknownFields()
	return decoder.Decode(&result)
}

func EncodeJSON(writer io.Writer, data interface{}) error {
	encoder := json.NewEncoder(writer)
	return encoder.Encode(data)
}
