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
	"fmt"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . <input_file> <output_file>")
		return
	}

	input, _ := os.ReadFile(os.Args[1])
	output := os.Args[2]

	rawLines := strings.Split(string(input), "\n")
	outputLines := []string{}
	header := fmt.Sprintf("(THE-INTERFACE SENTINEL FIELD REPORT — PROCESSED)\n")
	outputLines = append(outputLines, header)

	for _, line := range rawLines {
		transformed := pipe(line)
		if transformed != "" {
			outputLines = append(outputLines, transformed)
		}
	}

	summary := fmt.Sprintf("\n==================== \nLines read: %d\nLines written: %d\nLines removed: %d\nRules applied : [\n1. Convert ALL CAPS lines to Title Case, \n2. Convert all lowercase lines to uppercase, \n3. Trim all leading and trailing whitespace, \n4. Reverse the words in any line that contains the word REVERSE, \n5. Remove lines that are only dashes or blanks\n]", len(rawLines), len(outputLines), len(rawLines)-len(outputLines))
	outputLines = append(outputLines, summary)

	os.WriteFile(output, []byte(strings.Join(outputLines, "\n")), 0644)
}

func pipe(s string) string {
	if s == strings.ToUpper(s) {s = cap(s)}
	if s == strings.ToLower(s) {s = strings.ToUpper(s)}
	if strings.Contains(s, "REVERSE") {s = reverse(s)}

	s = strings.TrimSpace(s)

	if strings.Trim(s, " -") == "" {
		return ""
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
