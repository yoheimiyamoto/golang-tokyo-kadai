package tree

import "testing"

// func TestWalk(t *testing.T) {
// 	FileWalk("./sample")
// }

// func TestPrint(t *testing.T) {
// 	f := &File{
// 		Name: "one",
// 		FirstChild: &File{
// 			Name: "two",
// 		},
// 	}
// 	Print(f)
// }

func TestDirWalk(t *testing.T) {
	dirWalk("./sample")
}
