package grep

import "regexp"

// String returns a list of matches from the pattern you pass in
func String(pattern string, list []string) []string {
	var matches []string

	re := regexp.MustCompile(pattern)

	for _, item := range list {
		if re.Match([]byte(item)) {
			matches = append(matches, item)
		}
	}

	return matches
}
