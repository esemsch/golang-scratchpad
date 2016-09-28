package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
	"strings"
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

	allLines := strings.Split(string(wholeFile), "\n")

	wordsWithCounts := make(map[string]int)
	pattern := regexp.MustCompile("^\\s*(\\w+)[.,;?!]\\s*")
	for _, l := range allLines {
		for _, w := range strings.Split(l, " ") {
			sanitizedWord := strings.ToLower(pattern.ReplaceAllString(w, "$1"))
			if len(sanitizedWord) > 0 {
				count, ok := wordsWithCounts[sanitizedWord]
				if ok {
					wordsWithCounts[sanitizedWord] = count + 1
				} else {
					wordsWithCounts[sanitizedWord] = 1
				}
			}
		}
	}

	wordCounts := WordCounts{}
	for word, count := range wordsWithCounts {
		wordCounts = append(wordCounts, WordCount{count: count, word: word})
	}

	sort.Sort(wordCounts)

	for _, wc := range wordCounts {
		fmt.Printf("%s %dx\n", wc.word, wc.count)
	}
}
