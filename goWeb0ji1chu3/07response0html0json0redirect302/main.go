package main

import (
	//"fmt"
	"encoding/json"
	"net/http"
)

//todo:向客户端返回页面，下面返回去的HTML会作为页面呈现，而且像其他的HTML一样是可用的页面
func myHandler(rw http.ResponseWriter, r *http.Request) {
	//rw.Write([]byte("这是来自服务器的字符串"))
	str := `<html>
    <head>
        <meta charset="UTF-8"/>
    </head>
    <body>
        <form action="http://localhost:9999/aqie" method="POST" enctype="application/x-www-form-urlencoded">
        序号：<input type="text" name="number"/><br/>
        信息:<input type="text" name="message"/><br/>
        密码：<input type="password" name="password"/><br/>
		<input type="submit"/><br/>
		这是服务器返回的原HTML文件
    </form>
    </body>
</html>`
	rw.Write([]byte(str))
}

//todo:向客户端返回json
func myHandler2(rw http.ResponseWriter, r *http.Request) {
	//rw.Header().Set("Content-Type","application/json")//可以不设定，rw.WriteHeader(int)会自动识别
	type student struct { //todo:注意，序列化结构体，类型名可以不大写，
		Name  string //todo:字段必须大写，否则不大写的字段会不进行json
		Age   int
		email string
	}
	a := student{Name: "jack", Age: 27, email: "123@sohu.com"}
	qiepian, err := json.Marshal(a)
	if err != nil {
		rw.Write([]byte("序列化失败"))
		return
	}
	rw.Write(qiepian) //todo:在客户端看，是在一个全新页面看到这里的qiepian代表的内容的
}

//todo:返回客户端，让客户端重定向到新的链接指向的页面
func myHandler3(w http.ResponseWriter, r *http.Request) {
	//todo:重定向方式一
	w.Header().Set("Location", "https://studygolang.com/pkgdoc")
	w.WriteHeader(302) //这个必须使用与重定向配套的状态码302,用错了比如304
	//w.Write([]byte("重定向"))//todo:没有这个w.Write()也没影响

	//todo:重定向方式二
	//http.Redirect(w,r,"https://studygolang.com/pkgdoc",302)
}

func main() {
	http.HandleFunc("/aqie", myHandler)
	http.HandleFunc("/oho", myHandler2)
	http.HandleFunc("/nani", myHandler3)
	http.ListenAndServe(":9999", nil)
}
