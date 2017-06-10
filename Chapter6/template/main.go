package main

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"os"
)

type UserInfo struct {
	Name string
	Age  int
}

type TemplateData struct {
	Title string
	Infos []UserInfo
}

func testTemplate() {
	data := &TemplateData{Title: "this is a go template test"}
	for i := 0; i < 100; i++ {
		//随机创建100 一个UserInfo  每个人的年龄都是随机
		data.Infos = append(data.Infos, UserInfo{fmt.Sprintf("User%d", i+1), rand.Intn(20) + 10})
	}
	t, err := template.ParseFiles("tpl/index.html")
	if err != nil {
		fmt.Println(err.Error())
		log.Fatalln("模板文件解析错误")
	}
	t.Execute(os.Stdout, data)
}

type HttpHandler struct{}

func (this *HttpHandler) ServeHTTP(wr http.ResponseWriter, r *http.Request) {
	data := &TemplateData{Title: "this is a go template test"}
	for i := 0; i < 100; i++ {
		//随机创建100 一个UserInfo  每个人的年龄都是随机
		data.Infos = append(data.Infos, UserInfo{fmt.Sprintf("User%d", i+1), rand.Intn(20) + 10})
	}
	t, err := template.ParseFiles("tpl/index.html")
	if err != nil {
		fmt.Println(err.Error())
		log.Fatalln("模板文件解析错误")
	}
	t.Execute(wr, data)
}

func main() {
	server := &http.Server{Addr: ":8888"}
	serverMux := http.NewServeMux()
	serverMux.Handle("/info", &HttpHandler{})
	//server.Handler = &HttpHandler{}
	//将handler设置成多路复用器
	server.Handler = serverMux
	server.ListenAndServe()
}
