//todo:以下创建服务器的方法四
package main

import (
	"fmt"
	"net/http"
)

func main() {
	myMux := http.NewServeMux()
	//向自己新建的myMux注册一个处理器
	myMux.HandleFunc("/aqie", func(rw http.ResponseWriter, w *http.Request) {
		fmt.Fprintln(rw, "这是采用的newServeMux的方式创建服务器")
	})
	http.ListenAndServe(":9090", myMux) //todo:这里就写自己的多路复用器myMux
}
