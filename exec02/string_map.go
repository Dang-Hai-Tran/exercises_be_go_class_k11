package main

func CreateMapByString(str string) map[string]int {
	var mapCharCount = make(map[string]int)
	for _, char := range str {
		mapCharCount[string(char)] += 1
	}

	return mapCharCount
}
