package tasks

import (
	"regexp"
	"strings"
)

//CheckDuplicateStrings
func CheckDuplicateStrings(str string) []string {
	duplicates := []string{}
	occurance := make(map[string]bool)

	splitStr := strings.Split(str, " ")

	for _, s := range splitStr {

		lowstr := strings.ToLower(s)

		//remove all numbers and special characters
		re := regexp.MustCompile(`[^\w0-9]`)
		cleanedStr := re.ReplaceAllString(lowstr, "")

		_, occcured := occurance[cleanedStr]
		if occcured {
			occurance[cleanedStr] = true
		} else {
			occurance[cleanedStr] = false
		}
	}

	for k, v := range occurance {
		if v {
			duplicates = append(duplicates, k)
		}
	}
	return duplicates
}
