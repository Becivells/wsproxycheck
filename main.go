package main

import (
	_ "embed"
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"os"
)

var (
	Version = "unknown"
	Commit  = "unknown"
	Date    = "unknown"
	Branch  = "unknown"
)
var addr = flag.String("addr", "0.0.0.0:8083", "http service address")
var sv = flag.Bool("v", false, "show version")

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1,
	WriteBufferSize: 1,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
} // use default options

func showVersion() {
	fmt.Printf("Current Version: %s\n", Version)
	fmt.Printf("Current branch: %s\n", Branch)
	fmt.Printf("Current commit: %s\n", Commit)
	fmt.Printf("Current date: %s\n", Date)
	os.Exit(0)
}

//go:embed index.html
var indexPage []byte

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "no-cache")
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
	//允许客户端传递校验信息比如 cookie (重要)
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Cache-Control", "no-cache")
	c, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	c.WriteMessage(1, []byte(""))
	c.SetReadLimit(1)

	for {

		_, message, err := c.ReadMessage()
		//time.Sleep(time.Duration(1)*time.Second)
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
	}
}

func main() {
	flag.Parse()
	if *sv {
		showVersion()
	}
	log.SetFlags(0)
	http.HandleFunc("/", index)
	http.HandleFunc("/echo", echo)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
