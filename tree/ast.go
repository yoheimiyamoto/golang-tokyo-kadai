package tree

import (
	"bytes"
	"encoding/json"
)

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

func (t *tree) json() string {
	data, _ := json.Marshal(t)
	var buf bytes.Buffer
	json.Indent(&buf, data, "", "  ")
	return buf.String()
}
