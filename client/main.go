package main

import (
	"log"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"

	"github.com/gorilla/websocket"
	"github.com/mem/ws-rpc/internal/common"
)

func main() {
	ws, res, err := dialer.Dial("ws://127.0.0.1:8010/ws", http.Header{})
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	defer ws.Close()

	log.Println(res)

	handle(ws)
}

var dialer = websocket.Dialer{
	ReadBufferSize:  common.MaxMessageSize,
	WriteBufferSize: common.MaxMessageSize,
}

func handle(ws *websocket.Conn) {
	defer func() {
		ws.Close()
	}()

	rwc := &common.ReadWriteCloser{WS: ws}
	codec := jsonrpc.NewClientCodec(rwc)
	c := rpc.NewClientWithCodec(codec)
	// c := rpc.NewClient(rwc)

	for {
		args := &common.HelloArgs{Msg: "Hello, World"}
		var reply common.HelloReply
		err := c.Call("Comm.Hello", args, &reply)
		if err != nil {
			log.Printf("%v", err)
			break
		}
		log.Printf("%v", reply)
	}
}
