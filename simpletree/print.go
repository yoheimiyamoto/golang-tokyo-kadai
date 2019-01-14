package simpletree

import "fmt"

func print(t *tree) {
	fn := func(t *tree) {
		fmt.Println(t.Name)
	}
	eachTree(t, fn, nil)
}

func eachTree(t *tree, pre, post func(t *tree)) {
	if pre != nil {
		pre(t)
	}
	for c := t.FirstChild; c != nil; c = c.NextSibling {
		eachTree(c, pre, post)
	}
	if post != nil {
		post(t)
	}
}
