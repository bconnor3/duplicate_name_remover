package filereader

import (
	"io/ioutil"
	"testing"
)

func TestCanOpenFile(t *testing.T) {
	name, err := ioutil.TempDir("duplicate_word_remover", "tempDir")

	tempFile, err := ioutil.TempFile(name, "tempFile")

	if err != nil {
		t.Error(err)
	}

	message := []byte("howyadoin")
	err = ioutil.WriteFile(tempFile.Name(), message, 0644)

	if err != nil {
		t.Error(err)
	}

	file, err := openFileForReading(tempFile.Name())

	if err != nil {
		t.Error(err)
	}

	stat, err := file.Stat()

	if stat.Size() == 0 {
		t.Errorf("Error opening file")
	}

}

func TestCanReadFromFile(t *testing.T) {
	var openFileStruct openFile
	name, err := ioutil.TempDir("duplicate_word_remover", "tempDir")
	tempFile, err := ioutil.TempFile(name, "tempFile")

	if err != nil {
		t.Error(err)
	}

	expected := "howyadoin"

	message := []byte(expected)

	err = ioutil.WriteFile(tempFile.Name(), message, 0644)

	actual, err := readFromFile(openFileStruct, tempFile)

	if err != nil {
		t.Error(err)
	}

	if actual.contents[0] != expected {
		t.Errorf("Got %s, expected %s", actual.contents[0], expected)
	}

}
