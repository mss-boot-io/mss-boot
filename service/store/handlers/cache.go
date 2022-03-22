/*
 * @Author: lwnmengjing
 * @Date: 2022/3/22 10:17
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2022/3/22 10:17
 */

package handlers

import (
	"context"
	"store/pkg/storage"
	"time"

	pb "github.com/mss-boot-io/mss-boot/proto/store/v1"
)

func (e StoreHandler) Get(c context.Context, req *pb.GetReq) (resp *pb.GetResp, err error) {
	e.Make(c)
	e.Log.Info(req)
	resp = &pb.GetResp{}
	resp.Value, _ = storage.Cache.Get(c, req.Key)
	return
}

func (e StoreHandler) Set(c context.Context, req *pb.SetReq) (resp *pb.SetResp, err error) {
	e.Make(c)
	e.Log.Info(req)
	resp = &pb.SetResp{}
	err = storage.Cache.Set(c, req.Key, req.Value, int(req.Expire))
	return
}

func (e StoreHandler) Del(c context.Context, req *pb.DelReq) (resp *pb.DelResp, err error) {
	e.Make(c)
	e.Log.Info(req)
	resp = &pb.DelResp{}
	err = storage.Cache.Del(c, req.Key)
	return
}

func (e StoreHandler) HashGet(c context.Context, req *pb.HashGetReq) (resp *pb.HashGetResp, err error) {
	e.Make(c)
	e.Log.Info(req)
	resp = &pb.HashGetResp{}
	resp.Value, err = storage.Cache.HashGet(c, req.HashKey, req.Key)
	return
}

func (e StoreHandler) HashDel(c context.Context, req *pb.HashDelReq) (resp *pb.HashDelResp, err error) {
	e.Make(c)
	e.Log.Info(req)
	resp = &pb.HashDelResp{}
	err = storage.Cache.HashDel(c, req.HashKey, req.Key)
	return
}

func (e StoreHandler) Increase(c context.Context, req *pb.IncreaseReq) (resp *pb.IncreaseResp, err error) {
	e.Make(c)
	e.Log.Info(req)
	resp = &pb.IncreaseResp{}
	err = storage.Cache.Increase(c, req.Key)
	return
}

func (e StoreHandler) Decrease(c context.Context, req *pb.DecreaseReq) (resp *pb.DecreaseResp, err error) {
	e.Make(c)
	e.Log.Info(req)
	resp = &pb.DecreaseResp{}
	err = storage.Cache.Decrease(c, req.Key)
	return
}

func (e StoreHandler) Expire(c context.Context, req *pb.ExpireReq) (resp *pb.ExpireResp, err error) {
	e.Make(c)
	e.Log.Info(req)
	resp = &pb.ExpireResp{}
	err = storage.Cache.Expire(c, req.Key, time.Duration(req.Expire)*time.Second)
	return
}
