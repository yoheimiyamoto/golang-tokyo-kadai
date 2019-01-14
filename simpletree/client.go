package simpletree

import (
	"io/ioutil"
	"path/filepath"
)

/*
true = last
*/
// type Depth []bool

// func (d Depth) add(b bool) Depth {
// 	return append(d, b)
// }

// func (d Depth) remove() Depth {
// 	if len(d) == 0 {
// 		return d
// 	}
// 	return d[0 : len(d)-1]
// }

// func (d Depth) int() int {
// 	return len(d)
// }

// var depth = make(Depth, 0)

func dirWalk(dir string, pre, post func(path string)) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, f := range files {
		path := filepath.Join(dir, f.Name())
		if f.IsDir() {
			if pre != nil {
				pre(path)
			}
			dirWalk(path, pre, post)
			if post != nil {
				post(path)
			}
		}
	}
	return nil
}

// func parse(dir string) (*tree, error) {

// 	err := dir(dir, pre, nil)
// 	if err != nil {
// 		return nil, err
// 	}
// }

// type PrintType uint32

// const (
// 	TypeFirst PrintType = iota
// 	TypeMiddle
// 	TypeEnd
// )

// func print(name string, depth Depth) string {
// 	var out string
// 	for i, d := range depth {
// 		if i == len(depth)-1 {
// 			if d {
// 				out += "└──"
// 			} else {
// 				out += "├──"
// 			}
// 			continue
// 		}
// 		if !d {
// 			out += "│   "
// 		}
// 	}
// 	out += fmt.Sprintf(" %s, depth: %v", name, depth)
// 	return out
// }
