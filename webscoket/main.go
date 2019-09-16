package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

type Ping struct {
	Ping int `json:"ping"`
}

type Pong struct {
	Pong int `json:"pong"`
}

var addr = flag.String("addr", "api.huobi.pro", "http service address")

func main() {
	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "wss", Host: *addr, Path: "/ws"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})
	in := make(chan Ping)

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			r, err := gzip.NewReader(bytes.NewReader(message))
			if err != nil {
				log.Println("gzip err:", err)
				return
			}

			d := json.NewDecoder(r)
			p := Ping{}
			d.Decode(&p)
			log.Printf("recv: %d", p.Ping)
			in <- p
			r.Close()
		}
	}()

	for {
		select {
		case <-done:
			return
		case msg := <-in:
			pong := Pong{
				Pong: msg.Ping,
			}
			sengmsg, _ := json.Marshal(pong)
			var zBuf bytes.Buffer
			zw := gzip.NewWriter(&zBuf)
			if _, err = zw.Write(sengmsg); err != nil {
				fmt.Println("-----gzip is faild,err:", err)
			}
			zw.Close()
			log.Printf("send: %s", sengmsg)
			err := c.WriteMessage(websocket.TextMessage, zBuf.Bytes())
			if err != nil {
				log.Println("write:", err)
				return
			}
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
