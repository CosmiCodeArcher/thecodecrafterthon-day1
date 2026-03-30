// CodeCrafters — Operation Gopher Protocol
// Module: String Transformer
// Author: [Hamza Musa]
// Squad:  [The Interfaces]

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// Your program accepts a command followed by a string and applies the correct transformation.

func main() {
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Println("Type your command and a string")
	line, _ :=  reader.ReadString('\n')
	line = strings.TrimSpace(line)

	parts := strings.Fields(line)
	

	command := parts[0]
	textSlice := parts[1:]
	text := strings.Join(textSlice, " ")

	if command == "upper" {fmt.Println(upper(text))}
	if command == "lower" {fmt.Println(lower(text))}
	if command == "cap" {fmt.Println(cap(text))}
	if command == "title" {fmt.Println(title(text))}
	if command == "snake" {fmt.Println(snake(text))}
	// if command == "reverse" {fmt.Println(reverse(text))}
}

func upper(s string) string {
	return strings.ToUpper(s)
}

func lower(s string) string {
	return strings.ToLower(s)
}

func cap(s string) string {
	words := strings.Fields(s)

	for i, word := range words {
		words[i] = strings.ToUpper(string(word[0])) + strings.ToLower(string(word[1:]))
	}
	return strings.Join(words, " ")
}

// [Threat, Level, of, Elevated,]


func title(s string) string {
	words := strings.Fields(s)
	connectors := []string{"a", "an", "the", "and", "but", "or", "for", "nor", "on", "at", "to", "by", "in", "of", "up", "as", "is", "it"}

	for i, word := range words {
		for _, c := range connectors {
			if word == c && word != string(word[0]) {
				words[i] = word
			} else {
				words[i] = strings.ToUpper(string(word[0])) + strings.ToLower(string(word[1:]))
			}
		}
	}

	return strings.Join(words, " ")
}

func snake(s string) string {
	words := strings.Fields(s)

	for i, word := range words {
		newSlice := []string{}
		newString := ""
		for _, v := range word {
			if unicode.IsPunct(v) {
				continue
			} else {
				newSlice = append(newSlice, string(v))
			}
			newString = strings.Join(newSlice, "")
			words[i] = newString
		}
	}
	
	camelCase := strings.Join(words, "_")

	return strings.TrimSpace(strings.ToLower(camelCase))
}

func reverse(s string) string {
	words := strings.Fields(s)

	for i, word := range words {
		reverse := []string{}
		for j = len(word)-1; i >=0; j++{
			reverse = append(reverse, word[j])
			words[i] = 
		}
		
	}
}