package virtual

/*
 * @Author: lwnmengjing<lwnmengjing@qq.com>
 * @Date: 2023/9/17 08:12:38
 * @Last Modified by: lwnmengjing<lwnmengjing@qq.com>
 * @Last Modified time: 2023/9/17 08:12:38
 */

// Pagination pagination params
type Pagination struct {
	Page     int64 `form:"page" query:"page"`
	PageSize int64 `form:"pageSize" query:"pageSize"`
}

// GetPage get page
func (e *Pagination) GetPage() int64 {
	if e.Page <= 0 {
		return 1
	}
	return e.Page
}

// GetPageSize get page size
func (e *Pagination) GetPageSize() int64 {
	if e.PageSize <= 0 {
		return 10
	}
	return e.PageSize
}
