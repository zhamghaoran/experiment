package service

import (
	"context"
	"encoding/json"
	"experiment/3_cs/client/client"
	"experiment/3_cs/server/kitex_gen/kitex_gen/service"
	"fmt"
)

func afterOperation() {
	fmt.Println("按任意键继续")
	var a string
	fmt.Scan(&a)
}

func clear() {
	for i := 1; i <= 100; i++ {
		fmt.Println("")
	}
}
func ShowAll() {
	clear()
	defer clear()
	defer afterOperation()
	all, err := client.C.Select(context.Background(), &service.AddReq{})
	if err != nil {
		return
	}
	for _, v := range all.AddReqs {
		b, _ := json.Marshal(v)
		fmt.Println(string(b))
	}
}

func Add() {
	clear()
	defer afterOperation()
	defer clear()
	res := &service.AddReq{}
	fmt.Println("姓名：")
	fmt.Scan(&res.Name)
	fmt.Println("电话：")
	fmt.Scan(&res.Tel)
	fmt.Println("地址")
	fmt.Scan(&res.Add)
	ans, _ := client.C.Add(context.Background(), res)
	if ans.Code > 0 {
		fmt.Println("添加成功")
	} else {
		fmt.Println("添加")
	}
}
func Modify() {
	clear()
	defer clear()
	defer afterOperation()
	res := &service.AddReq{}
	fmt.Println("请输入ID")
	fmt.Scan(&res.Id)
	fmt.Println("姓名：")
	fmt.Scan(&res.Name)
	fmt.Println("电话")
	fmt.Scan(&res.Tel)
	fmt.Println("地址")
	fmt.Scan(&res.Add)
	ans, _ := client.C.Modify(context.Background(), res)
	if ans.Code > 0 {
		fmt.Println("更新成功")
	} else {
		fmt.Println("更新失败")
	}
}
func Delete() {
	clear()
	defer clear()
	defer afterOperation()
	res := &service.AddReq{}
	fmt.Println("请输入ID")
	fmt.Scan(&res.Id)
	s, _ := client.C.Delete(context.Background(), res)
	if s.Code > 0 {
		fmt.Println("删除成功")
	} else {
		fmt.Println("删除失败")
	}
}
