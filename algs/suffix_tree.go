// Package algs
package algs

type STreeNode struct {
	childs map[byte]*STreeNode
	from   int
	to     int
	sl     *STreeNode
	parent *STreeNode
}

type ActivePoint struct {
	node        *STreeNode
	l           int
}

func (ap *ActivePoint) isEndOfNode() bool {
	if isRoot(ap.node) {
		return true
	}
	return ap.node.to == ap.node.from+ap.l
}

func (ap *ActivePoint) hasNextSymb(c byte, s string) bool {
	if ap.isEndOfNode() {
		_, existance := ap.node.childs[c]
		return existance
	}
	return s[ap.node.from+ap.l+1] == c
}

func (ap *ActivePoint) goNext(c byte) {
	if ap.isEndOfNode() {
		ap.node = ap.node.childs[c]
		ap.l = 0
	} else {
		ap.l++
	}
}

func (ap *ActivePoint) sl(s string) {
	tNode := ap.node.parent.sl.childs[s[ap.node.from]]
	if isRoot(ap.node.parent) {
		// если корень, то суф ссылки не работают и мы должны остаться в корне для смещения на один, так как суффикс при перехода меньше на один
		tNode = ap.node.parent
	}
	// skip-count
	tl := ap.l
	offset := 0
	for tNode.from+tl > tNode.to {
		offset += tNode.to - tNode.from + 1
		tl -= tNode.to - tNode.from + 1
		tNode = tNode.childs[s[ap.node.from+offset]]
	}
	ap.node = tNode
	ap.l = tl
}

func BuildUkkonen(s string) *STreeNode {
	root := &STreeNode{
		childs: make(map[byte]*STreeNode),
		from:   -1,
		to:     -1,
	}
	root.sl = root
	root.parent = root
	ap := ActivePoint{
		node: root,
		l:    0,
	}
	i := 0
	prevNode := &STreeNode{}
	for i < len(s) {
		if ap.hasNextSymb(s[i], s) {
			ap.goNext(s[i])
			i++
			prevNode = &STreeNode{}
			continue
		}
		extendNode(&ap, i, s, root)
		prevNode.sl = ap.node
		prevNode = ap.node
		if isRoot(ap.node) {
			prevNode.sl = root
			i++
			continue
		}
		ap.sl(s)
		prevNode.sl = ap.node
	}
	return root
}

func extendNode(ap *ActivePoint, i int, s string, root *STreeNode) {
	if ap.isEndOfNode() {
		add(ap.node, s[i], i, len(s)-1)
		return
	}
	ap.node = split(ap.node, root, ap.l, s)
	add(ap.node, s[i], i, len(s)-1)
}

func split(node, root *STreeNode, to int, s string) *STreeNode {
	newNode := &STreeNode{
		childs: make(map[byte]*STreeNode),
		from:   node.from,
		to:     node.from + to,
		parent: node.parent,
		sl:     root,
	}
	node.parent.childs[s[node.from]] = newNode
	node.parent = newNode
	node.from = node.from + to + 1
	newNode.childs[s[node.from]] = node
	return newNode
}

func add(node *STreeNode, c byte, from, to int) {
	node.childs[c] = &STreeNode{
		childs: make(map[byte]*STreeNode),
		from:   from,
		to:     to,
		parent: node,
	}
}

func isRoot(node *STreeNode) bool {
	return node.to == -1
}
