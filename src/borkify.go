// borkify
package main

import (
	"fmt"
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

	fmt.Print(*str + " became ")

	var endString = ""

	var strcpy = *str

	for i := 0; i < len(*str); i++ {
		endString += borks_in[int8(str[i])]
	}

	fmt.Println(endString)

	*str = endString
}

func unborkify(str string) {

}

func borkify_init() {
	for i := 0; i < 16; i++ {
		for j := 0; j < 16; j++ {
			fmt.Printf("i=%d, j=%d\n", i, j)
			borks_in[j*16+i] = string(borks[j] + string(start_char+byte(i)))
			fmt.Println("String created: " + borks_in[j*16+i])
			start_char++
		}
	}

	for i := 0; i < len(borks_in); i++ {
		fmt.Println(borks_in[i])
	}
}
