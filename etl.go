package fetl

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

func (e *FileETL) SetExtractor(extractor Extractor) {
	e.Extract = extractor.Extract
}

func (e *FileETL) SetTransformer(transformer Transformer) {
	e.Transform = transformer.Transform
}

func (e *FileETL) SetLoader(loader Loader) {
	e.Load = loader.Load
}

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
		var tranformed fmt.Stringer

		text := scanner.Text()

		extracted, err = e.Extract(text)
		if err != nil {
			return
		}

		tranformed, err = e.Transform(extracted)
		if err != nil {
			return
		}

		err = e.Load(tranformed)
		if err != nil {
			return
		}
	}

	return scanner.Err()
}
