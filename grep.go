package grep

import "regexp"

// String returns a list of matches from the pattern you pass in
func String(pattern string, list []string) []string {
	return _string(pattern, list, false)
}

// NotString returns a list of strings that DONT match the pattern you pass in
func NotString(pattern string, list []string) []string {
	return _string(pattern, list, true)
}

func _string(pattern string, list []string, negateMatch bool) []string {
	var matches []string

	re := regexp.MustCompile(pattern)

	for _, item := range list {
		if re.Match([]byte(item)) {
			if negateMatch == false {
				matches = append(matches, item)
			}
		} else if negateMatch {
			matches = append(matches, item)
		}
	}

	return matches
}
