package main

import (
	"fmt"
	"log"

	"github.com/VadimRight/GraficalTreeWork/entities"
	"github.com/VadimRight/GraficalTreeWork/internal/parsers"
)

func main() {
	// Создание пустого дерева
	tree := &entities.BSTFuzzy{}

	// Чтение и добавление нечетких чисел из файла input.txt
	err := parsers.ParseAndInsertFuzzyNumbers("input.txt", tree, true)
	if err != nil {
		log.Fatalf("Error reading fuzzy numbers: %v", err)
	}

	fmt.Println("Дерево до добавления нового узла:")
	tree.PrettyPrint()

	// Чтение и добавление нечеткого числа для нового узла из файла new_node.txt
	err = parsers.ParseAndInsertFuzzyNumbers("new_node.txt", tree, false)
	if err != nil {
		log.Fatalf("Error reading new node fuzzy number: %v", err)
	}

	fmt.Println("\nДерево после добавления нового узла:")
	tree.PrettyPrint()
}
