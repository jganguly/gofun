package iq

import (
	"fmt"
	"strings"
)

// splits a hashtag into words.
// exclusions: no word in dictionary can be a substring of another word
func HashtagIter(dict []string, ht string) []string {

	var wl = make([]string, 0)

	for i := 0; i < len(ht); i++ {
		if len(ht) == 0 {
			return wl
		} else {
			w, pos := hashtag(dict, ht)
			wl = append(wl, w)
			ht = ht[pos:]
			fmt.Println(w, ht, wl)
		}
	}
	fmt.Println("###", wl, ht)
	return wl
}

func hashtag(dict []string, ht string) (string, int) {
	for i := 0; i < len(ht)+1; i++ {
		sub := ht[0:i]
		if contains(dict, sub) {
			//fmt.Println("#", sub,wl)
			return sub, i
		}
	}
	return "", 0
}

func HashtagRecurse(listOfWords []string, ht string, pos int, tokens []string) []string {

	if pos == len(ht)-1 {
		return tokens
	} else {
		for i := pos; i < len(ht)+1; i++ {
			subword := ht[pos:i]
			if contains(listOfWords, subword) {
				tokens = append(tokens, subword)
				pos = i
				HashtagRecurse(listOfWords, ht, pos, tokens)
			}
		}
		return tokens
	}
}

func contains(wl []string, s string) bool {
	for _, w := range wl {
		if strings.Compare(w, s) == 0 {
			return true
		}
	}
	return false
}
