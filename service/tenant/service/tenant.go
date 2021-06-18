/*
 * @Author: lwnmengjing
 * @Date: 2021/6/17 5:04 下午
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2021/6/17 5:04 下午
 */

package service

import (
	"github.com/lwnmengjing/mss-boot/pkg/errors"
	"github.com/lwnmengjing/mss-boot/pkg/service"
	pb "github.com/lwnmengjing/mss-boot/proto/tenant/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strings"
	"tenant/models"
)

type Tenant struct {
	service.Service
}

func (e Tenant) DBError(err error) error {
	if err != nil {
		return nil
	}
	return e.MakeError("", "", errors.TenantSvcOperateDBFailed)
}

func (e Tenant) Create(req *pb.CreateReq) error {
	m := models.Tenant{
		Name:        req.Name,
		Status:      uint8(req.Status),
		System:      uint8(req.System),
		Contact:     req.Contact,
		Description: req.Description,
	}
	m.SetDomains(req.Domains...)
	return e.DBError(e.Orm.Create(&m).Error)
}

func (e Tenant) Get(req *pb.GetReq, resp *pb.GetResp) error {
	m := &models.Tenant{ID: req.Id}
	err := e.Orm.First(m).Error
	if err != nil {
		return e.DBError(err)
	}
	*resp = pb.GetResp{
		Id:          m.ID,
		Name:        m.Name,
		Status:      uint32(m.Status),
		System:      uint32(m.System),
		Contact:     m.Contact,
		Description: m.Description,
		Domains:     m.GetDomains(),
		CreatedAt:   timestamppb.New(m.CreatedAt),
		UpdatedAt:   timestamppb.New(m.UpdatedAt),
	}
	return err
}

func (e Tenant) Update(req *pb.UpdateReq) error {
	m := models.Tenant{
		ID:          req.Id,
		Name:        req.Name,
		Status:      uint8(req.Status),
		System:      uint8(req.System),
		Contact:     req.Contact,
		Domains: strings.Join(req.Domains, ","),
		Description: req.Description,
	}
	return e.DBError(e.Orm.Save(&m).Error)
}

func (e Tenant) Delete(req *pb.DeleteReq) error {
	m := &models.Tenant{ID: req.Id}
	return e.DBError(e.Orm.Delete(m).Error)
}

//func (e Tenant) List(req *pb.ListReq, resp *pb.ListResp) error {
//
//}
