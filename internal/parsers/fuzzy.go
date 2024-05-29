package parsers

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/VadimRight/Grafical-Tree-Work/entities"
)

func ParseTxtFile(filePath string) (*entities.BSTFuzzy, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	tree := &entities.BSTFuzzy{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) != 2 {
			continue
		}
		value, _ := strconv.Atoi(parts[0])
		membership, _ := strconv.ParseFloat(parts[1], 64)
		tree.InsertFuzzy(tree.Root, value, membership)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return tree, nil
}
