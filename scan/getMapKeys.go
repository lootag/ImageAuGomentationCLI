package scan

func(this ScanningService) getMapKeys(stringIntMap map[string]int) []string {
	keys := []string{}
	for key, _ := range stringIntMap {
		keys = append(keys, key)
	}

	return keys
}