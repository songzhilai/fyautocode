package tpl

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

var Template = "template" //模板文件路劲
var DEPTH = 100           //模板文件深度
var TPLSuffix = ".tmpl"   //模板文件后缀

//DirTemplate 解析模板文件
func DirTemplate(paths, project, typestr string) error {
	absolutePath := strings.TrimSuffix(paths, project)
	fmt.Println("absolutePath", absolutePath)
	// files := getPathFileList(paths)
	filespathmap := make(map[string][]string)
	err := walkDir(fmt.Sprintf("%s/%s/%s", paths, Template, typestr), 10, filespathmap)
	if err != nil {
		return err
	}
	fmt.Println("walkDir", filespathmap)
	for k, v := range filespathmap {
		for _, fv := range v {
			tmpl, err := template.ParseFiles(fmt.Sprintf("%s/%s", k, fv))
			if err != nil {
				return err
			}
			//去掉模板文件后缀
			newfilename := strings.TrimSuffix(fv, TPLSuffix)
			structinfo := tplStruct(newfilename, project)
			// if structinfo == nil {
			// 	return errors.New("struct is nil")
			// }
			fmt.Println("path", k, fmt.Sprintf("%s/%s/%s", absolutePath, Template, typestr))
			//替换成相对路劲
			relativePath := fmt.Sprintf("./%s/%s", project, strings.TrimPrefix(k, fmt.Sprintf("%s/%s/%s", absolutePath, Template, typestr)))
			if relativePath != "" {
				//创建文件夹
				FYMkdir(relativePath)
			}
			ff, err := os.Create(fmt.Sprintf("./%s/%s", relativePath, newfilename))
			if err != nil {
				return err
			}
			err = tmpl.Execute(ff, structinfo)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func walkDir(dirpath string, depth int, filespathmap map[string][]string) error {
	if depth > DEPTH { //大于设定的深度
		return errors.New("设置的文件最大深度为10，超过该深度")
	}
	files, err := ioutil.ReadDir(dirpath) //读取目录下文件
	if err != nil {
		return err
	}
	for _, file := range files {
		if file.IsDir() {
			//创建生成的文件夹
			walkDir(dirpath+"/"+file.Name(), depth+1, filespathmap)
			continue
		} else {
			filespathmap[dirpath] = append(filespathmap[dirpath], file.Name())
		}
	}
	return nil
}

//PathExists 判断文件、文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// FYMkdir 创建文件夹
func FYMkdir(dir string) {
	exist, err := PathExists(dir)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		if exist {
			fmt.Println(dir + "文件夹已存在！")
		} else {
			// 文件夹名称，权限
			err := os.MkdirAll(dir, os.ModePerm)
			if err != nil {
				fmt.Println(dir+"文件夹创建失败：", err.Error())
			} else {
				fmt.Println(dir + "文件夹创建成功！")
			}
		}
	}
}
