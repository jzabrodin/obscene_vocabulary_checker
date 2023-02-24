package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	tabooWords := extractTabooWords()

	isTabooWord := false

	for {
		var sentence string
		_, err := fmt.Scanln(&sentence)
		if err != nil {
			return
		}

		if strings.EqualFold(sentence, "exit") {
			fmt.Println("Bye!")
			break
		}

		words := strings.Split(sentence, " ")

		for _, word := range words {
			isTabooWord = tabooWords[strings.ToLower(word)] != ""
			tmp := ""
			if isTabooWord {
				for range sentence {
					tmp += "*"
				}
				sentence = strings.ReplaceAll(sentence, word, tmp)
			}
		}

		fmt.Println(sentence)
	}
}

func extractTabooWords() map[string]string {
	var filename string
	_, err := fmt.Scanln(&filename)
	if err != nil {
		panic("can't get filename")
	}

	file, err := os.Open(filename)
	if err != nil {
		panic("can't open file")
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	var tabooWords map[string]string
	tabooWords = make(map[string]string)

	for scanner.Scan() {
		tmp := scanner.Text()
		tabooWords[strings.ToLower(tmp)] = tmp
	}
	return tabooWords
}
