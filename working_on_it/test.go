package main 

import (
	"fmt"
	"strings"
)

func main() {
	split("        Hello World!        ")
}

func split(s string) []string {
	list := strings.Fields(s)
	finalList := []string{}
	for i, word := range list {	
		list[i] = strings.TrimSpace(word)
		list[i] = strings.ToLower(word)
		finalList = append(finalList, list[i])
	}
	fmt.Println(finalList)
	fmt.Println(finalList[0])
	fmt.Println(finalList[1])
	return list
}
