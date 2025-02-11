package parser

import (
	"github.com/tjhorner/grammar/sequence"
)

/*
Keyword: whose
Src: _ been
Dst: [who's] been
*/

func (r *ruleMatcher) rule_whoseBeen(cur *sequence.Word) {
	if !r.HasNextCont(1) {
		return
	}
	next1 := r.NextWord(1)
	if next1.Lower != "been" {
		return
	}

	// Exception: whose BEEN (noun abbreviation)
	if next1.Caps == sequence.WC_UPPER {
		return
	}

	r.Matched("whose_has", "‘whose’ is possessive; ‘who's’ means ‘who has’")
	cur.ReplaceCap("who's")
	next1.MarkCommon()
}
