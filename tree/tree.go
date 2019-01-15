package tree

func Do(dir string) error {
	tre, err := parse(dir)
	if err != nil {
		return err
	}
	eachTree(tre, print, nil)
	return nil
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
