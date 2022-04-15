/*
 * @Author: lwnmengjing<lwnmengjing@qq.com>
 * @Date: 2022/4/16 1:11
 * @Last Modified by: lwnmengjing<lwnmengjing@qq.com>
 * @Last Modified time: 2022/4/16 1:11
 */

package curd

type OneID struct {
	ID string `uri:"id" json:"-"`
}

func (e *OneID) GetID() string {
	return e.ID
}

func (e *OneID) GetIDS() []string {
	return []string{e.ID}
}

type ManyID struct {
	ID  string   `uri:"id" json:"-"`
	IDS []string `json:"ids"`
}

func (e *ManyID) GetIDS() []string {
	if len(e.IDS) == 0 && e.ID == "" {
		return nil
	}
	if e.IDS == nil {
		e.IDS = []string{e.ID}
		return e.IDS
	}
	e.IDS = append(e.IDS, e.ID)
	return e.IDS
}

type Pagination struct {
	Page     int64 `form:"page" query:"page"`
	PageSize int64 `form:"pageSize" query:"pageSize"`
}

func (e *Pagination) GetPage() int64 {
	if e.Page <= 0 {
		return 1
	}
	return e.Page
}

func (e *Pagination) GetPageSize() int64 {
	if e.PageSize <= 0 {
		return 10
	}
	return e.PageSize
}
