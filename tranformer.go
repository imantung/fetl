package fetl

import "fmt"

type Transformer interface {
	Transform(extracted interface{}) (tranformed fmt.Stringer, err error)
}

type TransformFunc func(extracted interface{}) (tranformed fmt.Stringer, err error)
