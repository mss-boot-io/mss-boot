package security

/*
 * @Author: lwnmengjing
 * @Date: 2021/6/23 5:44 下午
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2021/6/23 5:44 下午
 */

// Verifier 用户验证器
type Verifier interface {
	GetUserID() string
	GetTenantID() string
	GetRoleID() string
	GetEmail() string
	GetUsername() string
	Verify(tenantID string, username string, password string) (bool, Verifier, error)
}
