package algs

import (
	"fmt"
	"slices"
)

func Match(text, pattern string) int {
	text = text + "$"
	cur := BuildUkkonen(text)
	pos := 1
	for i := range len(pattern) {
		if cur.from+pos > cur.to {
			next, ok := cur.childs[pattern[i]]
			if !ok {
				return -1
			}
			cur = next
			pos = 1
		} else {
			if text[cur.from+pos] == pattern[i] {
				pos++
			} else {
				return -1
			}
		}
	}
	return cur.from + pos - len(pattern)
}

func PrintSTree(node *STreeNode, prefix string, isLast bool, s string) {
	if node == nil {
		return
	}
	connector := "├── "
	nextPrefix := prefix + "│   "
	if isLast {
		connector = "└── "
		nextPrefix = prefix + "    "
	}
	st := "root"
	if node.from >= 0 {
		st = s[node.from : node.to+1]
	}
	fmt.Printf(
		"%s%s%s[%d:%d]\n",
		prefix,
		connector,
		st,
		node.from,
		node.to,
	)
	keys := make([]byte, 0, len(node.childs))
	for k := range node.childs {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	for i, k := range keys {
		child := node.childs[k]
		isLastChild := i == len(keys)-1
		fmt.Printf("%s%s\n", nextPrefix, "│   ")
		PrintSTree(child, nextPrefix, isLastChild, s)
	}
}
