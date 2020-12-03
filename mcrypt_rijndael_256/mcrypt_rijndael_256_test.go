/**
 * Created by lock
 */
package mcrypt_rijndael_256

import "testing"

func Test_TestDecrypt(t *testing.T) {
	srcData := []byte("your data")
	key := []byte("your key, 16/24/32 bit")
	//这Decrypt里面中blockSize=32
	Decrypt(srcData, key)
}
