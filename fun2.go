package main

import "strings"

func contarVogais(s string) int {
	vogais := []string{"a", "e", "i", "o", "u"}
	count := 0

	for _, char := range s {
		if contains(vogais, strings.ToLower(string(char))) {
			count++
		}
	}

	return count
}
