package filereader

import (
	"io/ioutil"
	"os"
	"testing"
)

var fileReader FileReader

func TestCanOpenFile(t *testing.T) {
	tempFile, err := ioutil.TempFile("", "tempFile")
	defer os.Remove(tempFile.Name())

	checkError(err, t)

	message := []byte("howyadoin")
	err = ioutil.WriteFile(tempFile.Name(), message, 0644)

	checkError(err, t)

	file, err := fileReader.OpenFileForReading(tempFile.Name())

	checkError(err, t)

	stat, err := file.Stat()

	if stat.Size() == 0 {
		t.Errorf("Error opening file")
	}

}

func TestCanReadFromFile(t *testing.T) {
	tempFile, err := ioutil.TempFile("", "tempFile")
	defer os.Remove(tempFile.Name())

	checkError(err, t)

	expected := "howyadoin"

	message := []byte(expected)

	err = ioutil.WriteFile(tempFile.Name(), message, 0644)

	actual, err := fileReader.ReadFromFile(tempFile)

	checkError(err, t)

	if actual[0] != expected {
		t.Errorf("Got %s, expected %s", actual[0], expected)
	}

}

func TestCanReadFromMultipleLinesInFile(t *testing.T) {
	tempFile, err := ioutil.TempFile("", "tempFile")
	defer os.Remove(tempFile.Name())

	testData := []byte("howdy yall\nhow y'all doin over here\ny'all have a good time, ya hear?")

	err = ioutil.WriteFile(tempFile.Name(), testData, 0644)

	actual, err := fileReader.ReadFromFile(tempFile)

	checkError(err, t)

	readError := "Error reading from file"

	if actual[0] != "howdy yall" {
		t.Error(readError)
	}

	if actual[1] != "how y'all doin over here" {
		t.Error(readError)
	}

	if actual[2] != "y'all have a good time, ya hear?" {
		t.Error(readError)

	}
}

//todo: create test to overwrite file with new []string

func checkError(err error, t *testing.T) {
	if err != nil {
		t.Error(err)
	}
}
