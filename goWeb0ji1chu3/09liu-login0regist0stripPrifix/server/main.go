package main

import (
	"go2web/wang3ye4shu1cheng2/200226liu/DB"
	"go2wang3shang4shu1cheng2/wang3ye4shu1cheng2/200226liu/material"
	"html/template"
	"net/http"
)

//以下各个处理器，parsefiles函数内的文件路径，都应当是以当前文件为基准找到目标文件的路径，另外，开头不能再
//加斜杠。以斜杠为开头表示省略了斜杠前面的东西，例如主函数中的handlefunc里，第一个参数以斜杠开头，表示斜杠前面
//是localhost：80,只是省略了。如果是不借助服务器，直接用浏览器打开各个HTML文件，通过点击HTML中的链接来进入另一个
//HTML文件页面，链接的路径写法同理，是以链接所在文件所处的位置为基准找到链接指向文件的路径，不能多加一个开头斜杠

func index(rw http.ResponseWriter, r *http.Request) {
	mytemplate, err := template.ParseFiles("../pages/index.html") //这里的路径需是以当前文件所属位置为基准
	if err != nil {
		return
	}
	mytemplate.Execute(rw, "")
}
func index2(rw http.ResponseWriter, r *http.Request) {
	mytemplate, err := template.ParseFiles("../pages/index2.html") //这里的路径需是以当前文件所属位置为基准
	if err != nil {
		return
	}
	mytemplate.Execute(rw, "")
}

func loginPage(rw http.ResponseWriter, r *http.Request) {
	mytemplate, _ := template.ParseFiles("../pages/login.html")
	mytemplate.Execute(rw, "")
}
func registPage(rw http.ResponseWriter, r *http.Request) {
	mytemplate, _ := template.ParseFiles("../pages/regist.html")
	mytemplate.Execute(rw, "")
}

//todo:图片是不能类比HTML文件模板的方式来用的
// func tu(rw http.ResponseWriter,r *http.Request){
// 	mytemplate,_ := template.ParseFiles("../piture/tu.jpg")
// 	mytemplate.Execute(rw,"")
// }

func login(rw http.ResponseWriter, r *http.Request) {
	myname := r.PostFormValue("yonghuming")
	mymima := r.PostFormValue("mima")
	if DB.Get(myname).Password == mymima { //todo:下面一行再次使用了login.html
		mytemplate := template.Must(template.ParseFiles("../pages/login.html"))
		mytemplate.Execute(rw, "登录成功,3秒后返回主页")
		//time.Sleep(time.Second*3)
		mytemplate2 := template.Must(template.ParseFiles("../pages/index2.html"))
		mytemplate2.Execute(rw, "")
	} else {
		mytemplate := template.Must(template.ParseFiles("../pages/login.html"))
		mytemplate.Execute(rw, "用户名或密码不正确")
	}
}

func regist(rw http.ResponseWriter, r *http.Request) {
	myyonghuming := r.PostFormValue("yonghuming")
	mymima := r.PostFormValue("mima")
	myyouxiang := r.PostFormValue("youxiang")
	var myuser = material.User{Name: myyonghuming, Password: mymima, Email: myyouxiang}
	niluser := material.User{ID: 0}
	if DB.Get(myyonghuming) == niluser {
		if DB.Save(myuser) {
			mytemplate := template.Must(template.ParseFiles("../pages/regist.html"))
			mytemplate.Execute(rw, "注册成功，3秒后返回主页")
			//time.Sleep(time.Second*3)
			//下面这两种办法都不行，没法做到网页的自动跳转
			// mytemplate2:=template.Must(template.ParseFiles("../pages/index2.html"))
			// mytemplate2.Execute(rw,"")
			// rw.Header().Set("Location","https:www.baidu.com")
			// rw.WriteHeader(302)

			//todo:下面这个返回的内容也是和上面返回的内容在同一个页面
			// 也就是说，一个HTML文件只代表一个内容，而不是一个页面
			mytemplate3 := template.Must(template.ParseFiles("../pages/success.html"))
			mytemplate3.Execute(rw, "")
			//todo:总结来说，就是一个处理器内，再多的execute都
			// 会都显示在同一页面上，就像将多个HTML按这里的代码顺序拼成一个HTML
		}
	} else {
		mytemplate := template.Must(template.ParseFiles("../pages/regist.html"))
		mytemplate.Execute(rw, "用户已存在，请重新注册")

	}
}

func main() {
	http.HandleFunc("/main", index)
	http.HandleFunc("/index2", index2)
	//http.HandleFunc("/picture/tu.jpg",tu)

	//todo:注意！下面这个strip是指将第一个参数引号内的,除了两斜杠(即两斜杠保留)的两斜杠之间的内容替换为http.Dir
	// 中括号内的不含引号的两引号内的内容。故http.Dir中的内容首位不能再加斜杠，路径为以当前文件为基准的到达目标文件的路径

	http.Handle("/picture/", http.StripPrefix("/picture/", http.FileServer(http.Dir("../picture"))))

	//todo:注册处理器中的第一个参数也是可以像下面这样有多个 斜杠的
	http.HandleFunc("/pages/login.html", loginPage)
	http.HandleFunc("/pages/regist.html", registPage)

	http.HandleFunc("/login", login)
	http.HandleFunc("/regist", regist)
	http.ListenAndServe(":9999", nil)
}
