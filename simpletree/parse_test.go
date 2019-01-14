package simpletree

import (
	"fmt"
	"testing"
)

func TestDirWalk(t *testing.T) {
	fn := func(path string) {
		fmt.Println(path)
	}
	dirWalk("./sample", fn, nil)
}
