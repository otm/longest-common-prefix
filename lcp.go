// Package lcp (Longest Common Prefix) is used for text compleation
package lcp

// Find returns an array of strings that matches the longest needle
func Find(needle string, haystack []string) ([]string, error) {
	result := make([]string, len(haystack))
	temp := make([]string, 0, len(result))
	copy(result, haystack)

	for i, l := 0, len(needle); i < l; i++ {
		for j := 0; j < len(result); j++ {
			if i >= len(result[j]) {
				continue
			}

			if needle[i] == result[j][i] {
				temp = append(temp, result[j])
			}
		}

		if len(temp) == 0 {
			if i == 0 {
				return make([]string, 0), nil
			}

			return result, nil
		}

		result, temp = temp, make([]string, 0, len(result))
	}

	return result, nil
}
