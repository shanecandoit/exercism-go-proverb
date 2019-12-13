// Package proverb create a proverb that uses the given strings. Pairs from a list.
package proverb

import "strings"

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {
	/*
	Given a list of inputs: ["nail", "shoe", "horse", ...]
	output the full text of this proverbial rhyme:

		For want of a $1 the $2 was lost.
		For want of a $2 the $3 was lost.
		...
		And all for the want of a $1.
	*/

	line := "For want of a $1 the $2 was lost."
	last := "And all for the want of a $1."

	var result []string

	// empty input, empty output
	sz := len(rhyme)
	if sz == 0 {
		return result
	}

	for i, str := range rhyme {
		// and grab the next str too
		if i < sz-1 {
			str2 := rhyme[i+1]
			line1 := strings.ReplaceAll(line, "$1", str)
			line2 := strings.ReplaceAll(line1, "$2", str2)
			result = append(result, line2)
		}
	}

	// end with last line. "And all for the want of a $1." rhyme[0]
	result = append(result, strings.Replace(last, "$1", rhyme[0], 1))

	return result
}
