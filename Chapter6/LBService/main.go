package main

import (
	"math/rand"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var Mux = http.NewServeMux()

//Http处理
type HttpHandler struct{}

func (this *HttpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	Mux.ServeHTTP(w, r)
}

//反向代理处理
type ProxyHander struct{}

func (this *ProxyHander) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	serverList := []string{
		"http://127.0.0.1:8001",
		"http://127.0.0.1:8002",
		"http://127.0.0.1:8003",
	}
	url, _ := url.Parse(serverList[rand.Intn(3)])
	//确保后端服务器能收到真实的HOST
	r.Host = ""
	httputil.NewSingleHostReverseProxy(url).ServeHTTP(w, r)

}

func main() {
	//初始化多路复用器
	Mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`hello,api! <br/>Host:` + r.Host))
	})
	//单独启动一个协程启动后端http服务 端口8000
	go func() {
		server := &http.Server{
			Addr:    ":8001",
			Handler: &HttpHandler{},
		}
		server.ListenAndServe()
	}()

	//单独启动一个协程启动后端http服务 端口8000
	go func() {
		server := &http.Server{
			Addr:    ":8002",
			Handler: &HttpHandler{},
		}
		server.ListenAndServe()
	}()

	//单独启动一个协程启动后端http服务 端口8000
	go func() {
		server := &http.Server{
			Addr:    ":8003",
			Handler: &HttpHandler{},
		}
		server.ListenAndServe()
	}()

	//反向代理upstream 80端口
	http.ListenAndServe(":8888", &ProxyHander{})

}
