//对应range.html文件
package myrange

type student struct {
	Name string
	Age  int
}

func School() *([]student) {
	a := student{"jack", 27}
	b := student{"tom", 28}
	c := student{"mary", 24}
	d := &[]student{a, b, c} //可见切片里也可以填符合类型的变量
	return d
}
