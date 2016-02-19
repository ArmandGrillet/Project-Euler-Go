/*
Using names.txt (right click and 'Save Link/Target As...'), a 46K text file containing over five-thousand first names, begin by sorting it into alphabetical order.
Then working out the alphabetical value for each name, multiply this value by its alphabetical position in the list to obtain a name score.

For example, when the list is sorted into alphabetical order, COLIN, which is worth 3 + 15 + 12 + 9 + 14 = 53, is the 938th name in the list. So, COLIN would obtain a score of 938 × 53 = 49714.

What is the total of all the name scores in the file?
*/

package algo

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func E22() {
	dat, err := ioutil.ReadFile("files/22.txt")
	if err != nil {
		panic(err)
	}
	namesList := string(dat)
	namesList = strings.Replace(namesList, "\"", "", -1)

	names := strings.Split(namesList, ",")
	sort.Strings(names)

	namesScore := 0
	for index, name := range names {
		namesScore += nameScore(name) * (index + 1)
	}
	fmt.Println(namesScore)
}

func nameScore(name string) (score int) {
	for _, char := range name {
		score += (int)(char) - 64
	}
	return
}
