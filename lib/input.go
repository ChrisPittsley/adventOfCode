package lib

import (
	"io"
	"os"
	"path/filepath"
)

func GetInput() ([]byte, error) {
	var in io.Reader
	switch len(os.Args) {
	case 1:
		in = os.Stdin
	case 2:
		path, err := filepath.Abs(os.Args[1])
		if err != nil {
			return nil, err
		}
		in, err = os.Open(path)
		if err != nil {
			return nil, err
		}
	}
	data, err := io.ReadAll(in)
	if err != nil {
		return nil, err
	}
	return data, nil
}