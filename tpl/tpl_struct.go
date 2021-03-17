package tpl

//GoModTemplate 模板结构体
type GoModTemplate struct {
}

//ReturnInfoStruct 赋值结构体
func (g GoModTemplate) ReturnInfoStruct(name string) map[string]interface{} {
	result := make(map[string]interface{})
	result["Name"] = name
	return result
}

//GoSumTemplate 模板结构体
type GoSumTemplate struct {
}

//ReturnInfoStruct 赋值结构体
func (g GoSumTemplate) ReturnInfoStruct(name string) map[string]interface{} {
	result := make(map[string]interface{})
	result["Name"] = "Name"
	result["Name2"] = "Name2"
	result["Name3"] = "Name3"
	result["Name4"] = "Name4"
	result["Name5"] = "Name5"
	result["Name6"] = "Name6"
	result["Name7"] = "Name7"
	result["Name8"] = "Name8"
	result["Name9"] = "Name9"
	result["Name10"] = "Name10"
	return result
}
