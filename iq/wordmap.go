package iq

import (
	"fmt"
	"intques/reuse"
)

// Wordmap map of words containing line number and column position
func Wordmap(word string) {

	listOfLines := reuse.ReadLineByLineFromFile("iq/text.dat")

	// map of word
	mapW := make(map[string][]structLoc)

	for i, line := range listOfLines {

		pattern := "[\\/\\:\\,\\.\\s]+"
		listOfwords := reuse.Tokenize(line,pattern)

		for _, word := range listOfwords {
			indices := reuse.GetColPos(line, word)

			if mapW[word] == nil {
				arrStructLoc := make([]structLoc, 0)
				arrStructLoc = addToArrStructLoc(i, indices, arrStructLoc)
				mapW[word] = arrStructLoc
			} else {
				arrStructLoc := mapW[word]
				arrStructLoc = addToArrStructLoc(i, indices, arrStructLoc)
				mapW[word] = arrStructLoc
			}
		}
	}
	fmt.Println(mapW[word])
}

func addToArrStructLoc(lineno int, indices []int, arrStructLoc []structLoc) []structLoc {
	for _, index := range indices {
		loc := structLoc{lineno, index}
		arrStructLoc = append(arrStructLoc, loc)
	}
	return arrStructLoc
}
