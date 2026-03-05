package main

/*
 * 
 * Tokenize will be splitting text into words and then normalises them and remove very short tokens
 * TODO -- convert to lowercase and replace the punctuatins with digits and spaces
 * this will keep only the letters
 */

import (
	"regexp"
	"strings"
)

func Tokenize(text string) []string {
	text = strings.ToLower(text)
	re := regexp.MustCompile(`[^a-z]+`)
	text = re.ReplaceAllString(text, " ")
	
	//spliting it into words and optional thing we can do is to remove the stop words for simplicity
	words := strings.Fields(text)
	filtered := make([]string, 0, len(words))
	for _, w := range words {
		if len(w) >= 2 {
			filtered = append(filtered, w)
		}
	}
	
	return filtered
}