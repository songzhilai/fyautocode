package commands

import (
	"fmt"
	"fyautocode/tpl"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func modPath(p string) string {
	dir := filepath.Dir(p)
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			content, _ := ioutil.ReadFile(filepath.Join(dir, "go.mod"))
			mod := RegexpReplace(`module\s+(?P<name>[\S]+)`, string(content), "$name")
			name := strings.TrimPrefix(filepath.Dir(p), dir)
			name = strings.TrimPrefix(name, string(os.PathSeparator))
			if name == "" {
				return fmt.Sprintf("%s/", mod)
			}
			return fmt.Sprintf("%s/%s/", mod, name)
		}
		parent := filepath.Dir(dir)
		if dir == parent {
			return ""
		}
		dir = parent
	}
}

// RegexpReplace replace regexp
func RegexpReplace(reg, src, temp string) string {
	var result []byte
	pattern := regexp.MustCompile(reg)
	for _, subMatches := range pattern.FindAllStringSubmatchIndex(src, -1) {
		result = pattern.ExpandString(result, temp, src, subMatches)
	}
	return string(result)
}

//go:generate packr2
func create(path, project string) (err error) {
	err = os.RemoveAll(f.Path)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if err = os.MkdirAll(f.Path, 0755); err != nil {
		return err
	}
	return tpl.DirTemplate(path, project, f.Type)
}

// func generate(path string) error {
// 	cmd := exec.Command("go", "generate", path)
// 	cmd.Dir = f.Path
// 	cmd.Stdout = os.Stdout
// 	cmd.Stderr = os.Stderr
// 	return cmd.Run()
// }

// func write(path, tpl string) (err error) {
// 	data, err := parse(tpl)
// 	if err != nil {
// 		return
// 	}
// 	return ioutil.WriteFile(path, data, 0644)
// }

// func parse(s string) ([]byte, error) {
// 	t, err := template.New("").Parse(s)
// 	if err != nil {
// 		return nil, err
// 	}
// 	var buf bytes.Buffer
// 	if err = t.Execute(&buf, f); err != nil {
// 		return nil, err
// 	}
// 	return buf.Bytes(), nil
// }

// func buildDir(base string, cmd string, n int) string {
// 	dirs, err := ioutil.ReadDir(base)
// 	if err != nil {
// 		panic(err)
// 	}
// 	for _, d := range dirs {
// 		if d.IsDir() && d.Name() == cmd {
// 			return path.Join(base, cmd)
// 		}
// 	}
// 	if n <= 1 {
// 		return base
// 	}
// 	return buildDir(filepath.Dir(base), cmd, n-1)
// }
