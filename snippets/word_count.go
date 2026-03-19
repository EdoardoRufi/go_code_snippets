package snippets

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

// 1) count all strings and substrings in an array of strings given a string as input.
// 2) given the input strings, find all the consecutive substrings and count every of them
// 3) do it asynchronosly
// 4) asynchronously with wait group

var _ TaskExecutioner = (*WordCount)(nil)

type TaskExecutioner interface {
	Exec()
	Print()
}

type WordCount struct {
	WordsList []string
	Str       string
	Count     int
}

func NewWordCount(wordsList []string, str string) *WordCount {
	return &WordCount{WordsList: wordsList, Str: str}
}

// 1)
func (w *WordCount) Exec() {
	for _, word := range w.WordsList {
		if strings.Contains(word, w.Str) {
			w.Count++
		}
	}
	w.Print()
	time.Sleep(1 * time.Second)
}

func (w *WordCount) Print() {
	fmt.Println("WordList: ", w.WordsList)
	fmt.Println("Str to count: ", w.Str)
	fmt.Println("Count: ", w.Count)
}

// 2)
type WordCountSubstrings struct {
	All []*WordCount
}

func NewWordCountSubstrings(wordsList []string, str string) *WordCountSubstrings {

	const minSubstrLen = 2
	strLen := len(str)

	all := make([]*WordCount, 0)

	strChars := strings.Split(str, "")

	// j arriva da 2 a 4. Va addizionata all'indice da prendere
	for j := minSubstrLen; j <= strLen; j++ {

		//for che itera sui caratteri,  quindi va da 0 a 5. Dovrò prender il charIndex + j
		for charIndex := 0; charIndex <= len(strChars); charIndex++ {
			substring := ""
			if (charIndex + j) > len(strChars) {
				continue
			}
			// se parti da 2, prendi wordChars[0] e wordChars[1], concatena e salva una stringa
			for i := 0; i < j; i++ {
				if (charIndex + i) >= len(strChars) {
					continue
				} else {
					substring += strChars[charIndex+i]
				}
			}

			all = append(all, &WordCount{WordsList: wordsList, Str: substring})
		}
	}
	return &WordCountSubstrings{All: all}
}

// func (w *WordCountSubstrings) Exec() {
// 	for _, executer := range w.All {
// 		executer.Exec()
// 	}
// }

// 3) ASYNC EXEC
// func (w *WordCountSubstrings) Exec() {

// 	wg := &sync.WaitGroup{}

// 	for _, executer := range w.All {

// 		wg.Add(1)
// 		go func(e *WordCount) {
// 			defer wg.Done()
// 			e.Exec()

// 		}(executer)
// 	}
// 	wg.Wait()

// }

// 4) ASYNC EXEC with wait group
func (w *WordCountSubstrings) Exec() {

	const workerPoolSize = 3

	wordCountChannel := make(chan *WordCount)

	wg := &sync.WaitGroup{}
	wg.Add(workerPoolSize)

	for i := 0; i < workerPoolSize; i++ {
		go func() {
			defer wg.Done()
			for job := range wordCountChannel {
				job.Exec()
			}
		}()
	}

	for _, job := range w.All {
		wordCountChannel <- job
	}
	close(wordCountChannel)
	wg.Wait()

}
