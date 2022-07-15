package utils

// FindDuplicatesFromStringSlice finds the duplicate values present in the given string slice
func FindDuplicatesFromStringSlice(values []string) []string {
	presentValues := make(map[string]bool)
	duplicateValues := make([]string, 0)
	for _, value := range values {
		if present := presentValues[value]; !present {
			presentValues[value] = true
		} else {
			duplicateValues = append(duplicateValues, value)
		}
	}
	return duplicateValues
}
