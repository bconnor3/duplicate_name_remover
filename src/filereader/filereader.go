package filereader

import (
	"bufio"
	"log"
	"os"
)

type openFile struct {
	contents []string
}

func openFileForReading(fileName string) (*os.File, error) {
	file, err := os.Open(fileName)

	return file, err
}

func readFromFile(of openFile, file *os.File) (openFile, error) {
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		of.contents = append(of.contents, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return of, nil
}
