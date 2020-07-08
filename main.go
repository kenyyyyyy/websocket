package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
)

// 全部连接
var Clients =make([]*websocket.Conn,0)
var lock sync.Mutex

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
// 加入房间
func join(w http.ResponseWriter, r *http.Request) {
	// 升级协议
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	Clients=append(Clients,conn)

	w.WriteHeader(200)
	fmt.Println(len(Clients))
	for  {
		if _,data,err:=conn.ReadMessage();err==nil{
			write(data)
		}else{
			return
		}
	}
}

func write(msg []byte)  {
	for _,c:=range Clients{
		lock.Lock()
		err:=c.WriteMessage(websocket.TextMessage,msg)
		if err!=nil{
			fmt.Println(err)
		}
		lock.Unlock()
	}
}

func main() {
	//
	http.HandleFunc("/ws", join)

	fmt.Println("Starting v1 server ...")
	//启动
	http.ListenAndServe(":8080", nil)
}
