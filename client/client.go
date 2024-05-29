package main

import (
	"fmt"
	"log"
	"net/rpc"
	"net/rpc/jsonrpc"

	"github.com/khusainnov/newbee/app/specs"
)

func main() {
	client, _ := jsonrpc.Dial("tcp", ":8001")

	echo(client)
}

func echo(client *rpc.Client) {
	send := &specs.EchoReq{Message: "Hello JSON-RPC"}

	var resp specs.Resp
	err := client.Call("API.Echo", send, &resp)
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("\n%v\n", resp.Message)
}
