package filemanipulator

import (
	"duplicate_word_remover/filereader"
	"sort"
)

//SortFileElements to be used elsewhere
func SortFileElements(fr filereader.Interface, fileName string) ([]string, error) {
	file, err := fr.OpenFileForReading(fileName)

	if err != nil {
		return nil, err
	}

	data, err := fr.ReadFromFile(file)

	if err != nil {
		return nil, err
	}

	sort.Strings(data)

	return data, nil
}
