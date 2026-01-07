package main 

import (
	"fmt"
	"strings"
)

func main() {
	split("Hello World!")
}
func split(s string) []string {
	list := strings.Split(s, " ")
	fmt.Println(list)
	fmt.Println(list[0])
	fmt.Println(list[1])
	return list
}
