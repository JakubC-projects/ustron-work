package schema

import (
	"github.com/gorilla/schema"
)

var queryDecoder *schema.Decoder
var queryEncoder *schema.Encoder

func init() {
	queryDecoder = schema.NewDecoder()
	queryDecoder.IgnoreUnknownKeys(true)

	queryEncoder = schema.NewEncoder()
}

func Encode(data any, dst map[string][]string) error {
	return queryEncoder.Encode(data, dst)
}

func Decode(dst any, src map[string][]string) error {
	return queryDecoder.Decode(dst, src)
}
