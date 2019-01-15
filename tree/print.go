package tree

import "fmt"

func print(t *tree) {
	pre := func(t *tree) {
		var s string

		if t.Parent == nil {
			s = "."
			fmt.Println(s)
			return
		}

		s = t.Name
		if t.NextSibling == nil {
			s = fmt.Sprintf("└── %s", s)
		} else {
			s = fmt.Sprintf("├── %s", s)
		}

		for c := t.Parent; c.Parent != nil; c = c.Parent {
			if c.NextSibling == nil {
				s = fmt.Sprintf("     %s", s)
			} else {
				s = fmt.Sprintf("│    %s", s)
			}
		}
		fmt.Println(s)
	}
	eachTree(t, pre, nil)
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
