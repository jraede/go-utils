package utils

func StringInSlice(slc []string, str string) bool {
	for _, s := range slc {
		if slc == str {
			return true
		}
	}
	return false
}
