package schema

import (
	"reflect"

	"github.com/google/uuid"
	"github.com/gorilla/schema"
)

var queryDecoder *schema.Decoder
var queryEncoder *schema.Encoder

func init() {
	queryDecoder = schema.NewDecoder()
	queryDecoder.IgnoreUnknownKeys(true)
	queryDecoder.RegisterConverter(uuid.UUID{}, func(s string) reflect.Value {
		uid, err := uuid.Parse(s)
		if err != nil {
			return reflect.ValueOf(uuid.Nil)
		}
		return reflect.ValueOf(uid)
	})

	queryEncoder = schema.NewEncoder()
	queryEncoder.RegisterEncoder(uuid.UUID{}, func(r reflect.Value) string {
		value := r.Interface().(uuid.UUID)
		return value.String()
	})
}

func Encode(data any, dst map[string][]string) error {
	return queryEncoder.Encode(data, dst)
}

func Decode(dst any, src map[string][]string) error {
	return queryDecoder.Decode(dst, src)
}
