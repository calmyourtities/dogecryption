// borkify
package dogecrypt

import (
	_ "fmt"
	_ "strings"
)

// this algorithim is open for improvement
// it's supposed to translate a string of random characters
// to a sequence of strings like "bork" and "ruff"
// without making it slower
// it is open for improvement

var borks = [16]string{"bork", "bark", "woof", "ruff", "rarf", "hooman", "doggo", "pleb", "bone", "woww", "much", "such", "very", "fren", "doge", "waht"}

var borks_in [256]string  // the bork, hooman, etc
var borks_out [256]string // the aa, ab, cd, fg, etc

var start_char byte = 'a'

func borkify(str *string) {

	//fmt.Print(*str + " became ")

	var endString = ""

	var strcpy string = *str

	for i := 0; i < len(strcpy); i++ {
		endString += borks_in[int8(strcpy[i])]
	}

	//fmt.Println(endString)

	*str = endString
}

// WARNING: this function is very inefficient and may burn your eyes
// please use the glasses we were supposed to wear during the eclipse
// and view from a safe distance
func unborkify(str *string) {

	//fmt.Print(*str + " became ")

	var i int = 0 // global current character
	var endString string = ""
	var strcpy string = *str

	for {

		////fmt.Println("i: " + string(i))
		//fmt.Printf("1. i: %d\n", i)
		firstWord := nextDogeWord(strcpy, i)
		i += len(firstWord)
		//fmt.Printf("2. i: %d\n", i)
		secondWord := nextDogeWord(strcpy, i)
		i += len(secondWord)
		//fmt.Printf("3. i: %d\n", i)

		currentWord := firstWord + secondWord
		//fmt.Println("currentWord: " + currentWord)
		////fmt.Println("\n1 i: " + string(i))

		////fmt.Println("currentword: " + currentWord)
		////fmt.Println("2 i: " + string(i))

		//convert it back to letters
		for j := 0; j < len(borks_in); j++ {
			if borks_in[j] == currentWord {
				endString += string(byte(j))
				//fmt.Println("added " + string(byte(j)) + " to " + endString)
				//fmt.Println(endString)
				break
			}
		}

		////fmt.Println("3 i: " + string(i))
		//i += len(currentWord) //make the incremental go up

		if i >= len(strcpy) {
			//fmt.Println("finished")
			//fmt.Println("unborked: " + endString)
			break
		}
	}

	//fmt.Println(endString)

	*str = endString
}

func borkify_init() {
	for i := 0; i < 16; i++ {
		for j := 0; j < 16; j++ {

			borks_in[i*16+j] = borks[i] + borks[j]

			// old algo
			/*//fmt.Printf("i=%d, j=%d\n", i, j)
			borks_in[j*16+i] = string(borks[j] + string(start_char+byte(i)))
			//fmt.Println("String created: " + borks_in[j*16+i])
			start_char++*/
		}
	}

	// don't print the output
	/*for i := 0; i < len(borks_in); i++ {
		//fmt.Println(borks_in[i])
	}*/
}

func nextDogeWord(str string, start_index int) string {

	current := ""

	//fmt.Printf("start index: %d\n", start_index)

	for i := start_index; i < len(str); i++ {
		current += string(str[i])
		if contains16(current, borks) {
			//fmt.Println("found dogeword: " + current)
			return current
		}
	}

	//fmt.Println("current: " + current)
	return "error"
}

func contains256(str string, strs [256]string) bool {
	for i := 0; i < len(strs); i++ {
		if strs[i] == str {
			return true
		}
	}

	return false
}

func contains16(str string, strs [16]string) bool {
	for i := 0; i < len(strs); i++ {
		if strs[i] == str {
			return true
		}
	}

	return false
}
