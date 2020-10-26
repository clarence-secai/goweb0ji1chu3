//todo:下面这种方式因繁琐而不常用，1handleFunc文件夹中的是最常用的方式
// 另外可以看看其他一共四种创建服务器的方式。
package main

// func myhandler(rw http.ResponseWriter,r *http.Request){
// 	fmt.Fprintln(rw,"hello,world")
// }
//type meHandler struct{}
//func (m *meHandler)ServeHTTP(rw http.ResponseWriter,r *http.Request){
//	io.WriteString(rw,"你好，世界")
//}
//
//func main(){
//	//http.HandleFunc("/",myhandler)
//	woHandler := &meHandler{}
//	http.Handle("/haha",woHandler)//少了该行代码，跑起来会啥也没有就结束
//	http.ListenAndServe(":8080",woHandler)//以上两行代码不能互换顺序
////这里括号内本来是填nil或mux 也可以直接填处理器handler，像上面这样
//}
