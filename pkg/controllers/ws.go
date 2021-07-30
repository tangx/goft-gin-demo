package controllers

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/shenyisyn/goft-gin/goft"
)

type WebSocket struct {
	hset      sync.Map    // 保存客户端
	broadcast chan []byte // 广播消息队列
}

// NewWebSocket 初始化
func NewWebSocket() *WebSocket {
	return &WebSocket{
		broadcast: make(chan []byte, 100),
	}
}

func (ws *WebSocket) Name() string {
	return "WebSocket"
}

func (ws *WebSocket) Build(goft *goft.Goft) {
	goft.Handle("GET", "/ws", ws.handlerWS)
}

func (ws *WebSocket) handlerWS(c *gin.Context) goft.Void {

	// 后台轮训广播
	go ws.Broadcast()

	// 初始化 http -> ws 协议升级
	wsu := &websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			// 允许所有跨域
			// 需要根据实际情况更改
			return true
		},
	}

	// 创建链接
	// 	这里与一般的 tcp server 不同， 不需要使用使用 for 循环等待新链接
	// 	http 框架本身已经实现新链接的等待
	conn, err := wsu.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("connect failed:%s\n", err)
	}

	ws.ConnLifeCycle(conn)

	return goft.Void{}
}

func (ws *WebSocket) StoreConn(conn *websocket.Conn) {
	key := wsConnKey(conn)
	ws.hset.Store(key, conn)
	log.Printf("info: %s connect\n", key)
}

func (ws *WebSocket) RemoveConn(conn *websocket.Conn) {

	key := wsConnKey(conn)
	ws.hset.Delete(key)
	conn.Close()
	log.Printf("info: %s disconnect\n", key)
}

func (ws *WebSocket) ConnLifeCycle(conn *websocket.Conn) {
	ws.StoreConn(conn)
	// 退出时删除保存链接
	defer ws.RemoveConn(conn)

	key := wsConnKey(conn)

	for {
		mt, r, err := conn.NextReader()
		if err != nil {

			if err != io.EOF {
				log.Println("NextReader failed:", err)
			}

			// 退出
			return
		}

		// 读取消息
		body, err := ioutil.ReadAll(r)
		if err != nil {
			log.Printf("read message failed: %v\n", err)
			continue
		}

		// 将消息广播到所有用户
		if mt == websocket.TextMessage {
			fmt.Printf("%s ->: %s\n", key, body)

			ws.broadcast <- body
		}
	}
}

// Broadcast 向所有用户广播消息
func (ws *WebSocket) Broadcast() {

	// 生成满足 sync.Map Range 的函数
	fnGen := func(msg []byte) func(k, v interface{}) bool {
		return func(k, v interface{}) bool {
			// 断言 k,v 类型
			_k := k.(string)
			_conn := v.(*websocket.Conn)

			msg := fmt.Sprintf("%s <- : %s", _k, msg)

			// 发送消息
			err := _conn.WriteMessage(websocket.TextMessage, []byte(msg))
			if err != nil {
				fmt.Printf("Send msg to failed: %v\n", err)
			}

			return true
		}
	}

	for {
		// 从广播消息队列中获取消息
		msg := <-ws.broadcast
		fn := fnGen(msg)
		ws.hset.Range(fn)
	}
}

func wsConnKey(conn *websocket.Conn) string {
	return conn.RemoteAddr().String()
}
