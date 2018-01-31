// Playbook - http://play.golang.org/p/3wFl4lacjX

package dogecrypt

import (
	_ "bufio"
	_ "fmt"
	_ "os"
)

/*func main() {

	dogecrypt_init()

	fmt.Print("Enter your key: ")
	reader := bufio.NewReader(os.Stdin)
	strKey, _ := reader.ReadString('\n')
	fmt.Println("What should be encrypted?")
	msg, _ := reader.ReadString('\n')
	msg = msg[:len(msg)-1]                         // \n should not be in message
	strKey = strKey[:len(strKey)-1]                // \n should not be in key
	dogecrypted_text := dogecrypt_enc(msg, strKey) // sample key: 7E892875A52C59A3B588306B13C31FBD
	fmt.Printf("Dogecrypted: %s\n", dogecrypted_text)
	undogecrypted_text := dogecrypt_dec(msg, strKey)
	fmt.Printf("Undogecrypted: %s\n", undogecrypted_text)
}*/

func Dogecrypt_init() {
	borkify_init()
}

func AesEnc(raw, key string) string {
	encrypted, _ := encrypt([]byte(key), raw)
	borkify(&encrypted)
	return encrypted
}

func AesDec(raw, key string) string {
	unborkify(&raw)
	decrypted, _ := decrypt([]byte(key), raw)
	return decrypted
}
