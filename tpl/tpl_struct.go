package tpl

//GoModTemplate 模板结构体
type GoModTemplate struct {
	Name string
}

//ReturnInfoStruct 赋值结构体
func (g *GoModTemplate) ReturnInfoStruct(name string) {
	g.Name = name
}

//GoSumTemplate 模板结构体
type GoSumTemplate struct {
	Name string
	Age  int
}

//ReturnInfoStruct 赋值结构体
func (g *GoSumTemplate) ReturnInfoStruct(name string) {
	g.Name = name
	g.Age = 16
}
