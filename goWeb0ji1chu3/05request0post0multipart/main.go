//todo:注意，如表单中有上传文件的功能，则HTML中的enctype要改成enctype="multipart/form-data"
// 也正是由于由application/x-www-form-urlencoded改成了multipart/form-data，故需要
// 换成使用r.MultipartForm或r.FormFile()来获得相应表单和上传文件的内容。如果还用原来
// 的r.Form和r.PostForm就取不到任何东西了。
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	//"io"
	//"io/ioutil"
)

func myHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "客户端发来的协议版本", r.Proto)        //客户端发来的协议版本 HTTP/1.1
	fmt.Fprintln(rw, "客户端发来内容长度", r.ContentLength) //todo:客户端发来请求体的内容长度 418

	r.ParseForm()
	fmt.Fprintln(rw, "ho=", r.Form)     //ho= map[]   即取不到任何东西
	fmt.Fprintln(rw, "ha=", r.PostForm) //ha= map[]   即取不到任何东西

	r.ParseMultipartForm(1024)
	fmt.Fprintln(rw, "文件MultipartForm内容是", r.MultipartForm) //这个会打开表单内容和上传的文件内容
	//文件MultipartForm内容是 &{map[mima:[123] yonghuming:[jack]] map[wenjian:[0xc0000d6190]]}
	fmt.Fprint(rw, "-value-: ", r.MultipartForm.Value) //-value-: map[mima:[123] yonghuming:[jack]]
	fmt.Fprint(rw, "-file-: ", r.MultipartForm.File)   //-file-: map[wenjian:[0xc0000d6190]]

	myFile, err := r.MultipartForm.File["wenjian"][0].Open()
	if err != nil {
		fmt.Fprintln(rw, "打开文件出错", err)
		return
	}
	//can := make([]byte,1024)
	//myFile.Read(can)
	//fmt.Fprintln(rw,"上传的文件内容是",string(can))
	//上面这种方式会将所有信息(包括非文件的表单中的内容)返回到一个文件中，要求客户端下载才能看
	//下面这种跟讲义一样的方法则是直接以页面的形式返回给客户端
	//原因见multipart包中对type File的说明

	qiepian, err := ioutil.ReadAll(myFile)
	if err != nil {
		fmt.Fprintln(rw, "读取文件出错", err)
		return
	}
	fmt.Fprintln(rw, "上传的文件内容是：", string(qiepian)) //上传的文件内容是:我就随便写写东西this is a txt file .

}
func main() {
	http.HandleFunc("/ha", myHandler)
	http.ListenAndServe(":8888", nil)
}
