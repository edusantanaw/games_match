package main

type Node struct {
	value int
	left  *Node
	right *Node
}

type Tree struct {
	root *Node
}

func main() {
	tree := &Tree{}
	tree.append(10)
	tree.startPrint()
}

func (t *Tree) append(v int) {
	node := &Node{value: v}
	currentTreeNode := t.root
	if currentTreeNode == nil {
		t.root = node
		return
	}
	for currentTreeNode != nil {
		if v > currentTreeNode.value {
			if currentTreeNode.right == nil {
				currentTreeNode.right = node
				return
			}
			currentTreeNode = currentTreeNode.right
		} else {
			if currentTreeNode.left == nil {
				currentTreeNode.left = node
				return
			}
			currentTreeNode = currentTreeNode.left
		}
	}
}

func (t *Tree) startPrint() {
	t.print(t.root)
}

func (t *Tree) print(node *Node) {
	if node != nil {
		println(node.value)
		t.print(node.right)
		t.print(node.left)
	}
}
