package main

import (
	service "experiment/3_cs/server/kitex_gen/kitex_gen/service/address"
	"log"
)

func main() {
	svr := service.NewServer(new(AddressImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
