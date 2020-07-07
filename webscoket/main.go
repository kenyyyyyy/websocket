package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"strconv"
	"sync"
)

type WebSocket struct {
	conn  *websocket.Conn
	urlId string
	Id int
}

var wait sync.WaitGroup
var origin = "http://127.0.0.1:8080/"
var url = "ws://127.0.0.1:8080/join"

func NewWebSocket(Id int) (ws *WebSocket) {
	s := strconv.Itoa(Id)
	ws = &WebSocket{
		urlId: url + "?Id=" + s,
		Id: Id,
	}
	return
}
func (w *WebSocket) GetConn() (err error) {
	conn, err := websocket.Dial(w.urlId, "", origin)
	if err != nil {
		fmt.Println("建立连接失败:",err)
		return
	}
	w.conn = conn
	fmt.Println("建立连接")
	return
}

// 接收数据
func (w *WebSocket) Read() {
	fmt.Println("准备接收数据:")
	if w.conn == nil {
		fmt.Println("建立连接")
	}
	msg := make([]byte, 512)
	for  {
		n, err := w.conn.Read(msg)
		if err != nil {
			fmt.Println("接收数据失败:", err)
			wait.Done()
		}
		fmt.Printf("成功接收数据:%s，ID为:%d: \n",string(msg[:n]),w.Id)
	}

}
func main() {
	wait.Add(20)
	for i := 0; i < 20; i++ {
		go func(i int) {
			WSconn := NewWebSocket(i)
			WSconn.GetConn()
			WSconn.Read()
		}(i)
	}
	wait.Wait()
}
