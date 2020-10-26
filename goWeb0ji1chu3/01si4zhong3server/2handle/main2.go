//todo:一下是创建服务器的方法二
package main

import (
	"fmt"
	"net/http"
)

type myHandler struct{}

func (m *myHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "这是用的实现ServeHTTP方法的最原始的方法")
}
func main() {
	meHandler := &myHandler{}
	http.Handle("/wo", meHandler)
	http.Handle("lala",&myHandler{})//todo:也可以这样用匿名的结构体变量
	http.ListenAndServe(":60", nil)
	//下面这种也可以，相当于不向defaultmux注册处理器了，直接固定
	//在该端口下固定调用这一个处理器
	// meHandler := &myHandler{}
	// http.ListenAndServe(":60",meHandler)
}
