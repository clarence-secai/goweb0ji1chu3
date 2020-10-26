package DB

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"go2wang3shang4shu1cheng2/wang3ye4shu1cheng2/200226liu/material"
)

func Get(name string) material.User {
	//打开数据库
	sqlstr := "select id,username,password,email from users where username=?"
	myrow := material.MyDB.QueryRow(sqlstr, name)
	var myuser material.User
	err := myrow.Scan(&myuser.ID, &myuser.Name, &myuser.Password, &myuser.Email)
	if err == sql.ErrNoRows {
		return material.User{}
	} else {
		return myuser
	}
}
func Save(user material.User) bool {
	sqlstr := "insert into users(username,password,email) values (?,?,?) "
	material.MyDB.Exec(sqlstr, user.Name, user.Password, user.Email)
	// if err != nil {
	// 	return true
	// }else {
	return true
	// }
}
