//对应range2.HTML文件
package myrange

type user struct {
	Name     string
	Password int
}

func Usermap() map[int]user {
	a := make(map[int]user) //这里需要先make
	a[1] = user{"jack", 111}
	a[2] = user{"tom", 222}
	a[3] = user{"mary", 333}
	return nil
}
