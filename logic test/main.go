package main

import "fmt"

func countAlphabets(word string) map[rune]int {
	counts := make(map[rune]int)
	for _, char := range word {
		counts[char]++
	}
	return counts
}

func compareMaps(a, b map[rune]int) bool {
	if len(a) != len(b) {
		return false
	}
	for key, value := range a {
		if bValue, exists := b[key]; !exists || bValue != value {
			return false
		}
	}
	return true
}

func anagrams(input []string) [][]string {
	anagrams := make([][]string, 0)
	temporaries := make(map[string]map[rune]int)

	for _, word := range input {
		temporaries[word] = countAlphabets(word)
	}

	for key, value := range temporaries {
		compareResult := false
		for index, words := range anagrams {
			compareResult = compareMaps(temporaries[words[0]], value)
			if compareResult {
				anagrams[index] = append(anagrams[index], key)
				break
			}
		}
		if !compareResult {
			anagrams = append(anagrams, []string{key})
		}
	}

	return anagrams
}

func main() {
	input := []string{"cook", "save", "taste", "aves", "vase", "state", "map"}
	result := anagrams(input)
	fmt.Println("Result : ", result)
}
