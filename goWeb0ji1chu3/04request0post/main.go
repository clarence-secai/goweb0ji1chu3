package main

import (
	"fmt"
	"net/http"
	//"io"
)

func myHandler(rw http.ResponseWriter, r *http.Request) {
	/*todo:获取url里的详细信息 详细的可以见文档ner/url包下 Type Url结构体的解释*/
	fmt.Fprintln(rw, "客户端发来时请求主机是", r.Host)        //客户端发来时请求主机是 localhost:8888
	fmt.Fprintln(rw, "客户端发来时方法是", r.Method)        //客户端发来时方法是 POST
	fmt.Fprintln(rw, "客户端发来时url是", r.URL)          //客户端发来时url是 /ha
	fmt.Fprintln(rw, "客户端发来时url是", r.URL.String()) //客户端发来时url是 /ha
	fmt.Fprintln(rw, "客户端发来时url是", r.URL.Scheme)   //客户端发来时url是
	fmt.Fprintln(rw, "客户端发来时url是", r.URL.RawQuery) //客户端发来时url是

	/*todo:获取请求头的详细信息*/
	fmt.Fprintln(rw, "客户端发来时请求头是", r.Header)
	//客户端发来时请求头是 map[Accept:[text/html,application/xhtml+xml,application/xml;
	//q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9]
	//Accept-Encoding:[gzip, deflate, br]
	//Accept-Language:[zh-CN,zh;q=0.9,en;q=0.8, en-GB;q=0.7,en-US;q=0.6]
	//Cache-Control:[max-age=0]
	//Connection:[keep-alive]
	//Content-Length:[24]
	//Content-Type:[application/x-www-form-urlencoded]
	//Origin:[null]
	//Sec-Fetch-Dest:[document]
	//Sec-Fetch-Mode:[navigate]
	//Sec-Fetch-Site:[cross-site] Sec-Fetch-User:[?1]
	//Upgrade-Insecure-Requests:[1]
	//User-Agent:[Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36 Edg/84.0.522.63]]

	fmt.Fprintln(rw, "客户端发来时请求头语言是", r.Header["Accept-Language"])
	//客户端发来时请求头语言是 [zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6]
	fmt.Fprintln(rw, "客户端发来时请求头接受编码是", r.Header.Get("Accept-Encoding"))
	//客户端发来时请求头接受编码是 gzip, deflate, br

	/*todo:获取请求体的全部信息*/
	//lenth := r.ContentLength   //todo:ContentLength记录请求体的内容的字节长度
	//mybody := make([]byte,lenth)
	////r.Body.Read(mybody) //不能下面接收返回值不能去判断是否是err != nil，否则报读取请求体失败,详见io包Reader接口的解释
	//n,err := r.Body.Read(mybody)
	//if err != io.EOF {
	//	fmt.Fprintln(rw,"读取请求体失败",err,n)
	//	return
	//}
	//fmt.Fprintln(rw,"客户端发来时请求体是",string(mybody))//客户端发来时请求体是 yonghuming=jack&mima=123

	//注意，上面的读取请求体和下面的获得表单参数不能同时都运行，不然下面的代码输出的内容都是空的，对调顺序则输出请请求体的代码输出的内容也是乱码

	r.ParseForm()
	fmt.Fprintln(rw, "表单用户名参数为", r.Form)                //表单用户名参数为 map[mima:[123] yonghuming:[jack]]
	fmt.Fprintln(rw, "表单用户名参数为", r.Form["yonghuming"])  //表单用户名参数为 [jack]
	fmt.Fprintln(rw, "表单密码参数为", r.Form.Get("mima"))     //表单密码参数为 123
	fmt.Fprintln(rw, "表单密码参数为", r.PostForm["mima"])     //表单密码参数为 [123]
	fmt.Fprintln(rw, "表单密码参数为", r.PostForm.Get("mima")) //表单密码参数为 123

	//todo:下面这种方式是获取表单参数的最快方式，可以和上面的同时运行
	fmt.Fprintln(rw, "表单的用户名是", r.FormValue("yonghuming")) //表单的用户名是 jack
	fmt.Fprintln(rw, "表单的密码是", r.PostFormValue("mima"))    //表单的密码是 123
}
func main() {
	http.HandleFunc("/ha", myHandler)
	http.ListenAndServe(":8888", nil)
}
