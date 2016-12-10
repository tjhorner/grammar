package main

import (
	"victorz.ca/grammar"

	"fmt"
)

func runTest(s string) (undetected bool) {
	r := grammar.MakeTweetReply(s, "@")
	if r == "" {
		r = "(No errors detected!)"
		undetected = true
	}
	fmt.Println(r)
	return
}

func main() {
	for _, test := range [...]string{
		"Nothing's wrong with this sentence.",
		"But it's true that their is a problem with this sentence.",
	} {
		for i := 0; i < 20; i++ {
			if runTest(test) {
				break
			}
		}
	}
}
