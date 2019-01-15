package tree

import (
	"fmt"
	"os"
)

// treeをコマンドライン出力する関数です
func print(t *tree) {
	var s string

	// rootディレクトリの場合
	if t.Parent == nil {
		s = "."
		fmt.Fprintln(os.Stdout, s)
		return
	}

	// 直上ディレクトリからファイルまでのノード
	s = t.Name
	if t.NextSibling == nil {
		s = fmt.Sprintf("└── %s", s)
	} else {
		s = fmt.Sprintf("├── %s", s)
	}

	// rootから直上ディレクトリまでのノード
	for c := t.Parent; c.Parent != nil; c = c.Parent {
		if c.NextSibling == nil {
			s = fmt.Sprintf("     %s", s)
		} else {
			s = fmt.Sprintf("│    %s", s)
		}
	}
	fmt.Fprintln(os.Stdout, s)
}
