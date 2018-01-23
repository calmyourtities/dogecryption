// borkify
package main

import (
	"fmt"
	"strings"
)

// this algorithim is open for improvement
// it's supposed to translate a string of random characters
// to a sequence of strings like "bork" and "ruff"
// without making it slower
// it is open for improvement

var borks = [16]string{"bork", "bark", "woof", "ruff", "rarf", "hooman", "doggo", "pleb", "bone", "woww", "much", "such", "very", "fren", "doge", "waht"}

var borks_in [256]string
var borks_out [256]string

var a byte = 'a'

func borkify(str string) {
	strings.Replace(str, "", "", 0)
}

func borkify_init() {
	for i := 0; i < 16; i++ {
		for j := 0; j < 16; j++ {
			borks_in[j*16+i] = string(borks[i][:len(borks[i])-(i%4)] + string(int8(i/2)) + borks[i][len(borks[i]):])
		}
	}

	for i := 0; i < len(borks_in); i++ {
		fmt.Print(string(i))
		fmt.Println(borks_in[i])
	}
}
