package simpletree

import (
	"encoding/json"
	"fmt"
)

type TreeType uint32

const (
	TreeTypeDir TreeType = iota
	TreeTypeFile
)

type tree struct {
	// Parent, FirstChild, PrevSibling, NextSibling *tree
	Parent, FirstChild, NextSibling *tree
	Name                            string
	TreeType                        TreeType
	// Dirs   []*Dir
}

func (t *tree) json() string {
	data, _ := json.Marshal(t)
	return string(data)
}

// func (d dir) consume() dir {
// 	return d[:len(d)-1]
// }

// func (d dir) first() string {
// 	return d[0]
// }

// func add(t *tree, d dir, name string) {
// 	pre = func(t *tree) {
// 		if d.first() == t.Name
// 	}
// }

// 深さ優先
func eachTree(t *tree, pre, post func(t *tree)) {
	if pre != nil {
		pre(t)
	}
	for c := t.FirstChild; c != nil; c = c.NextSibling {
		eachTree(c, pre, post)
	}
	if post != nil {
		post(t)
	}
}

type dirs []string

func (d dirs) consume() dirs {
	if len(d) > 0 {
		return d[1:]
	}
	return d
}

func (d dirs) isComplete() bool {
	if len(d) == 0 {
		return true
	}
	return false
}

func (d dirs) current() string {
	if len(d) == 0 {
		return ""
	}
	return d[0]
}

/*
ディレクトをみつける、もしくは、ディレクトリがなければ新しいディレクトリを作る
上記の後にdirsをconsumeしていく。

*/
func add(t *tree, d dirs, v string) *tree {
	for c := t.FirstChild; c != nil; c = c.NextSibling {
		fmt.Println("hello")
		d.consume()
		fmt.Println(d.current())
	}
	// if t == nil {
	// 	t = new(tree)
	// 	t.Name = value
	// 	t.TreeType = TreeTypeDir
	// 	t.Namec
	// 	return t
	// }

	// add(t, d, v)
	// for c := t.FirstChild; c != nil; c = c.NextSibling {
	// 	if c.Name == d.current() {
	// 		d = d.consume()
	// 		if d.isComplete() {
	// 			c.Name = value
	// 			return t
	// 		}
	// 		// fmt.Println(c.FirstChild)
	// 		add(c.FirstChild, d, value)
	// 	}
	// }
	return t
}

// func add(t *Tree, value int) *Tree {
// 	if t == nil {
// 		// Equivalent to return &tree{value: value}.
// 		t = new(Tree)
// 		t.Value = value
// 		return t
// 	}
// 	if value < t.Value {
// 		t.Left = add(t.Left, value)
// 	} else {
// 		t.Right = add(t.Right, value)
// 	}
// 	return t
// }

// // 幅優先
// func eachTree2(t *tree) {
// 	for c :=
// }

// func add(t *tree, p path, name string) *tree {
// 	return t
// }
