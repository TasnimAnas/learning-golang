package fileSystem

import (
	"fmt"
	"os"
	"strconv"
)

func IsFileExist(fileName string) bool {
	_, err := os.Stat(fileName)
	return err == nil
}

func ConcatStringToFile(fileName string, data string) {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic("Error opening statement")
	}
	defer file.Close()
	_, err = file.WriteString(data)
	if err != nil {
		panic("Error writing to statement")
	}
}

func ReadFloatFromFile(fileName string) (data float64) {
	dataByte, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("New user, account initiated with balance 0!")
		return 0
	}

	dataString := string(dataByte)
	data, err = strconv.ParseFloat(dataString, 64)
	if err != nil {
		fmt.Println("Balance is corrupted, initiated with balance 0!")
		return 0
	}
	return data
}

func WriteFloatToFile(fileName string, data float64) {
	os.WriteFile(fileName, fmt.Append(nil, data), 0664)
}
