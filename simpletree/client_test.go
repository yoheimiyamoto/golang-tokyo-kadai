package simpletree

import (
	"testing"
)

func TestDirWalk(t *testing.T) {
	dirWalk("./sample")
}

func TestDepthAdd(t *testing.T) {
	depth := Depth{}
	depth.add(true)
	t.Log(depth)
}

func Test(t *testing.T) {
	a := make([]string, 0)
	a = add(a, "hello")
	t.Log(a)
}

func add(a []string, v string) []string {
	a = append(a, v)
	return a
	// fmt.Println(a)
}

// func TestPrint(t *testing.T) {
// 	s := print("hello.txt", 3, TypeFirst)
// 	t.Log(s)
// }
