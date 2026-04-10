package utils

func GetExtension(filename string) string {
	for i := len(filename) - 1; i >= 0; i-- {
		if filename[i] == '.' {
			return filename[i+1:] // Return everything after the dot
		}
	}
	return "" // Return an empty string if no dot is found
}
