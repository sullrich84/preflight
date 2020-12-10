package array

func Contains(source []string, val string) bool {
	for _, item := range source {
		if item == val {
			return true
		}
	}
	return false
}

func ContainsAll(source []string, slice []string) bool {
	for _, item := range slice {
		if !Contains(source, item) {
			return false
		}
	}
	return true
}
