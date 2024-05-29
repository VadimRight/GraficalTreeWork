package modules

type Node struct {
	Value      int
	Membership float64
	Left       *Node
	Right      *Node
}

type BST struct {
	Root *Node
}

func (bst *BST) InsertNode(node *Node, value int, membership float64) *Node {
	if bst.Root == nil {
		bst.Root = &Node{value, membership, nil, nil}
		return bst.Root
	}
	if node == nil {
		return &Node{value, membership, nil, nil}
	}
	if membership <= node.Membership {
		node.Left = bst.InsertNode(node.Left, value, membership)
	}
	if membership >= node.Membership {
		node.Right = bst.InsertNode(node.Right, value, membership)
	}
	return node
}
