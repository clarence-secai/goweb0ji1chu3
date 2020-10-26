package main

import (
	"fmt"
	"html/template"
	"net/http"
)
func page(rw http.ResponseWriter, w *http.Request) {
	t, _ := template.ParseFiles("./a.html")
	t.Execute(rw, "hello!")
}
//该网站的该端口给客户端返回设置好cookie
func cookieHandler(rw http.ResponseWriter, w *http.Request) {
	mycookie := http.Cookie{
		Name:   "ticket",
		Value:  "chifan",
		MaxAge: 120, //todo:该参数如未设置，关闭浏览器则cookie失效(是关闭浏览器，不是退出该网站)，
		//如设置了该参数，从服务器最近一次给客户端cookie算起，不
		//论是否期间关闭浏览器（注意再次进入该网站服务器再次给客户
		//端发新cookie的情形)  到达该时间后失效  该字段单位是秒
		HttpOnly: true,
	}

	hercookie := http.Cookie{Name: "mary", Value: "haha"}

	//todo:方法一如下
	//http.SetCookie(rw,&hercookie)
	////todo:上一行为在rw的响应头域中添加Set-Cookie为键mycookie为值，无后面的t.Execute(rw, user)
	////也已添加到浏览器客户端上了。下面方法二也是同理

	//todo:方法二如下(和上一种方式同时用时，会打架导致上一种失效)
	// 采用下面这种方式设置cookie时，key必须是"set-cookie",开头字母可以不分大小写
	rw.Header().Set("set-cookie", mycookie.String())
	rw.Header().Add("set-cookie", hercookie.String())//追加一个cookie
}

func getcookie(rw http.ResponseWriter, r *http.Request) {
	//todo:方式一如下：
	str := r.Header.Get("Cookie")//todo:键必须是"Cookie"
	fmt.Fprintln(rw, "之前服务器收到的cookie是", str)

	//todo:方式二如下：
	fmt.Fprintln(rw, "我是分割线")
	fmt.Fprintln(rw, r.Cookies())
	//返回一个cookie内容的切片，也可r.Cookie(name string)，返回的就是相应的cookie
}
func main() {
	http.HandleFunc("/a", page)
	http.HandleFunc("/binggan", cookieHandler)
	http.HandleFunc("/b", getcookie)

	http.ListenAndServe(":6789", nil)
}
