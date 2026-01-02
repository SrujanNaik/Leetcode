package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}


	letter := make(map[string]int)
	for _, vari := range s {
		letter[string(vari)] += 1
	}

	for _, vari := range t {
		if letter[string(vari)] == 0 {
			return false
		}
		letter[string(vari)] -= 1
	}
	return true
}

func main () {
	s := "racecar"
	t := "arrace"

	fmt.Println(isAnagram(s, t))
}
