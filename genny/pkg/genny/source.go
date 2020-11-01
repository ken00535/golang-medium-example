package genny

import "github.com/cheekybits/genny/generic"

//go:generate genny -in ./source.go -out ./gencode.go gen "Something=int,string,float32"

// NOTE: this is how easy it is to define a generic type
type Something generic.Type

//FilterWithSomething filter Something type
func FilterWithSomething(source []Something, callback func(Something) bool) []Something {
	var out []Something
	for _, element := range source {
		if callback(element) {
			out = append(out, element)
		}
	}
	return out
}
