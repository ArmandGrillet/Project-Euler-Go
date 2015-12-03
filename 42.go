/*
The nth term of the sequence of triangle numbers is given by, tn = ½n(n+1); so the first ten triangle numbers are:

1, 3, 6, 10, 15, 21, 28, 36, 45, 55, ...

By converting each letter in a word to a number corresponding to its alphabetical position and adding these values we form a word value. For example, the word value for SKY is 19 + 11 + 25 = 55 = t10. If the word value is a triangle number then we shall call the word a triangle word.

Using words.txt (right click and 'Save Link/Target As...'), a 16K text file containing nearly two-thousand common English words, how many are triangle words?
*/

package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func isTriangleNumber(number int) bool {
	triangleNumber := 1
	for i := 1; triangleNumber < number; i++ {
		triangleNumber = int(float64(0.5*float64(i)) * float64(i+1))
	}
	if number == triangleNumber {
		return true
	} else {
		return false
	}
}

func main() {
	alphabet := strings.Split("abcdefghijklmnopqrstuvwxyz", "")
	valuedAlphabet := make(map[string]int)
	for i := 0; i < len(alphabet); i++ {
		valuedAlphabet[alphabet[i]] = i + 1
	}

	dat, err := ioutil.ReadFile("42.txt")
	if err != nil {
		panic(err)
	}
	text := string(dat)
	text = strings.Replace(text, "\"", "", -1)
	text = strings.ToLower(text)
	words := strings.Split(text, ",")

	triangleWords := 0
	var wordLetters []string
	var wordValue int
	for i := 0; i < len(words); i++ {
		wordLetters = strings.Split(words[i], "")
		wordValue = 0
		for j := 0; j < len(wordLetters); j++ {
			wordValue += valuedAlphabet[wordLetters[j]]
		}
		if isTriangleNumber(wordValue) {
			triangleWords++
		}
	}
	fmt.Println(triangleWords)
}
