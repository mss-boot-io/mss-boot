package json

import (
	json "github.com/json-iterator/go"

	"github.com/mss-boot-io/mss-boot/core/config/encoder"
)

type _encoder struct{}

// Encode func encode
func (j _encoder) Encode(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

// Decode func decode
func (j _encoder) Decode(d []byte, v interface{}) error {
	return json.Unmarshal(d, v)
}

// String string
func (j _encoder) String() string {
	return "json"
}

// NewEncoder new encoder
func NewEncoder() encoder.Encoder {
	return _encoder{}
}
