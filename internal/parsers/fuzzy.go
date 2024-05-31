package parsers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/VadimRight/GraficalTreeWork/entities"
)

func ParseAndInsertFuzzyNumbers(filePath string, tree *entities.BSTFuzzy, isInputFile bool) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCount := 0

	// Подсчет количества строк в файле
	for scanner.Scan() {
		lineCount++
	}

	// Проверка на минимальное количество строк для основного входного файла
	if isInputFile && lineCount < 7 {
		return fmt.Errorf("need more than 7")
	}

	// Сброс сканера для повторного чтения
	file.Seek(0, 0)
	scanner = bufio.NewScanner(file)

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
