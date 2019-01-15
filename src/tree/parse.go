package tree

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

/*
対象のディレクトリ配下をtreeにパースする関数です。
dir = 対象のディレクトリパス
ext = 対象のファイル拡張子
*/
func parse(dir, ext string) (*tree, error) {
	root := &tree{Name: "root", TreeType: TreeTypeDir}
	var c *tree // カレントディレクトリ
	c = root

	// 新しいフォルダをカレントディレクトの下の階層に作成
	// カレントディクレトリを新しく作成したディレクトに移動
	pre := func(name string) {
		children := &tree{Name: name, TreeType: TreeTypeDir}
		add(c, children)
		c = children
	}

	// カレントディレクトリを1つ上の階層に移動
	post := func(name string) {
		c = c.Parent
	}

	// カレントディレクトリにファイル作成
	main := func(name string) {
		if ext != "" {
			if _ext := filepath.Ext(name); _ext != "" {
				if _ext != fmt.Sprintf(".%s", ext) {
					return
				}
			}
		}
		f := &tree{Name: name, TreeType: TreeTypeFile}
		add(c, f)
	}

	fileWalk(dir, pre, post, main)
	return root, nil
}

func fileWalk(dir string, pre, post, main func(name string)) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, f := range files {
		if f.IsDir() {
			if pre != nil {
				pre(f.Name())
			}
			fileWalk(filepath.Join(dir, f.Name()), pre, post, main)
			if post != nil {
				post(f.Name())
			}
		} else {
			if main != nil {
				main(f.Name())
			}
		}
	}
	return nil
}

/*
treeの配下にtreeを追加する関数です。
dst = 追加する先のtree
t = 追加するtree
*/
func add(dst, t *tree) {
	if dst == nil {
		return
	}
	t.Parent = dst
	if dst.FirstChild == nil {
		dst.FirstChild = t
		dst.LastChild = t
		return
	}
	dst.LastChild.NextSibling = t
	dst.LastChild = t
}
