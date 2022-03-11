package xml

import (
	"encoding/xml"

	"github.com/mss-boot-io/mss-boot/core/config/encoder"
)

type _encoder struct{}

// Encode encode
func (x _encoder) Encode(v interface{}) ([]byte, error) {
	return xml.Marshal(v)
}

// Decode decode
func (x _encoder) Decode(d []byte, v interface{}) error {
	return xml.Unmarshal(d, v)
}

// String string
func (x _encoder) String() string {
	return "xml"
}

// NewEncoder new a xml encoder
func NewEncoder() encoder.Encoder {
	return _encoder{}
}
