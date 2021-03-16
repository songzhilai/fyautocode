package tpl

var tplToStructMap = map[string]interface{}{
	"go.mod": GoModTemplate{},
	"go.sum": GoSumTemplate{},
}

func tplStruct(name, project string) interface{} {
	tplinter, ok := tplToStructMap[name]
	if !ok {
		return nil
	}
	return TplTypeAssignment(tplinter, project)
}
