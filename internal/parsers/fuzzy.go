package parsers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/VadimRight/Grafical-Tree-Work/entities"
)

func ParseAndInsertFuzzyNumbers(filePath string, tree *entities.BSTFuzzy) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts)%2 != 0 {
			return fmt.Errorf("invalid line format: %s", line)
		}

		var y, mu []float64
		for i := 0; i < len(parts); i += 2 {
			value, err := strconv.ParseFloat(parts[i], 64)
			if err != nil {
				return err
			}
			membership, err := strconv.ParseFloat(parts[i+1], 64)
			if err != nil {
				return err
			}
			y = append(y, value)
			mu = append(mu, membership)
		}
		crispValue := entities.Defuzzify(y, mu)
		tree.InsertFuzzy(crispValue)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}