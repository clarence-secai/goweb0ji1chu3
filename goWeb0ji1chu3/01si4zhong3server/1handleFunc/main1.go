//todo:以下是创建服务器的方法一  这种方式最常用
package main

import (
	"fmt"
	//"fmt"
	"io"
	"net/http"
)

func myHandler2(rw http.ResponseWriter, r *http.Request) {
	io.WriteString(rw, "我用的io包，使用的handlefunc方法创建服务器")
}
func main() {
	http.HandleFunc("/haha", myHandler2)

	//处理器也可以是符合签名的匿名函数
	http.HandleFunc("/hoho", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("我用的rw.Write,我是hoho"))
	})
	http.HandleFunc("/ei", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw,"我用的fmt.Fprint，我的ei")
	})
	http.ListenAndServe("http://192.168.1.106:70", nil)
	//http.ListenAndServe("",nil)  //这种是默认本机且是80端口
	//网址可以略填为":70",注意网址最多只能写到端口号，后面的不能再写
	//nil此时不能换填为myHandler 因为类型不匹配
}
