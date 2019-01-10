package tree

import (
	"fmt"
	"os"
	"path/filepath"
)

// type File interface {
// }

type File struct {
	Name string
}

type Dir struct {
	Name  string
	Files []File
}

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
