package yaml

import (
	"github.com/ghodss/yaml"
	"github.com/mss-boot-io/mss-boot/core/config/encoder"
)

type _encoder struct{}

// Encode func encode
func (y _encoder) Encode(v interface{}) ([]byte, error) {
	return yaml.Marshal(v)
}

// Decode func decode
func (y _encoder) Decode(d []byte, v interface{}) error {
	return yaml.Unmarshal(d, v)
}

// String string
func (y _encoder) String() string {
	return "yaml"
}

// NewEncoder new encoder
func NewEncoder() encoder.Encoder {
	return _encoder{}
}
