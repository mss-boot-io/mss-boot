// Code generated by "stringer -type ErrCode -output error_code_string.go"; DO NOT EDIT.

package errors

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[SUCCESS-0]
	_ = x[StoreSvcOperateAdapterFailed-10000]
	_ = x[GeneratorSvcParamsInvalid-20000]
	_ = x[GeneratorSvcOperateDBFailed-20001]
	_ = x[GeneratorSvcRecordIsExist-20002]
	_ = x[GeneratorSvcRecordNotFound-20003]
	_ = x[GeneratorSvcObjectIDInvalid-20004]
	_ = x[TenantSvcParamsInvalid-30000]
	_ = x[TenantSvcOperateDBFailed-30001]
	_ = x[TenantSvcRecordIsExist-30002]
	_ = x[TenantSvcRecordNotFound-30003]
	_ = x[TenantSvcObjectIDInvalid-30004]
	_ = x[AdminSvcParamsInvalid-40000]
	_ = x[AdminSvcUnauthorized-40001]
	_ = x[AdminSvcOperateDBFailed-40002]
	_ = x[AdminSvcForbidden-40003]
	_ = x[AdminSvcRecordIsExist-40004]
	_ = x[AdminSvcDeleteRecordNotExist-40005]
}

const (
	_ErrCode_name_0 = "SUCCESS"
	_ErrCode_name_1 = "StoreSvcOperateAdapterFailed"
	_ErrCode_name_2 = "GeneratorSvcParamsInvalidGeneratorSvcOperateDBFailedGeneratorSvcRecordIsExistGeneratorSvcRecordNotFoundGeneratorSvcObjectIDInvalid"
	_ErrCode_name_3 = "TenantSvcParamsInvalidTenantSvcOperateDBFailedTenantSvcRecordIsExistTenantSvcRecordNotFoundTenantSvcObjectIDInvalid"
	_ErrCode_name_4 = "AdminSvcParamsInvalidAdminSvcUnauthorizedAdminSvcOperateDBFailedAdminSvcForbiddenAdminSvcRecordIsExistAdminSvcDeleteRecordNotExist"
)

var (
	_ErrCode_index_2 = [...]uint8{0, 25, 52, 77, 103, 130}
	_ErrCode_index_3 = [...]uint8{0, 22, 46, 68, 91, 115}
	_ErrCode_index_4 = [...]uint8{0, 21, 41, 64, 81, 102, 130}
)

func (i ErrCode) String() string {
	switch {
	case i == 0:
		return _ErrCode_name_0
	case i == 10000:
		return _ErrCode_name_1
	case 20000 <= i && i <= 20004:
		i -= 20000
		return _ErrCode_name_2[_ErrCode_index_2[i]:_ErrCode_index_2[i+1]]
	case 30000 <= i && i <= 30004:
		i -= 30000
		return _ErrCode_name_3[_ErrCode_index_3[i]:_ErrCode_index_3[i+1]]
	case 40000 <= i && i <= 40005:
		i -= 40000
		return _ErrCode_name_4[_ErrCode_index_4[i]:_ErrCode_index_4[i+1]]
	default:
		return "ErrCode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
