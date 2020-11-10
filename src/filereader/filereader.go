package filereader

import (
	"bufio"
	"os"
)

//Interface to be used elsewhere
type Interface interface {
	OpenFileForReading(fileName string) (*os.File, error)
	ReadFromFile(file *os.File) ([]string, error)
}

func returnBiggerNumber(a int, b int) (result int) {
	if a == b {
		return
	} else if a > b {
		result = a
	} else {
		result = b
	}
	return
}

//FileReader to be used elsewhere
type FileReader struct{}

//OpenFileForReading blah blah blah
func (fr *FileReader) OpenFileForReading(fileName string) (*os.File, error) {
	file, err := os.Open(fileName)
	return file, err
}

//ReadFromFile blah blah
func (fr *FileReader) ReadFromFile(file *os.File) ([]string, error) {
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
