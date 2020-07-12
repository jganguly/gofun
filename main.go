package main

import (
	"fmt"
	"intques/iq"
)

func main() {
	/* hastag */
	ht := "#bindrawinsgold"
	ht = ht[1:]
	dict := []string {"wins", "loses", "gold", "silver", "bindra", "tom"}
	var tokens = make([]string,0)
	tokens = iq.HashtagIter(dict, ht)
	tokens = iq.HashtagRecurse(dict,ht,0,tokens)
	fmt.Println(tokens)
	
	/* wordmap */
	iq.Wordmap("of")
}

