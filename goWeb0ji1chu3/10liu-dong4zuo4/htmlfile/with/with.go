package with

type teacher struct{
	Name string
	Age int
	Students student
	Classes []class
}
type student struct {
	Name string
	score float64 //todo:这里不大写，前端模板取这个字段无法显示值
}
type class struct{
	Name string
	Num int
}
func Tec() teacher {
	alice := teacher{
		Name: "alice",
		Age:  22,
		Students: student{Name: "jack", score: 66.6},
		Classes: []class{
			class{Name:"class1",Num:45},
			class{Name:"class2",Num:30},
		},
	}
	return alice
}