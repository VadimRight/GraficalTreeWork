package modules

import (
	"errors"
	"log"
)

type FuzzyNode struct {
	Value      int
	Membership float64
	Left       *FuzzyNode
	Right      *FuzzyNode
}

type BSTFuzzy struct {
	Root *FuzzyNode
}

func (bst *BSTFuzzy) InsertFuzzy(node *FuzzyNode, value int, membership float64) {
	const opt string = "enteties.fuzzy.InsertFuzzy"
	if membership > 1.0 {
		err := errors.New("Membership can't be bigger than 1.0")
		log.Fatalf("%s: %v", opt, err)
	}
	bst.InsertFuzzyNode(bst.Root, value, membership)
}

func (bst *BSTFuzzy) SearchFuzzy(node *FuzzyNode, value int, membership float64) bool {
	found := bst.SearchFuzzyNode(node, value, membership)
	return found
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
