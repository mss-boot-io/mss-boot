/*
 * @Author: lwnmengjing
 * @Date: 2021/6/17 10:19 上午
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2021/6/17 10:19 上午
 */

package form

type CategoryCreateReq struct {
	//名称
	Name string `json:"name" comment:"名称"`
	//描述
	Description string `json:"description" comment:"描述"`
}

type CategoryCreateResp struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"createdAt"`
	UpdateAt    int64  `json:"updateAt"`
}
