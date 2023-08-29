package response

/*
 * @Author: lwnmengjing
 * @Date: 2021/6/8 5:51 下午
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2021/6/8 5:51 下午
 */

// Responses responses
type Responses interface {
	SetCode(int)
	SetTraceID(string)
	SetMsg(...string)
	SetList(interface{})
	SetStatus(string)
	Clone() Responses
}
