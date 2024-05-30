// /cmd/main.go

package main

import (
	"fmt"

	"github.com/VadimRight/Grafical-Tree-Work/internal/parsers"
)

func main() {
	filePath := "./txt-files/fuzzy.txt"
	treeFuzzy, err := parsers.ParseTxtFileFuzzy(filePath)
	if err != nil {
		fmt.Println("Error parsing file:", err)
		return
	}

	fmt.Println("InOrder Traversal:")
	treeFuzzy.InOrder(treeFuzzy.Root)
}
