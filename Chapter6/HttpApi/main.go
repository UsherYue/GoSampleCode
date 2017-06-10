package main

import (
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
	url, _ := url.Parse("http://127.0.0.1:8000")
	httputil.NewSingleHostReverseProxy(url).ServeHTTP(w, r)
}

func main() {
	//初始化多路复用器
	Mux.HandleFunc("/user/age", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		w.Write([]byte(`{"Age":"18"}`))
	})
	Mux.HandleFunc("/user/name", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		w.Write([]byte(`{"Name":"张三"}`))
	})
	Mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`hello,api!`))
	})
	//单独启动一个协程启动后端http服务 端口8000
	go func() {
		server := &http.Server{
			Addr:    ":8000",
			Handler: &HttpHandler{},
		}
		server.ListenAndServe()
	}()

	//反向代理upstream 80端口
	http.ListenAndServe(":8888", &ProxyHander{})

}
