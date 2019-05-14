package fileo

import (
	"bufio"
	"fmt"
	"os"
)

// FileETL
type FileETL struct {
	Filename  string
	Extract   ExtractFunc
	Transform TransformFunc
	Load      LoadFunc
}

type ExtractFunc func(text string) (extracted interface{}, err error)

type TransformFunc func(extracted interface{}) (tranformed fmt.Stringer, err error)

type LoadFunc func(tranformed fmt.Stringer) (err error)

// Start reading
func (r FileETL) Start() (err error) {
	file, err := os.Open(r.Filename)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		var extracted interface{}
		var tranformed fmt.Stringer

		text := scanner.Text()

		extracted, err = r.Extract(text)
		if err != nil {
			return
		}

		tranformed, err = r.Transform(extracted)
		if err != nil {
			return
		}

		err = r.Load(tranformed)
		if err != nil {
			return
		}
	}

	return scanner.Err()
}
