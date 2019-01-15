package tree

type TreeType uint32

const (
	TreeTypeDir TreeType = iota
	TreeTypeFile
)

type tree struct {
	Parent, FirstChild, LastChild, NextSibling *tree
	Name                                       string
	TreeType                                   TreeType
}
