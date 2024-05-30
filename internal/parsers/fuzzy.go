package parsers

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/VadimRight/Grafical-Tree-Work/entities"
)

func ParseTxtFileFuzzy(filePath string) (*entities.BSTFuzzy, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	tree := &entities.BSTFuzzy{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		pairs := strings.Split(line, ";")
		for _, pair := range pairs {
			parts := strings.Split(pair, ",")
			if len(parts) != 2 {
				continue
			}
			value, err1 := strconv.Atoi(parts[0])
			membership, err2 := strconv.ParseFloat(parts[1], 64)
			if err1 != nil || err2 != nil {
				continue
			}
			tree.InsertFuzzy(tree.Root, value, membership)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return tree, nil
}
