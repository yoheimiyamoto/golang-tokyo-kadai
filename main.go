package main

import (
	"flag"

	"github.com/YoheiMiyamoto/golang-tokyo-kadai/src/tree"
)

func main() {
	dir := flag.String("d", "./", "dir path")
	maxDepth := flag.Int("l", 100, "max depth")
	ext := flag.String("e", "", "ext")
	flag.Parse()
	tree.Do(*dir, *maxDepth, *ext)
}
