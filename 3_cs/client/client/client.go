package client

import (
	"experiment/3_cs/server/kitex_gen/kitex_gen/service/address"
	"github.com/cloudwego/kitex/client"
)

var C address.Client

func init() {
	newClient, err := address.NewClient("hello", client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		panic(err)
	}
	C = newClient
}
