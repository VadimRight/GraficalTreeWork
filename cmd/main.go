// /cmd/main.go

package main

import (
	"fmt"

	"github.com/VadimRight/Grafical-Tree-Work/internal/parsers"
)

func main() {
	filePath := "./fuzzy.txt"
	tree, err := parsers.ParseTxtFile(filePath)
	if err != nil {
		fmt.Println("Error parsing file:", err)
		return
	}

	fmt.Println("InOrder Traversal:")
	tree.InOrder(tree.Root)
}
