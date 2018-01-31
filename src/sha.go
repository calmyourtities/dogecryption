// sha.go
package dogecrypt

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

func Bamboozle(str string) string {
	temp := sha1.Sum([]byte(str)) // for some reason, golang has a
	// compile time error because of
	// how its dumb memory works,
	// see https://www.reddit.com/r/golang/comments/3oa9au/arrayslicing_rules/
	temp2 := hex.EncodeToString(temp[:])
	borkify(&temp2)
	return temp2
}

func Bamboozle224(str string) string {
	temp := sha256.Sum224([]byte(str))
	temp2 := hex.EncodeToString(temp[:])
	borkify(&temp2)
	return temp2
}

func Bamboozle256(str string) string {
	temp := sha256.Sum256([]byte(str))
	temp2 := hex.EncodeToString(temp[:])
	borkify(&temp2)
	return temp2
}

func Bamboozle384(str string) string {
	temp := sha512.Sum384([]byte(str))
	temp2 := hex.EncodeToString(temp[:])
	borkify(&temp2)
	return temp2
}

func Bamboozle512(str string) string {
	temp := sha512.Sum512([]byte(str))
	temp2 := hex.EncodeToString(temp[:])
	borkify(&temp2)
	return temp2
}
