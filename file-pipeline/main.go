// CodeCrafters — Operation Gopher Protocol
// Module: File Pipeline
// Author: [Hamza Musa]
// Squad:  [The Interface]

/*```text
// ═══════════════════════════════════════════
// SQUAD PIPELINE CONTRACT
// Squad: [The Interface]
// ───────────────────────────────────────────
// Input line types:
//   [
        Lines in ALL CAPS
        Lines in all lowercase
        Lines with extra leading/trailing spaces
        Lines that are only dashes or blanks
       ]
//
// Transformation rules (in order):
//   1. [Convert ALL CAPS lines to Title Case]
//   2. [Convert all lowercase lines to uppercase]
//   3. [Trim all leading and trailing whitespace]
//   4. [Reverse the words in any line that contains the word REVERSE]
//   5. [Remove lines that are only dashes or blanks]
//
// Output format:
//   [Header: yes]
//   [Line numbering format: "1"]
//   [Summary block: yes]
//
// Terminal summary fields:
//   [Lines read    : [number]]
//   [Lines written : [number]]
//   [Lines removed : [number]]
//   [Rules applied : [list your 5 rules]]
// ═══════════════════════════════════════════
```
*/

package main

import (
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(input), "\n")
	for i := range lines {
		lines[i] = pipe(lines[i])
	}
	os.WriteFile("output.txt", []byte(strings.Join(lines, "\n")), 0644)
}

func pipe(s string) string {
	result := []string{}
	if s == strings.ToUpper(s) {
		s = cap(s)
	}
	if s == strings.ToLower(s) {
		s = strings.ToUpper(s)
	}
	if strings.Contains(s, "REVERSE") {
		s = reverse(s)
	}
	s = strings.TrimSpace(s)

	if strings.Trim(s, "-") != "" && strings.Trim(s, " ") != "" {
		result = append(result, s)
	}

	return s
}

func cap(s string) string {
	words := strings.Fields(s)

	for i, word := range words {
		words[i] = strings.ToUpper(string(word[0])) + strings.ToLower(string(word[1:]))
	}
	return strings.Join(words, " ")
}

func reverse(s string) string {
	words := strings.Fields(s)

	for i, word := range words {
		reverse := []string{}
		for j := len(word) - 1; j >= 0; j-- {
			reverse = append(reverse, string(word[j]))
		}
		words[i] = strings.Join(reverse, "")
	}
	return strings.Join(words, " ")
}
