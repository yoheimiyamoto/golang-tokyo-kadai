package simpletree

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
	// Parent, FirstChild, PrevSibling, NextSibling *tree
	Parent, FirstChild, NextSibling *tree
	Name                            string
	TreeType                        TreeType
}

func (t *tree) json() string {
	data, _ := json.Marshal(t)
	var buf bytes.Buffer
	json.Indent(&buf, data, "", "  ")
	return buf.String()
}

func (t *tree) AddSiblint(_t *tree) {
	for c := t; c != nil; c = c.NextSibling {
		if c.NextSibling == nil {
			c.NextSibling = _t
			return
		}
	}
}

type dirs []string

func (d dirs) consume() dirs {
	if len(d) > 0 {
		return d[1:]
	}
	return d
}

func (d dirs) isComplete() bool {
	if len(d) == 0 {
		return true
	}
	return false
}

func (d dirs) current() string {
	if len(d) == 0 {
		return ""
	}
	return d[0]
}
