package main

import (
	//"fmt"
	//"time"
	"html/template"
	"net/http"
)

func myHandler(rw http.ResponseWriter, r *http.Request) {
	//todo:这里参数写与本主函数文件平级文件夹(路径需含该平级文件夹)路径下的该文件名，
	// 如和本主函数文件在同一个文件夹下，写文件名即可

	t, _ := template.ParseFiles("muban return.html")
	t.Execute(rw, "我是服务端代码去和{{.}}配合的字符串")
}

type bowl struct{}

func (b *bowl) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	t2 := template.Must(template.ParseFiles("muban2return.html", "muban3return.html"))
	t2.ExecuteTemplate(rw, "muban2return.html", "我是模板2HTML")
	//time.Sleep(time.Second*2)//todo:这一句没必要，上下两个执行可以返回在同一个服务器返回的页面中
	t2.ExecuteTemplate(rw, "muban3return.html", "我是模板3HTML")

}
func main() {
	http.HandleFunc("/yi", myHandler)
	me := &bowl{}
	http.Handle("/ei", me)
	http.ListenAndServe(":5555", nil)
}
