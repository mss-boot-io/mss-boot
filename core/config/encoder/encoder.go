// Package encoder handles source encoding formats
package encoder

import "fmt"

// Encoder is the encoder from which cfg is encoded
type Encoder interface {
	fmt.Stringer
	Encode(interface{}) ([]byte, error)
	Decode([]byte, interface{}) error
}
