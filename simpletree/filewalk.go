package simpletree

import (
	"fmt"
	"os"
	"path/filepath"
)

func parse(dirPath string) error {
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return err
	})
	return err
}
