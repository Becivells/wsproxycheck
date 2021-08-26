package main

import (
	_ "embed"
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "0.0.0.0:8083", "http service address")

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1,
	WriteBufferSize: 1,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
} // use default options

//go:embed index.html
var indexPage []byte

func index(w http.ResponseWriter, r *http.Request) {
	w.Write(indexPage)
}

func echo(w http.ResponseWriter, r *http.Request) {
	//接收客户端发送的origin （重要！）
	w.Header().Set("Access-Control-Allow-Origin", "*")
	//服务器支持的所有跨域请求的方法
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
	//允许跨域设置可以返回其他子段，可以自定义字段
	w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session")
	// 允许浏览器（客户端）可以解析的头部 （重要）
	w.Header().Set("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
	//设置缓存时间
	w.Header().Set("Access-Control-Max-Age", "172800")
	//允许客户端传递校验信息比如 cookie (重要)
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	c, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	c.WriteMessage(1, []byte("aaa"))
	c.SetReadLimit(1)
	for {

		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
	}
}

func main() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/", index)
	http.HandleFunc("/echo", echo)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
