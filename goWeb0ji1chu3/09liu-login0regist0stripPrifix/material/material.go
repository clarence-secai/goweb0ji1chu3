package material

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID       int
	Name     string
	Password string
	Email    string
}

var MyDB *sql.DB //todo:这里不var，直接在函数中:=会报myDB申明而未使用
func init() {
	MyDB, _ = sql.Open("mysql", "root:413188ok@tcp(localhost:3306)/test")
	fmt.Println(MyDB.Ping())
}

//todo:这个init函数没法放在db.go一起，因为没法在db.go里的函数
// 中去直接用这个打开获得的MyDB，原因是作用域的问题，放在一起则又在
// 同一个包中，所以不好弄，不然开一个新的文件夹专门放其他包都要用到
// 的东西
