package service

import (
	"encoding/json"
	"experiment/2_cs/DB"
	"fmt"
)

type AddressBook struct {
	ID      int    `gorm:"id" json:"id" `
	Name    string `gorm:"name" json:"姓名"`
	Tel     int64  `gorm:"tel" json:"电话"`
	Address string `gorm:"address" json:"地址"`
}

func (AddressBook) TableName() string {
	return "address_book"
}

func afterOperation() {
	fmt.Println("按任意键继续")
	var a string
	fmt.Scan(&a)
}
func SelectAllFromDB() {
	var add []AddressBook
	defer afterOperation()
	DB.Db.Find(&add)
	for _, v := range add {
		by, _ := json.Marshal(v)
		fmt.Println(string(by))
	}

}

func clear() {
	for i := 1; i <= 100; i++ {
		fmt.Println("")
	}
}
func ShowAll() {
	clear()
	defer clear()
	SelectAllFromDB()
}

func Add() {
	clear()
	defer afterOperation()
	defer clear()
	var add AddressBook
	fmt.Println("姓名：")
	fmt.Scan(&add.Name)
	fmt.Println("电话：")
	fmt.Scan(&add.Tel)
	fmt.Println("地址")
	fmt.Scan(&add.Address)
	affected := DB.Db.Create(&add).RowsAffected
	if affected > 0 {
		fmt.Println("添加成功")
	} else {
		fmt.Println("添加")
	}
}
func Modify() {
	clear()
	defer clear()
	defer afterOperation()
	var ad AddressBook
	fmt.Println("请输入ID")
	fmt.Scan(&ad.ID)
	fmt.Println("姓名：")
	fmt.Scan(&ad.Name)
	fmt.Println("电话")
	fmt.Scan(&ad.Tel)
	fmt.Println("地址")
	fmt.Scan(&ad.Address)
	affected := DB.Db.Save(&ad).RowsAffected
	if affected > 0 {
		fmt.Println("更新成功")
	} else {
		fmt.Println("更新失败")
	}
}
func Delete() {
	clear()
	defer clear()
	defer afterOperation()
	var add AddressBook
	fmt.Println("请输入ID")
	fmt.Scan(&add.ID)
	affected := DB.Db.Delete(&add).RowsAffected
	if affected > 0 {
		fmt.Println("删除成功")
	} else {
		fmt.Println("删除失败")
	}
}
