// challenges/packages/begin/proverbs.go
package main

import (
	"fmt"
	"github.com/jboursiquot/go-proverbs"
)

// import the proverbs package

// getRandomProverb returns a random proverb from the proverbs package

func getRandomProverv() string {
	return proverbs.Random().Saying
}
func main() {
	fmt.Println(getRandomProverv())
}
