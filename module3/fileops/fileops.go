package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteFloatToFile(fileName string, value float64) {
	valueStr := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueStr), 0644)
}

func ReadFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return 0.0, errors.New("Failed to find file")
	}
	valueText := string(data)
	value, _ := strconv.ParseFloat(valueText, 64)
	return value, nil
}
