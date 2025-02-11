package parser

import (
	"github.com/tjhorner/grammar/sequence"
)

/*
Keyword: hear/board
Src: (I am|* are|* be) _hear_
Dst: ... [here]
Src: (I am|* are|* be) _board_ with
Dst: ... [bored] with

Removed: they {are board of} directors
*/

func (r *ruleMatcher) rule_beNoun(cur *sequence.Word) {
	var repl, why string

	switch cur.Lower {
	// default: return
	case "hear":
		repl = "here"
		why = "I am ‘here’ to ‘hear’"
	case "board":
		if !r.HasNextCont(1) ||
			r.NextWord(1).Lower != "with" {
			return
		}
		repl = "bored"
		why = "‘board’ is a noun; ‘bored’ is a verb"
	}

	if !r.HasPrevCont(2) {
		return
	}
	prev1 := r.PrevWord(1)
	switch {
	default:
		return
	case prev1.Lower == "am":
		prev2 := r.PrevWord(2)
		if prev2.Lower == "i" {
			prev2.MarkCommon()
		} else {
			return
		}
	case prev1.IsCopula():
	}

	r.Matched(cur.Lower, why)
	prev1.MarkCommon()
	cur.Replace(repl)
	return
}
