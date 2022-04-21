/*
 * @Author: lwnmengjing<lwnmengjing@qq.com>
 * @Date: 2022/4/16 0:23
 * @Last Modified by: lwnmengjing<lwnmengjing@qq.com>
 * @Last Modified time: 2022/4/16 0:23
 */

package curd

type ListRequester interface {
	GetPage() int64
	GetPageSize() int64
}

type DeleteRequester interface {
	GetIDS() []string
}

type GetRequester interface {
	GetID() string
}

type UpdateRequester interface {
	GetID() string
	SetUpdatedAt()
}

type CreateRequester interface {
	SetCreatedAt()
}
