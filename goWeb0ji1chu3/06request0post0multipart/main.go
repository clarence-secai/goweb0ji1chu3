package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func myHandler(rw http.ResponseWriter, r *http.Request) {
	myFile, myHeader, err := r.FormFile("fileupload") //括号内参数类似formvalue方法一样填HTML表单中的name=对应的属性
	if err != nil {
		fmt.Fprintln(rw, "便捷方式读取文件出错", err)
		return
	}
	file, _ := myHeader.Open()
	qiepian, _ := ioutil.ReadAll(file)
	fmt.Fprintln(rw, "上传的文件内容为", string(qiepian))

	//这上下两种作用是一样的，不同的是上面的是以页面的
	//方式返回给客户端，下面这种方式是以要客户端下载文
	//件的方式返回个客户端。二者同时运行的话，上面的
	//内容跟着下面的这个一起以文件形式返回给客户端
	can := make([]byte, 1024)
	myFile.Read(can)
	fmt.Fprintln(rw, "上传的文件内容是", string(can))

}
func main() {
	http.HandleFunc("/wo", myHandler)
	http.ListenAndServe(":8899", nil)
}
