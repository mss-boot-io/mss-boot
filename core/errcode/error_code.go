package errcode

//go:generate stringer -type ErrCode -output error_code_string.go

// ErrCode error code
type ErrCode int32

const (
	ok   = 0
	gRPC = 50000
)

const (
	// Success success code
	Success ErrCode = ok
)

const (
	// GRPCInternalServerError gRPC Internal Server Error
	GRPCInternalServerError ErrCode = gRPC + iota
)

// Code code
func (e ErrCode) Code() int32 {
	return int32(e)
}
