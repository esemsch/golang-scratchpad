package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
	"time"
	"unicode"
)

type WordCount struct {
	count int
	word  string
}

type WordCounts []WordCount

func (wc WordCounts) Len() int      { return len(wc) }
func (wc WordCounts) Swap(i, j int) { wc[i], wc[j] = wc[j], wc[i] }
func (wc WordCounts) Less(i, j int) bool {
	if wc[i].count > wc[j].count {
		return true
	} else if wc[i].count < wc[j].count {
		return false
	} else {
		return strings.Compare(wc[i].word, wc[j].word) < 0
	}
}

const MILLION int64 = 1000000

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Fprintln(os.Stderr, "You have to provide the name of the file")
		os.Exit(1)
	}

	fileName := args[1]

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't open file %s", fileName)
		os.Exit(1)
	}

	T := time.Now().UnixNano()
	wholeFile := []byte{}
	buffer := make([]byte, 4096)

	for {
		n, err := file.Read(buffer)
		if n == 0 && err == io.EOF {
			err := file.Close()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Couldn't close file %s", fileName)
				os.Exit(1)
			}
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading from file %s. Error: ", fileName, err)
			os.Exit(1)
		} else {
			wholeFile = append(wholeFile, buffer[0:n]...)
		}

	}
	fmt.Println("Reading", (time.Now().UnixNano()-T)/MILLION, "ms")

	T = time.Now().UnixNano()
	allLines := strings.Split(string(wholeFile), "\n")
	fmt.Println("Splitting", (time.Now().UnixNano()-T)/MILLION, "ms")

	T = time.Now().UnixNano()
	wordsWithCounts := make(map[string]int)
	punctuation := []rune{'.', ',', '!', '?', ';'}
	for _, l := range allLines {
		for _, w := range strings.Split(l, " ") {
			withoutPunctuation := make([]rune, len(w))
			pos := 0
			for _, r := range w {
				isPunctuation := false
				for _, p := range punctuation {
					if r == p {
						isPunctuation = true
						break
					}
				}
				if !isPunctuation {
					withoutPunctuation[pos] = unicode.ToLower(r)
					pos++
				}
			}
			if pos > 0 {
				sanitizedWord := string(withoutPunctuation[:pos])
				count, ok := wordsWithCounts[sanitizedWord]
				if ok {
					wordsWithCounts[sanitizedWord] = count + 1
				} else {
					wordsWithCounts[sanitizedWord] = 1
				}
			}
		}
	}
	fmt.Println("Counting", (time.Now().UnixNano()-T)/MILLION, "ms")

	T = time.Now().UnixNano()
	wordCounts := WordCounts{}
	for word, count := range wordsWithCounts {
		wordCounts = append(wordCounts, WordCount{count: count, word: word})
	}
	fmt.Println("Converting to list of structs", (time.Now().UnixNano()-T)/MILLION, "ms")

	T = time.Now().UnixNano()
	sort.Sort(wordCounts)
	fmt.Println("Sorting", (time.Now().UnixNano()-T)/MILLION, "ms")

	fmt.Println("==== RESULTS ====")
	for _, wc := range wordCounts {
		fmt.Printf("%s %dx\n", wc.word, wc.count)
	}
}
