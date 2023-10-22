package main

import "strings"

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func concatenarStrings(strs []string) string {
	return strings.Join(strs, "")
}
