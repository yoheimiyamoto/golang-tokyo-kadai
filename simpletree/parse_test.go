package simpletree

import (
	"testing"
)

// func TestDirWalk(t *testing.T) {
// 	fn := func(path string) {
// 		fmt.Println(path)
// 	}
// 	dirWalk("./sample", fn, nil)
// }

func TestAddNextSibling(t *testing.T) {
	dst := new(tree)
	dst.NextSibling = &tree{Name: "one"}
	tre := new(tree)
	tre.Name = "hello"
	addNextSibling(dst, tre)
	actual := dst.NextSibling.NextSibling.Name
	expected := "hello"
	if actual != expected {
		t.Errorf("actual: %s, expected: %s", actual, expected)
	}
}

func TestAddDir(t *testing.T) {
	tre := new(tree)
	tre.FirstChild = &tree{Name: "one"}
	addDir(tre, "test")
	actual := tre.FirstChild.NextSibling.Name
	expected := "test"
	if actual != expected {
		t.Errorf("actual: %s, expected: %s", actual, expected)
	}
}

func TestParse(t *testing.T) {
	tre, _ := parse("./sample")
	t.Log(tre.json())
}
