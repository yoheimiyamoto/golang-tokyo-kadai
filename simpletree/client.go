package simpletree

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

/*
true = last
*/
type Depth []bool

func (d Depth) add(b bool) Depth {
	return append(d, b)
}

func (d Depth) remove() Depth {
	if len(d) == 0 {
		return d
	}
	return d[0 : len(d)-1]
}

func (d Depth) int() int {
	return len(d)
}

var depth = make(Depth, 0)

func dirWalk(dir string) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	for i, f := range files {
		// fmt.Println(depth)
		if f.IsDir() {
			// fmt.Printf("%*s%s\n", depth.int()*4, "", f.Name())
			fmt.Printf("%s, %s\n", dir, print(f.Name(), depth))
			if i == len(files)-1 {
				depth = depth.add(true)
			} else {
				depth = depth.add(false)
			}
			dirWalk(filepath.Join(dir, f.Name()))
			depth = depth.remove()
		} else {
			fmt.Printf("%s, %s\n", dir, print(f.Name(), depth))
			// if i == 0 {
			// 	fmt.Println(print(f.Name(), depth))
			// 	// fmt.Printf("%*s%s!\n", depth*4, "", f.Name())
			// } else if i == len(files)-1 {
			// 	fmt.Println(print(f.Name(), depth))
			// } else {
			// 	fmt.Println(print(f.Name(), depth))
			// }
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

func print(name string, depth Depth) string {
	var out string
	for i, d := range depth {
		if i == len(depth)-1 {
			if d {
				out += "└──"
			} else {
				out += "├──"
			}
			continue
		}
		if !d {
			out += "│   "
		}
	}
	out += fmt.Sprintf(" %s, depth: %v", name, depth)
	return out
}
