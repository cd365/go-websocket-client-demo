package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

// dialer websocket dialer
var dialer = &websocket.Dialer{
	Proxy:            http.ProxyFromEnvironment,
	HandshakeTimeout: 15 * time.Second,
}

type WebsocketClient struct {
	addr string
	conn *websocket.Conn
	Msg  func(ws *WebsocketClient, messageType int, bts []byte)
}

func (s *WebsocketClient) Dial(addr string) (err error) {
	s.addr = addr
	s.conn, _, err = dialer.Dial(s.addr, nil)
	if err != nil {
		return
	}
	go s.read()
	return
}

func (s *WebsocketClient) read() {
	var err error
	var messageType int
	var bts []byte
	for {
		messageType, bts, err = s.conn.ReadMessage()
		if err != nil {
			log.Printf("websocket read error: %s", err.Error())
			s.reconnect()
			continue
		}
		s.Msg(s, messageType, bts)
	}
}

func (s *WebsocketClient) reconnect() {
	var err error
	// close the connection with a reading error
	_ = s.conn.Close()
	timer := time.NewTimer(time.Second)
	for range timer.C {
		s.conn, _, err = dialer.Dial(s.addr, nil)
		if err != nil {
			log.Printf("websocket reconnect error: %s", err.Error())
			timer.Reset(time.Second * 3)
			continue
		}
		timer.Stop()
		log.Println("websocket reconnect success")
		break
	}
}
