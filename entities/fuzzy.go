package entities

import (
	"fmt"
)

type NodeFuzzy struct {
	Value int
	Left  *NodeFuzzy
	Right *NodeFuzzy
}

type BSTFuzzy struct {
	Root *NodeFuzzy
}

func (bst *BSTFuzzy) InsertFuzzy(value int) {
	const opt string = "entities.InsertFuzzy"
	if bst.Root == nil {
		bst.Root = &NodeFuzzy{value, nil, nil}
	} else {
		bst.Root = bst.insertNodeFuzzy(bst.Root, value)
	}
}

func (bst *BSTFuzzy) insertNodeFuzzy(node *NodeFuzzy, value int) *NodeFuzzy {
	if bst.Root == nil {
		bst.Root = &NodeFuzzy{value, nil, nil}
		return bst.Root
	}
	if node == nil {
		return &NodeFuzzy{value, nil, nil}
	}
	if value <= node.Value {
		node.Left = bst.insertNodeFuzzy(node.Left, value)
	}
	if value > node.Value {
		node.Right = bst.insertNodeFuzzy(node.Right, value)
	}
	return node
}

// Defuzzify выполняет дефазификацию методом центра тяжести
func Defuzzify(y, mu []float64) int {
	fmt.Printf("Дефазификация для Y: %v и Mu: %v\n", y, mu)
	numerator := 0.0
	denominator := 0.0
	for i := 0; i < len(y); i++ {
		numerator += y[i] * mu[i]
		denominator += mu[i]
	}
	if denominator == 0 {
		return 0
	}
	result := numerator / denominator
	fmt.Printf("Результат дефазификации: %.2f\n", result)
	return int(result)
}

// PrettyPrint prints the tree in a visually appealing way
func (bst *BSTFuzzy) PrettyPrint() {
	bst.printHelper(bst.Root, "", true)
}

func (bst *BSTFuzzy) printHelper(node *NodeFuzzy, prefix string, isTail bool) {
	if node == nil {
		return
	}
	fmt.Printf("%s", prefix)
	if isTail {
		fmt.Printf("└── ")
		prefix += "    "
	} else {
		fmt.Printf("├── ")
		prefix += "│   "
	}
	fmt.Printf("%d\n", node.Value)
	bst.printHelper(node.Left, prefix, node.Right == nil)
	bst.printHelper(node.Right, prefix, true)
}
