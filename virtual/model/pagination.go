/*
 * @Author: lwnmengjing
 * @Date: 2023/9/13 10:49:06
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2023/9/13 10:49:06
 */

package model

type PaginationImp interface {
	SetPageSize(int)
	SetCurrent(int)
	SetTotal(int64)
	GetCurrent() int
	GetPageSize() int
	GetTotal() int64
}
