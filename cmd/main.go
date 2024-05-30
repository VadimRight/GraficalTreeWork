package main

import (
	"fmt"
	"log"

	"github.com/VadimRight/Grafical-Tree-Work/entities"
	"github.com/VadimRight/Grafical-Tree-Work/internal/parsers"
)

func main() {
	// Создание пустого дерева
	tree := &entities.BSTFuzzy{}

	// Чтение нечетких чисел из файла input.txt
	fuzzyTree, err := parsers.ParseTxtFileFuzzy("input.txt")
	if err != nil {
		log.Fatalf("Error reading fuzzy numbers: %v", err)
	}
	tree.Root = fuzzyTree.Root

	fmt.Println("Дерево до добавления нового узла:")
	tree.PrettyPrint()

	// Чтение нечеткого числа для нового узла из файла new_node.txt
	newFuzzyTree, err := parsers.ParseTxtFileFuzzy("new_node.txt")
	if err != nil {
		log.Fatalf("Error reading new node fuzzy number: %v", err)
	}

	// Добавление нового узла
	if newFuzzyTree.Root != nil {
		fmt.Println("\nЧтение и добавление нового узла:")
		tree.InsertFuzzy(newFuzzyTree.Root.Value)
	}

	fmt.Println("\nДерево после добавления нового узла:")
	tree.PrettyPrint()
}
