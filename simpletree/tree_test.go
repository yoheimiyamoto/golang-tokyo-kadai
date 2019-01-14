package simpletree

import (
	"testing"
)

func TestAddSiblint(t *testing.T) {
	tre := new(tree)
	_t := new(tree)
	_t.Name = "hello"
	tre.AddSiblint(_t)
	t.Log(tre.json())
}

func TestEachTree(t *testing.T) {
	// tre := &tree{
	// 	// Name: "one",
	// 	FirstChild: &tree{
	// 		Name: "one",
	// 		NextSibling: &tree{
	// 			Name: "two",
	// 		},
	// 	},
	// }
	// fn := func(t *tree) {
	// 	fmt.Println(t.Name)
	// }
	// eachTree(tre, fn, nil)
}

func TestDir(t *testing.T) {
	// d := dir{"one", "two"}

	// actual := d.current()
	// expected := "one"
	// if actual != expected {
	// 	t.Errorf("actual: %s, expected: %s", actual, expected)
	// }

	// d = d.consume()
	// actual = d.current()
	// expected = "two"
	// if actual != expected {
	// 	t.Errorf("actual: %s, expected: %s", actual, expected)
	// }
}

func TestAdd(t *testing.T) {
	// tre := &tree{
	// 	FirstChild: &tree{
	// 		Name: "one",
	// 		FirstChild: &tree{
	// 			Name: "two",
	// 		},
	// 		// NextSibling: &tree{
	// 		// 	Name: "two",
	// 		// },
	// 	},
	// }
	// tre := new(tree)
	// tre = add(tre, dir{"one", "two"}, "four")
	// add(tre, dirs{"one", "two"}, "four")
	// t.Log(tre.json())
}
