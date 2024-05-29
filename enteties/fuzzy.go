package modules

type FuzzyNode struct {
	Value      int
	Membership float64
	Left       *FuzzyNode
	Right      *FuzzyNode
}

type BSTFuzzy struct {
	Root *FuzzyNode
}

func (bst *BSTFuzzy) InsertFuzzyNode(node *FuzzyNode, value int, membership float64) *FuzzyNode {
	if bst.Root == nil {
		bst.Root = &FuzzyNode{value, membership, nil, nil}
		return bst.Root
	}
	if node == nil {
		return &FuzzyNode{value, membership, nil, nil}
	}
	if membership <= node.Membership {
		node.Left = bst.InsertFuzzyNode(node.Left, value, membership)
	}
	if membership >= node.Membership {
		node.Right = bst.InsertFuzzyNode(node.Right, value, membership)
	}
	return node
}

func (bst *BSTFuzzy) SearchFuzzyNode(node *FuzzyNode, value int, membership float64) bool {
	if node.Value == value && node.Membership == membership {
		return true
	}
	if node == nil {
		return false
	}
	if membership < node.Membership {
		return bst.SearchFuzzyNode(node.Left, value, membership)
	}
	if membership > node.Membership {
		return bst.SearchFuzzyNode(node.Right, value, membership)
	}
	return false
}
