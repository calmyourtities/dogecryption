// Playbook - http://play.golang.org/p/3wFl4lacjX

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	borkify_init()

	fmt.Print("Enter your key: ")
	reader := bufio.NewReader(os.Stdin)
	strKey, _ := reader.ReadString('\n')
	fmt.Println("What should be encrypted?")
	msg, _ := reader.ReadString('\n')
	msg = msg[:len(msg)-1]               // \n should not be in message
	strKey = strKey[:len(strKey)-1]      // \n should not be in key
	key := []byte(strKey)                // sample key: 7E892875A52C59A3B588306B13C31FBD
	encryptedMsg, _ := encrypt(key, msg) // sample text: Hello World
	fmt.Printf("Encrypted message: %s\n", encryptedMsg)
	msg, _ = decrypt(key, encryptedMsg)
	fmt.Print("Decrypted message: ")
	fmt.Println(msg) // Hello World
}
