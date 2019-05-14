package fetl

type Extractor interface {
	Extract(text string) (interface{}, error)
}

type ExtractFunc func(text string) (interface{}, error)
