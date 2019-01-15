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
	// dst := new(tree)
	// dst.NextSibling = &tree{Name: "one"}
	// tre := new(tree)
	// tre.Name = "hello"
	// addNextSibling(dst, tre)
	// actual := dst.NextSibling.NextSibling.Name
	// expected := "hello"
	// if actual != expected {
	// 	t.Errorf("actual: %s, expected: %s", actual, expected)
	// }
}

func TestAddDir(t *testing.T) {
	// tre := new(tree)
	// tre.FirstChild = &tree{Name: "one"}
	// addDir(tre, "test")
	// actual := tre.FirstChild.NextSibling.Name
	// expected := "test"
	// if actual != expected {
	// 	t.Errorf("actual: %s, expected: %s", actual, expected)
	// }
}

func TestParse(t *testing.T) {
	tre, _ := parse("./sample")
	print(tre)
	// t.Log(tre.FirstChild.NextSibling.Name)
	// t.Log(fmt.Sprintf("%v", tre))
}

func TestTest(t *testing.T) {
	// p := new(tree)
	// c := new(tree)
	// p.FirstChild = c
	// c.Parent = p
	// t.Log(p.json())
	// t.Log(c.json())

	// n.LastChild = c
	// c.Parent = n
}

// func TestTest2(t *testing.T) {
// 	p := new(html.Node)
// 	c := new(html.Node)
// 	p.FirstChild = c
// 	c.Parent = p

// }

func TestHTMLParse(t *testing.T) {
	// r := strings.NewReader("<html><p>hello</p><h2>h2!</h2></html>")
	// doc, err := html.Parse(r)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// t.Log(doc.FirstChild.FirstChild.NextSibling.FirstChild.NextSibling.Data)
}
