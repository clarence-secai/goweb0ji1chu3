package main

import (
	//"fmt"
	"go2web/goWeb0ji1chu3/200228liu-dong4zuo4/htmlfile/range"
	"go2web/goWeb0ji1chu3/200228liu-dong4zuo4/htmlfile/with"
	"html/template"
	"net/http"
)

func ifhandler(rw http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./htmlfile/if/if.html"))
	a := false
	t.Execute(rw, a)
	b := true //如果这里填别的，比如4、"字符串"等其他类型，发现也会当true处理
	t.Execute(rw, b)
}

func rangehandler(rw http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./htmlfile/range/range.html")
	t.Execute(rw, myrange.School())
}
func rangeMapHandler(rw http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./htmlfile/range/range2.html")
	t.Execute(rw, myrange.Usermap())
}

func withhandler(rw http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./htmlfile/with/with.html")
	t.Execute(rw, "太子")
}
func withhandler2(rw http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./htmlfile/with/with2.html")
	t.Execute(rw, with.Tec())
}

func templateHandler(rw http.ResponseWriter, r *http.Request) {
	//todo:想使用模板之间的包含，则两个模板文件都需要解析,且被包含的模板要作为第二个参数
	t := template.Must(template.ParseFiles("./htmlfile/template/template1.html", "./htmlfile/template/template2.html"))
	t.Execute(rw, "后台数据") //这里不指名操作哪个模板文件，则默认操作上一行解析的文件中的第一个
} //注意，这两个文件里的HTML代码不能互相包含，否则在浏览器端看到的就是无限循环互相包含

func definehandler(rw http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./htmlfile/define/define.html")
	t.Execute(rw, "我是后台数据")
	t.ExecuteTemplate(rw, "muban1", "我是后台数据")
	//todo:没有这一句，html文件中define muban1那部分内容就不会在浏览器显示！！！
	// define的主要用途是供其他HTML通过template关键词进行引用
	//这两句会分别将数据传给内外层两个模板，但如果两句顺序颠倒，会出现奇怪的现象
	// 因此最妥当的办法是只执行一个文件里的某一个模板，或按照由外到内的模板
	// 嵌套顺序来有序执行
}

func block1(rw http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./htmlfile/block/page1.html", "./htmlfile/block/page2.html") //
	t.ExecuteTemplate(rw, "page1.html", "我是后台数据")
}
func block2(rw http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./htmlfile/block/page1.html", "./htmlfile/block/page2.html") //
	t.ExecuteTemplate(rw, "page2.html", "我是后台数据")
}
//todo:注意block中的page2.HTML不能写上一行的./htmlfile/block/page2.html，只需写模板文件名称即可。
// 另外，试试执行page1.html，会发现自定义define的模板不会随着一起被显示,除非像definehandler处理
// 器那样，因此自定义的模板应统一放在一个文件中，方便其他文件用template来进行包含取用

func main() {
	http.HandleFunc("/if", ifhandler)
	http.HandleFunc("/range", rangehandler)
	http.HandleFunc("/rangemymap", rangeMapHandler)
	http.HandleFunc("/with", withhandler)
	http.HandleFunc("/with2", withhandler2)
	http.HandleFunc("/template", templateHandler)
	http.HandleFunc("/define", definehandler)
	http.HandleFunc("/block1", block1)
	http.HandleFunc("/block2", block2)

	http.ListenAndServe(":8888", nil)
}
