package main

import (
	"flag"

	"github.com/YoheiMiyamoto/golang-tokyo-kadai/src/tree"
)

func main() {
	dir := flag.String("dir", "./", "dir path")
	maxDepth := flag.Int("L", 100, "max depth")
	flag.Parse()
	tree.Do(*dir, *maxDepth)
}
