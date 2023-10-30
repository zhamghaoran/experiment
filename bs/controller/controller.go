package controller

import (
	"context"
	"experiment/2_cs/DB"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type AddressBook struct {
	ID      int    `gorm:"id" json:"id"`
	Name    string `gorm:"name" json:"姓名" `
	Tel     int64  `gorm:"tel" json:"电话"`
	Address string `gorm:"address" json:"地址"`
}

func (AddressBook) TableName() string {
	return "address_book"
}
func Add(c context.Context, ctx *app.RequestContext) {
	var a AddressBook
	err := ctx.Bind(&a)
	if err != nil {
		ctx.JSON(consts.StatusInternalServerError, utils.H{
			"message": "参数错误",
		})
	}
	if DB.Db.Create(&a).RowsAffected > 0 {
		ctx.JSON(consts.StatusOK, utils.H{
			"message": "ok",
		})
	} else {
		ctx.JSON(consts.StatusInternalServerError, utils.H{
			"message": "参数错误",
		})
	}
}
func Show(c context.Context, ctx *app.RequestContext) {
	var a []AddressBook
	DB.Db.Find(&a)
	ctx.JSON(consts.StatusOK, utils.H{
		"message": a,
	})
}

func Delete(c context.Context, ctx *app.RequestContext) {
	var a AddressBook
	err := ctx.Bind(&a)
	if err != nil {
		ctx.JSON(consts.StatusInternalServerError, utils.H{
			"message": "参数错误",
		})
	}
	if DB.Db.Delete(&a).RowsAffected > 0 {
		ctx.JSON(consts.StatusOK, utils.H{
			"message": "ok",
		})
	} else {
		ctx.JSON(consts.StatusInternalServerError, utils.H{
			"message": "参数错误",
		})
	}
}

func Modify(c context.Context, ctx *app.RequestContext) {
	var a AddressBook
	err := ctx.Bind(&a)
	if err != nil {
		ctx.JSON(consts.StatusInternalServerError, utils.H{
			"message": "参数错误",
		})
	}
	if DB.Db.Save(&a).RowsAffected > 0 {
		ctx.JSON(consts.StatusOK, utils.H{
			"message": "ok",
		})
	} else {
		ctx.JSON(consts.StatusInternalServerError, utils.H{
			"message": "参数错误",
		})
	}
}
