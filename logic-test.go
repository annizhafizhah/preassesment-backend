package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strings []string) [][]string {
	anagramGroups := make([][]string, 0)

	for _, str := range strings {
		// Membuat representasi yang terurut dari setiap string sebagai kunci untuk grup anagram
		sortedString := sortString(str)

		// Mencari grup anagram yang sudah ada
		groupFound := false
		for i, group := range anagramGroups {
			if sortString(group[0]) == sortedString {
				anagramGroups[i] = append(group, str)
				groupFound = true
				break
			}
		}

		// Jika grup anagram belum ada, maka membuat grup baru
		if !groupFound {
			anagramGroups = append(anagramGroups, []string{str})
		}
	}

	return anagramGroups
}

func sortString(str string) string {
	// Mengubah string menjadi slice rune dan mengurutkannya
	runes := []rune(str)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})

	// Mengembalikan string hasil pengurutan
	return string(runes)
}

func main() {
	strings := []string{"cook", "save", "taste", "aves", "vase", "state", "map"}
	result := groupAnagrams(strings)
	fmt.Println(result)
}
