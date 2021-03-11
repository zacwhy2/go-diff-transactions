package array

// IndexOf looks for a string from an array of strings and returns its index
func IndexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1
}
