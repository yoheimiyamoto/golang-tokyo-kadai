package simpletree

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

/*
filewalkしてparseする。
*/

// func parse(dirPath string) (*tree, error) {
// 	// add := func(t *tree) {

// 	// }
// 	c := new(tree) // カレントディレクトリ
// 	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
// 		if info.IsDir() {
// 			addDir(c, info.Name())
// 			return nil
// 		}
// 		c.AddSiblint(&tree{Name: info.Name()})
// 		fmt.Println(path)
// 		return err
// 	})
// 	// return nil, err
// }

// func fileWalk(dirPath string) error {
// 	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {

// 		fmt.Println(path)
// 		return err
// 	})
// 	return err
// }

func parse(dir string) (*tree, error) {
	t := new(tree)
	c := t
	// c := new(tree) // カレントディレクトリ
	// var depth int
	pre := func(name string) {
		// c.FirstChild = &tree{Name: name, TreeType: TreeTypeDir, Parent: c}
		t = &tree{Name: name, TreeType: TreeTypeDir}
		c.FirstChild = t
		t.Parent = c
		// c.FirstChild.Parent = c

		// c.FirstChild = new(tree)
		// c = c.FirstChild

		// depth++
		// fmt.Println(depth)

		// c.FirstChild = &tree{Name: name, TreeType: TreeTypeDir, Parent: c}
	}
	post := func(name string) {
		// c = c.Parent

		// fmt.Println(depth)

		// depth--
		// fmt.Println(depth)
		// c = c.Parent
	}
	main := func(name string) {
		// fmt.Println("hello")
		fmt.Println(name)
	}
	dirWalk(dir, pre, post, main)
	return t, nil
}

/*
dirWalkしてparseする
*/
func dirWalk(dir string, pre, post, main func(name string)) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, f := range files {
		// path := filepath.Join(dir, f.Name())
		if f.IsDir() {
			if pre != nil {
				// pre(path)
				pre(f.Name())
			}
			dirWalk(filepath.Join(dir, f.Name()), pre, post, main)
			if post != nil {
				// post(path)
				post(f.Name())
			}
		}
		if main != nil {
			fmt.Println(filepath.Join(dir, f.Name()))
			main(f.Name())
		}
	}
	return nil
}

/*
ディレクトリを作る。
FirstChildがなければ作る。
FirstChildがあればそのチャイルドのNextSiblingを追加する
*/
func addDir(t *tree, name string) {
	dir := &tree{Name: name, TreeType: TreeTypeDir}
	if t.FirstChild == nil {
		t.FirstChild = dir
		return
	}
	addNextSibling(t.FirstChild, dir)
	return
}

/*
dst = 追加する先
t = 追加する対象
*/
func addNextSibling(dst, t *tree) {
	if dst.NextSibling == nil {
		dst.NextSibling = t
		return
	}
	for c := dst.NextSibling; c != nil; c = c.NextSibling {
		if c.NextSibling == nil {
			c.NextSibling = t
			return
		}
	}
	return
}
