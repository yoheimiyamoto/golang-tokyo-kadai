package simpletree

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

var depth int

func dirWalk(dir string) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	for i, f := range files {
		if f.IsDir() {
			fmt.Printf("%*s%s\n", depth*4, "", f.Name())
			path := filepath.Join(dir, f.Name())
			depth++
			dirWalk(path)
			depth--
		} else {
			if i == 0 {
				fmt.Println(print(f.Name(), depth, TypeFirst))
				// fmt.Printf("%*s%s!\n", depth*4, "", f.Name())
			} else if i == len(files)-1 {
				fmt.Println(print(f.Name(), depth, TypeEnd))
			} else {
				fmt.Println(print(f.Name(), depth, TypeMiddle))
			}
			// fmt.Printf("%*s%s\n", depth*4, "", f.Name())
		}
	}
	return nil
}

type PrintType uint32

const (
	TypeFirst PrintType = iota
	TypeMiddle
	TypeEnd
)

func print(name string, depth int, p PrintType) string {
	var out string
	for i := 0; i < depth-1; i++ {
		out += "│   "
	}
	switch p {
	case TypeFirst:
		out += "└──"
	case TypeMiddle:
		out += "├──"
	case TypeEnd:
		out += "│  "
	}
	out += fmt.Sprintf(" %s", name)
	return out
}
