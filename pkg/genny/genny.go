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

// FilterWithInterface filter slice
func FilterWithInterface(source interface{}, callback func(interface{}) bool) interface{} {
	var outInt []int
	var outFloat32 []float32
	switch source.(type) {
	case []int:
		for _, element := range source.([]int) {
			if callback(element) {
				outInt = append(outInt, element)
			}
		}
		return outInt
	case []float32:
		for _, element := range source.([]float32) {
			if callback(element) {
				outFloat32 = append(outFloat32, element)
			}
		}
		return outFloat32
	}
	return nil
}
