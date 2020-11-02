package filereader

import (
	"bufio"
	"os"
)

func openFileForReading(fileName string) (*os.File, error) {
	file, err := os.Open(fileName)
	return file, err
}

func readFromFile(file *os.File) ([]string, error) {
	defer file.Close()
	var readValue []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		readValue = append(readValue, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return readValue, nil
}
