package tree

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type tree struct {
	Parent, FirstChild, LastChild, PrevSibling, NextSibling tree
	Name                                                    string
	// Type                                                    FileType
	// DataAtom                                                atom.Atom
}

type FileType uint32

const (
	TypeDir FileType = iota
	TypeFile
)

// func Print(f *File) {
// 	fmt.Println(f.Name)
// 	for f := f.FirstChild; f != nil; f = f.NextSibling {
// 		Print(f)
// 	}
// }

// func getDir(f *File, path []string) *File {

// }
var depth int

func dirWalk(dir string) {
	fmt.Println(dir)
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	for _, f := range files {
		fmt.Println(depth)
		if f.IsDir() {
			depth++
			fmt.Println(depth)
			dirWalk(filepath.Join(dir, f.Name()))
			depth--
			fmt.Println(depth)
		}
		fmt.Println(f.Name())
	}
}

func add(t *tree, v name) {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if f.IsDir() {
		t.FirstChild = 
	}
	t.NextSibling = f
}

if value < t.value {
	t.left = add(t.left, value)
} else {
	t.right = add(t.right, value)
}

// func add(f *File) *File {
// fn := func(f *File) {
// 	for f := f.FirstChild; f != nil; f = f.NextSibling {
// 		Print(f)
// 	}
// }
// for _, p := range path {

// }
// }

// func add(t *tree, value int) *tree {
// 	if t == nil {
// 		// Equivalent to return &tree{value: value}.
// 		t = new(tree)
// 		t.value = value
// 		return t
// 	}
// 	if value < t.value {
// 		t.left = add(t.left, value)
// 	} else {
// 		t.right = add(t.right, value)
// 	}
// 	return t
// }

func FileWalk(path string) error {
	var depth int
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			fmt.Printf("dirName: %s", filepath.Dir(path))
			depth++
		}
		fmt.Println(depth)
		fmt.Println(path)
		fmt.Println(info.Name())
		return nil
	})
	return err
}
