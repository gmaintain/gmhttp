package gmhttp

import (
	"strings"
)

type node struct {
	pattern  string
	part     string
	children []*node
	isWild   bool
}

func (n *node) Insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern
		return
	}
	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}
	child.Insert(pattern, parts, height+1)
}

func (n *node) Search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}
	part := parts[height]
	children := n.matchChildren(part)

	for _, child := range children {
		result := child.Search(parts, height+1)
		// 疑问这段代码能走下来吗?
		if result != nil {
			return result
		}
	}
	return nil
}

func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.isWild || child.part == part {
			return child
		}
	}
	return nil
}

func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

type noder interface {
	Insert(pattern string, parts []string, height int)
	Search(parts []string, height int) *node
}
