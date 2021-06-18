package errors

//go:generate stringer -type ErrCode -output error_code_string.go

// ErrCode 错误码类型
type ErrCode int32

const (

	// StoreSvcBaseErrorCode panic
	StoreSvcBaseErrorCode = 10000
	// TenantSvcBaseErrCode tenant error code
	TenantSvcBaseErrCode = 20000
)

const (
	// SUCCESS 约定成功为 0
	SUCCESS ErrCode = 0
)

const (
	// StoreSvcOperateAdapterFailed store适配器操作失败
	StoreSvcOperateAdapterFailed ErrCode = iota + StoreSvcBaseErrorCode
)

const (
	// TenantSvcBaseErrCode tenant数据库灿做失败
	TenantSvcOperateDBFailed ErrCode = iota + TenantSvcBaseErrCode
)

// Code 返回错误码
func (e ErrCode) Code() int32 {
	return int32(e)
}

// CheckErrorCode 判断错误码是否是成功错误吗
func CheckErrorCode(errCode int32) bool {
	return int32(SUCCESS) == errCode
}
