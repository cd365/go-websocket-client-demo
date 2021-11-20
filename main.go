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
	addr := "ws://host:port/addr"
	err := client.Dial(addr)
	if err != nil {
		log.Fatalln(err.Error())
	} else {
		select {}
	}
}
