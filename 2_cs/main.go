package main

import (
	"experiment/2_cs/service"
	"fmt"
)

func main() {
	display()
}

func display() {
	for {
		fmt.Println("------------")
		fmt.Println("1 查看联系人信息")
		fmt.Println("2 添加新联系人")
		fmt.Println("3 修改联系人")
		fmt.Println("4 删除联系人")
		fmt.Println("------------")
		var x int
		fmt.Scan(&x)
		switch x {
		case 1:
			service.ShowAll()
		case 2:
			service.Add()
		case 3:
			service.Modify()
		case 4:
			service.Delete()
		}
	}
}
