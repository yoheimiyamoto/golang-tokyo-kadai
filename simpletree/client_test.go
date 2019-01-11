package simpletree

import "testing"

func TestDirWalk(t *testing.T) {
	dirWalk("./sample")
}

func TestPrint(t *testing.T) {
	s := print("hello.txt", 3, TypeFirst)
	t.Log(s)
}
