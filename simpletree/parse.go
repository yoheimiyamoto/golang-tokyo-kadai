package simpletree

import (
	"io/ioutil"
	"path/filepath"
)

/*
filewalkしてparseする。
*/

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
