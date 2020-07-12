package reuse

import (
	"bufio"
	"encoding/json"
	"fmt"
	"index/suffixarray"
	"os"
	"regexp"
	"strings"
)

// ReadLineByLineFromFile read all the lines from a file
func ReadLineByLineFromFile(filename string) []string {

	path, err := os.Getwd()
	fmt.Println(path)

	// open file
	file, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	ErrorCheck(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	// read the lines one by one
	listOfLines := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		listOfLines = append(listOfLines, line)
	}

	return listOfLines

}

// Tokenize split a line into words using a regex pattern
// pattern := "[\\:\\,\\.\\s]+"
func Tokenize(line string, pattern string) []string {
	//pattern := "[\\/\\:\\,\\.\\s]+"
	listOfwords := regexp.MustCompile(pattern).Split(line, -1)
	return listOfwords
}

// GetColPos Get the indices of occurances of a word in a string
func GetColPos(line string, word string) []int {
	line = strings.Trim(line, " ")
	sa := suffixarray.New([]byte(line))
	indices := sa.Lookup([]byte(word), -1)
	return indices
}

// Camelcase converts a string to camel case
func Camelcase(word string) string {
	return strings.ToUpper(word[0:1]) + word[1:]
}

// ErrorCheck check for any error
func ErrorCheck(err error) {
	if err != nil {
		panic(err)
	}
}

// MarshallUnmarshallExample Marshal/Unmarshal example
func MarshallUnmarshallExample() {

	type address struct {
		Name string
		City string
		Zip  int
	}
	addr := address{"Jaideep", "Hyderabad", 500084}
	text, _ := json.Marshal(addr)
	fmt.Println(string(text))

	str := address{}
	json.Unmarshal([]byte(text), &str)
	fmt.Println(str.Name)
}



