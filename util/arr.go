package util

func Contains[c comparable](slice []c, value c) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}

func GetIndex[c comparable](slice []c, value c) int {
	for i, item := range slice {
		if item == value {
			return i
		}
	}

	return -1
}
