package toml

import (
	"bytes"

	"github.com/BurntSushi/toml"
	"github.com/mss-boot-io/mss-boot/core/config/encoder"
)

type _encoder struct{}

// Encode func encode
func (t _encoder) Encode(v interface{}) ([]byte, error) {
	b := bytes.NewBuffer(nil)
	defer b.Reset()
	err := toml.NewEncoder(b).Encode(v)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

// Decode func decode
func (t _encoder) Decode(d []byte, v interface{}) error {
	return toml.Unmarshal(d, v)
}

// String string
func (t _encoder) String() string {
	return "toml"
}

// NewEncoder new toml encoder
func NewEncoder() encoder.Encoder {
	return _encoder{}
}
