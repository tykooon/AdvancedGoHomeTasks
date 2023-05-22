package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"math"
	"os"
	"regexp"
	"sort"
	"strings"
)

type StatsDictionary map[string]int

func main() {

	var fileName string
	flag.StringVar(&fileName, "file", "", "file name (requiered parameter)")
	flag.StringVar(&fileName, "f", "", "file name (shorthand for -file)")

	var topNum int
	flag.IntVar(&topNum, "top", 10, "number of top results for output")
	flag.IntVar(&topNum, "t", 10, "(shorthand for -top)")

	if !isEntriesCorrect(&fileName, &topNum) {
		return
	}

	wordRegExp, _ := regexp.Compile(`\b\w+\b`)
	var dictionary StatsDictionary = make(map[string]int)

	file, _ := os.Open(fileName)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currentString := scanner.Text()
		currentString = strings.ToLower(currentString)
		dictionary.AddStatsFrom(currentString, wordRegExp)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Unexpected reading error.")
		return
	}

	var keys []string
	for w := range dictionary {
		keys = append(keys, w)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return dictionary[keys[i]] > dictionary[keys[j]]
	})

	for i := 0; i < int(math.Min(float64(topNum), float64(len(keys)))); i++ {
		fmt.Printf("%2d: %s -- %d \n", i+1, keys[i], dictionary[keys[i]])
	}
}

func isEntriesCorrect(f *string, t *int) bool {
	if len(os.Args) < 2 {
		fmt.Println("File name expected.")
		flag.Usage()
		return false
	}

	flag.Parse()

	if *f == "" {
		fmt.Println("Filename parameter is requiered. See using instructions:")
		flag.Usage()
		return false
	}

	if *t <= 0 {
		fmt.Println("top parameter cannot be negative. Default value will be used.")
		*t = 10
	}

	_, err := os.Stat(*f)
	if errors.Is(err, os.ErrNotExist) {
		fmt.Printf("Cannot find %s. Check the file name.", *f)
		return false
	} else if err != nil {
		fmt.Println("Unexpected error during access to file.")
		return false
	}
	return true
}

func (d StatsDictionary) AddStatsFrom(s string, re *regexp.Regexp) {
	words := re.FindAllString(s, -1)
	if words == nil {
		return
	}
	for _, w := range words {
		d[w] += 1
	}
}
