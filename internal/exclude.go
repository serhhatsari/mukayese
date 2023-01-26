package internal

import "strings"

var excludeList = []string{".git"}

func Exclude(path string) bool {
	isExist := false
	for _, s := range excludeList {
		if strings.Contains(path, s) {
			isExist = true
		}
	}
	return isExist
}
