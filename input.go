package main

import (
	"strings"
)

func cleanInput(text string) []string {
	list := strings.Split(text, " ")
	finalList := []string{}
	for i := range list {
		if list[i] == "" || list[i] == " " {
			
		}
		list[i] = strings.TrimSpace(list[i])
		list[i] = strings.ToLower(list[i])
		finalList = append(finalList, list[i])
	}
	return finalList
}
