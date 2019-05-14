package fetl

import "fmt"

type Loader interface {
	Load(tranformed fmt.Stringer) (err error)
}

type LoadFunc func(tranformed fmt.Stringer) (err error)
