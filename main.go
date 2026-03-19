package main

import (
	"go_code_snippets/snippets"
)

func main() {
	//snippets.LaunchSortByPrice()
	//snippets.LaunchCheckDuplicates()
	// datastructures.LaunchStack()
	// datastructures.LaunchConcurrentStack()
	// snippets.LaunchFindMinMax()
	//snippets.EnsureMaxVersionInString()
	// datastructures.LaunchConstructorDemonstration()
	snippets.Error()
	// err := endpoints.StartServers()
	// if err != nil {
	// 	panic(err)
	// }

	wordsList := []string{
		"smiao",
		"miao",
		"smiaog",
		"mao",
		"iao",
	}

	// wordCount := snippets.NewWordCount(wordsList, "miao")
	// wordCount.Exec()

	wordCountSubstrings := snippets.NewWordCountSubstrings(wordsList, "miao")
	wordCountSubstrings.Exec()
}
