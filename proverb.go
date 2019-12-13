// Package proverb match an input string to its proverb
package proverb

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {
	/*
		Given a list of inputs, generate the relevant proverb. For example, given the list
		["nail", "shoe", "horse", "rider", "message", "battle", "kingdom"], you will output the full text of this proverbial rhyme:

		For want of a nail the shoe was lost.
		For want of a shoe the horse was lost.
		For want of a horse the rider was lost.
		For want of a rider the message was lost.
		For want of a message the battle was lost.
		For want of a battle the kingdom was lost.
		And all for the want of a nail.

		For want of a s[0] the s[1] was lost.
		For want of a s[1] the s[2] was lost.
	*/

	head, mid, tail := "For want of a", "the", "was lost."

	var result []string

	result = append(result, head, "nail", mid, "shoe", tail)

	return result
}
