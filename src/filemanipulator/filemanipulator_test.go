package filemanipulator

import (
	"os"
	"testing"
)

type fakeFileReader struct {
	openFileForReadingSpy int
	readFromFileSpy       int
	fakeData              []string
	fakeFile              *os.File
}

func (fr *fakeFileReader) OpenFileForReading(fileName string) (*os.File, error) {
	fr.openFileForReadingSpy++
	return fr.fakeFile, nil
}

func (fr *fakeFileReader) ReadFromFile(file *os.File) ([]string, error) {
	fr.readFromFileSpy++
	return fr.fakeData, nil
}

func TestCanSortElementsOfATextFile(t *testing.T) {
	fakeData := []string{"shoe", "blanket", "horse", "zebra", "yellow"}
	expectedData := []string{"blanket", "horse", "shoe", "yellow", "zebra"}
	fileName := "fake_file.txt"
	var fakeFile *os.File
	ffr := fakeFileReader{
		openFileForReadingSpy: 0,
		readFromFileSpy:       0,
		fakeData:              fakeData,
		fakeFile:              fakeFile,
	}

	actual, err := SortFileElements(&ffr, fileName)

	checkError(err, t)

	compareStringSlices(actual, expectedData, t)
}

func checkError(err error, t *testing.T) {
	if err != nil {
		t.Error(err)
	}
}
func compareStringSlices(sliceA []string, sliceB []string, t *testing.T) {
for fileElement := range sliceB {
	if sliceA[fileElement] != sliceB[fileElement] {
		t.Error("Function didn't work")
	}
}
}