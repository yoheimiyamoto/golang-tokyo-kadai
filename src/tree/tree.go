package tree

func Do(dir string, maxDepth int) error {
	tre, err := parse(dir)
	if err != nil {
		return err
	}
	var depth int

	var eachTree func(t *tree, pre, post func(t *tree))
	eachTree = func(t *tree, pre, post func(t *tree)) {
		if pre != nil {
			pre(t)
		}
		depth++
		for c := t.FirstChild; c != nil; c = c.NextSibling {
			if depth <= maxDepth {
				eachTree(c, pre, post)
			}
		}
		depth--
		if post != nil {
			post(t)
		}
	}
	eachTree(tre, print, nil)
	return nil
}

// func eachTree(t *tree, pre, post func(t *tree), maxDepth int) {
// 	var depth int
// 	fmt.Printf("depth: %d", depth)
// 	if pre != nil {
// 		pre(t)
// 	}
// 	depth++
// 	for c := t.FirstChild; c != nil; c = c.NextSibling {
// 		if depth < maxDepth {
// 			eachTree(c, pre, post, maxDepth)
// 		}
// 	}
// 	depth--
// 	if post != nil {
// 		post(t)
// 	}
// }
