package filereader

import "os"

func openFileForReading(fileName string) (*os.File, error) {
	file, err := os.Open(fileName)

	return file, err
}
