package file

import (
	"encoding/csv"
	"git.darknebu.la/bit/logplus"
	"os"
)

type File struct {
	f *os.File
}

func Open(path string) (File, error) {
	file, err := os.Open(path)
	if err != nil {
		logplus.LogFError("openStarsCSV Panic! (cannot read file from %s)", path)
	}
	return File{f: file}, err
}

func (file *File) ReadCSV() ([][]string, error) {
	lines, err := csv.NewReader(file.f).ReadAll()
	if err != nil {
		logplus.LogError("openStarsCSV Panic! (cannot read the files content)")
	}
	return lines, err
}

func (file *File) Close() error {
	return file.f.Close()
}
