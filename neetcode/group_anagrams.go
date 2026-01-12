package main

import "fmt"


func group_Anagrams(strs []string) [][]string {
	m := make(map[[26]int][]string)

	for _, s := range strs {
		
		var count [26]int

		for i := 0; i < len(s); i++ {
			count[s[i] - 'a']++
		}

		m[count] = append(m[count], s)
	}

	result := [][]string{}
	for _, group := range m {
		result = append(result, group)
	}
	
	return result
}


func main() {
	strs := []string{"act", "pots", "tops", "cat", "stop", "hat"}

	//fmt.Println(string(strs[0][2]))
	fmt.Println(group_Anagrams(strs))
}
