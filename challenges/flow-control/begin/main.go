// challenges/flow-control/begin/main.go
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"unicode"

	"github.com/davecgh/go-spew/spew"
)


func panicHandler(){
	if err:= recover(); err != nil{
		fmt.Println("Recovered from panic:", err)
	}
}

func errorValidaror(err error){
	if err != nil{
		panic(fmt.Errorf("Error wraper: %s", err))
	}
}
func main() {
	// handle any panics that might occur anywhere in this function
	defer panicHandler()

	// use os package to read the file specified as a command line argument
	//
	filepath := os.Args[1]
	file, err := os.Open(filepath)
	errorValidaror(err)
	defer file.Close()
	// convert the bytes to a string
	contentBytes, err := ioutil.ReadAll(file)
	errorValidaror(err)

	content := string(contentBytes)
	// initialize a map to store the counts
	countMap := make(map[string]int)

	// use the standard library utility package that can help us split the string into words
	words := strings.Fields(content)

	// capture the length of the words slice
	countMap["words"] = len(words)
	countMap["letters"] = 0
	countMap["numbers"] = 0
	countMap["symbols"] = 0
 	// use for-range to iterate over the words slice and for each word, count the number of letters, numbers, and symbols, storing them in the map
	for _, word := range words {
		for _, char := range word {
			if unicode.IsLetter(char){
				countMap["letters"]++
			}else if unicode.IsDigit(char){
				countMap["numbers"]++
			}else if unicode.IsPunct(char) || unicode.IsSymbol(char){
				countMap["symbols"]++
			}
		}
	}

	// dump the map to the console using the spew package
	spew.Dump(countMap)
}
