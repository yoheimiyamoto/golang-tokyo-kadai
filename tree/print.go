package tree

import (
	"fmt"
	"os"
)

func print(t *tree) {
	var s string

	if t.Parent == nil {
		s = "."
		fmt.Fprintln(os.Stdout, s)
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
	fmt.Fprintln(os.Stdout, s)
}
