package main

import (
	"fmt"
	"log"
)

func main() {
	client := &WebsocketClient{
		Msg: func(ws *WebsocketClient, messageType int, bts []byte) {
			fmt.Printf("%s\n", bts)
		},
	}
	addr := "ws://127.0.0.1:9630/ws"
	err := client.Dial(addr)
	if err != nil {
		log.Fatalln(err.Error())
	} else {
		select {}
	}
}
