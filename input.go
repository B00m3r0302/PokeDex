package main

import (
	"strings"
)

func CleanInput(text string) []string {
	list := strings.Fields(text)
	finalList := []string{}
	for i := range list {
		list[i] = strings.TrimSpace(list[i])
		list[i] = strings.ToLower(list[i])
		finalList = append(finalList, list[i])
	}
	return finalList
}
