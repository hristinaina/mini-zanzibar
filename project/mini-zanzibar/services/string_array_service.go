package services

func Contains(array []string, item string) bool {
	for _, i := range array {
		if i == item {
			return true
		}
	}
	return false
}

func HasUniqueElements(firstArray []string, secondArray []string) bool {
	for _, i := range secondArray {
		if !Contains(firstArray, i) {
			return true
		}
	}
	return false
}
