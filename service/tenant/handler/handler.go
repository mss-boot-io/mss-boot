/*
 * @Author: lwnmengjing
 * @Date: 2021/6/17 4:12 下午
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2021/6/17 4:12 下午
 */

package handler

import (
	log "github.com/lwnmengjing/core-go/logger"
	"github.com/lwnmengjing/mss-boot/pkg/server/handler"
	pb "github.com/lwnmengjing/mss-boot/proto/tenant/v1"
	"gorm.io/gorm"
)

type Tenant struct {
	pb.UnimplementedTenantServer
	handler.Handler
	db *gorm.DB
}

// NewTenant new a tenant handler
func NewTenant(id string) *Tenant {
	return &Tenant{
		Handler: handler.Handler{
			ID:        id,
			Log:       log.NewHelper(log.DefaultLogger),
		},
	}
}
