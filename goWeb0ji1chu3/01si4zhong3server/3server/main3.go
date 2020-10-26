//todo:一下是创建服务器方法三
package main

import (
	"fmt"
	"net/http"
	"time"
)

type myHandler3 struct{}

func (m *myHandler3) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "这是用的手动配置服务器server参数的方法")
}

func main() {
	meHandler3 := &myHandler3{}
	myServer := &http.Server{
		Addr:        ":90",
		Handler:     meHandler3, //todo:这里填nil就会调用http.DefaultServeMux
		ReadTimeout: time.Second * 2}
	myServer.ListenAndServe() //这种方式的缺点似乎是不能细分网址+路径
}
