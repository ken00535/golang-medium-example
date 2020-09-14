package genny

// Filter filter slice
func Filter(source []int, callback func(int) bool) []int {
	var out []int
	for _, element := range source {
		if callback(element) {
			out = append(out, element)
		}
	}
	return out
}
