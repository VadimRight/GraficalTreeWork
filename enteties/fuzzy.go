package enteties

import (
	"errors"
	"fmt"
	"log"
)

type NodeFuzzy struct {
	Value      int
	Membership float64
	Left       *NodeFuzzy
	Right      *NodeFuzzy
}

type BSTFuzzy struct {
	Root *NodeFuzzy
}

func (bst *BSTFuzzy) InsertFuzzy(node *NodeFuzzy, value int, membership float64) {
	const opt string = "enteties.fuzzy.InsertFuzzy"
	if membership > 1.0 {
		err := errors.New("Membership can't be bigger than 1.0")
		log.Fatalf("%s: %v", opt, err)
	}
	bst.InsertNodeFuzzy(bst.Root, value, membership)
}

func (bst *BSTFuzzy) SearchFuzzy(node *NodeFuzzy, value int, membership float64) bool {
	found := bst.SearchNodeFuzzy(node, value, membership)
	return found
}

func (bst *BSTFuzzy) InsertNodeFuzzy(node *NodeFuzzy, value int, membership float64) *NodeFuzzy {
	if bst.Root == nil {
		bst.Root = &NodeFuzzy{value, membership, nil, nil}
		return bst.Root
	}
	if node == nil {
		return &NodeFuzzy{value, membership, nil, nil}
	}
	if membership <= node.Membership {
		node.Left = bst.InsertNodeFuzzy(node.Left, value, membership)
	}
	if membership >= node.Membership {
		node.Right = bst.InsertNodeFuzzy(node.Right, value, membership)
	}
	return node
}

func (bst *BSTFuzzy) SearchNodeFuzzy(node *NodeFuzzy, value int, membership float64) bool {
	if node.Value == value && node.Membership == membership {
		return true
	}
	if node == nil {
		return false
	}
	if membership < node.Membership {
		return bst.SearchNodeFuzzy(node.Left, value, membership)
	}
	if membership > node.Membership {
		return bst.SearchNodeFuzzy(node.Right, value, membership)
	}
	return false
}

func (bst *BSTFuzzy) InOrder(node *NodeFuzzy) {
	if node == nil {
		return
	} else {
		bst.InOrder(node.Left)
		fmt.Print(node.Value, node.Membership, " ")
		bst.InOrder(node.Right)
	}
}

func (bst *BSTFuzzy) Levelorder() {
	if bst.Root == nil {
		return
	}
	nodeList := make([](*NodeFuzzy), 0)
	nodeList = append(nodeList, bst.Root)
	for !(len(nodeList) == 0) {
		current := nodeList[0]
		fmt.Print(current.Value, " ")
		if current.Left != nil {
			nodeList = append(nodeList, current.Left)
		}
		if current.Right != nil {
			nodeList = append(nodeList, current.Right)
		}
		nodeList = nodeList[1:]
	}
}
