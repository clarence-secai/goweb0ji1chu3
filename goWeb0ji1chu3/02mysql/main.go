package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type user struct {
	id       int
	username string
	password string
	email    string
}

func main() {
	mydb, err := sql.Open("mysql", "root:413188ok@tcp(localhost:3306)/test")
	if err != nil {
		fmt.Println(err)
		return
	}
	//向数据库写入一行信息
	u := user{username: "bob", password: "123", email: "123@sohu.com"}
	_, err = mydb.Exec("insert into users(username,password,email)values(?,?,?)", u.username, u.password, u.email)
	if err != nil {
		fmt.Println("err2=", err)
		return
	}

	//查询id为2的信息
	myRow := mydb.QueryRow("select id,username,password,email from users where id=?", 2)
	var myres user
	myRow.Scan(&myres.id, &myres.username, &myres.password, &myres.email)
	fmt.Println("myres=", myres)

	//用有prepare查询id为4的信息
	myStmt, err2 := mydb.Prepare("select id,username,password,email from users where id =?")
	if err2 != nil {
		fmt.Println("err2=", err2)
		return
	}
	meRow := myStmt.QueryRow(4)
	var a user
	err3 := meRow.Scan(&a.id, &a.username, &a.password, &a.email)
	if err3 != nil {
		fmt.Println("err3=", err3)
		return
	}
	fmt.Println("a=", a)

	//查询数据库中test表格的所有信息
	allRow, err4 := mydb.Query("select id,username,password,email from users")
	if err4 != nil {
		fmt.Println("err4=", err4)
		return
	}
	defer allRow.Close() //这一步最好别缺，当然，根据官方文档，当Next()返回值为false时Close()会自动执行
	var alluser user
	for { //这里可以直接 for allRow.Next {} 从而省掉if语句，更简洁，详细见官方文档示例
		if allRow.Next() {
			allRow.Scan(&alluser.id, &alluser.username, &alluser.password, &alluser.email)
			fmt.Println(alluser)
		} else {
			break
		}
	}

}
