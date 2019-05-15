package fetl

import (
	"bufio"
	"os"
)

type FileETL struct {
	Filename  string
	Extract   ExtractFunc
	Transform TransformFunc
	Load      LoadFunc
}

type TransformFunc func(extracted interface{}) (tranformed interface{}, err error)

type LoadFunc func(tranformed interface{}) (err error)

type ExtractFunc func(text string) (interface{}, error)

// Start reading
func (e *FileETL) Start() (err error) {
	file, err := os.Open(e.Filename)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		var extracted interface{}
		var tranformed interface{}

		text := scanner.Text()

		extracted, err = e.Extract(text)
		if err != nil {
			return
		}

		if e.Transform != nil {
			tranformed, err = e.Transform(extracted)
			if err != nil {
				return
			}
		} else {
			tranformed = extracted
		}

		err = e.Load(tranformed)
		if err != nil {
			return
		}
	}

	return scanner.Err()
}
