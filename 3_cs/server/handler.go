package main

import (
	"context"
	service "experiment/3_cs/server/kitex_gen/kitex_gen/service"
)

// AddressImpl implements the last service interface defined in the IDL.
type AddressImpl struct{}

type AddressBook struct {
	ID      int    `gorm:"id" json:"id" `
	Name    string `gorm:"name" json:"姓名"`
	Tel     int64  `gorm:"tel" json:"电话"`
	Address string `gorm:"address" json:"地址"`
}

func (AddressBook) TableName() string {
	return "address_book"
}

// Add implements the AddressImpl interface.
func (s *AddressImpl) Add(ctx context.Context, req *service.AddReq) (resp *service.Res, err error) {
	resp = &service.Res{}
	var a AddressBook
	a.Address = req.Add
	a.Tel = req.Tel
	a.Name = req.Name
	affected := Db.Create(&a).RowsAffected
	if affected > 0 {
		resp.Code = 1
	} else {
		resp.Code = -1
	}
	return
}

// Modify implements the AddressImpl interface.
func (s *AddressImpl) Modify(ctx context.Context, req *service.AddReq) (resp *service.Res, err error) {
	resp = &service.Res{}
	var a AddressBook
	a.ID = int(req.Id)
	a.Name = req.Name
	a.Tel = req.Tel
	a.Address = req.Add
	affected := Db.Save(&a).RowsAffected
	if affected > 0 {
		resp.Code = 1
	} else {
		resp.Code = -1
	}
	return
}

// Delete implements the AddressImpl interface.
func (s *AddressImpl) Delete(ctx context.Context, req *service.AddReq) (resp *service.Res, err error) {
	resp = &service.Res{}
	var a AddressBook
	a.ID = int(req.Id)
	affected := Db.Delete(&a).RowsAffected
	if affected > 0 {
		resp.Code = 1
	} else {
		resp.Code = -1
	}
	return
}

// Select implements the AddressImpl interface.
func (s *AddressImpl) Select(ctx context.Context, req *service.AddReq) (resp *service.All, err error) {
	resp = &service.All{}
	resp.AddReqs = make([]*service.AddReq, 0)
	var a []AddressBook
	Db.Find(&a)
	for _, v := range a {
		res := &service.AddReq{
			Name: v.Name,
			Tel:  v.Tel,
			Add:  v.Address,
			Id:   int32(v.ID),
		}
		resp.AddReqs = append(resp.AddReqs, res)
	}
	return
}
