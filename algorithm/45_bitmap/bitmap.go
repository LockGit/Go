/**
 * Created by lock
bitmap demo
将数字 A 的第 k 位设置为1：A = A | (1 << (k - 1))
将数字 A 的第 k 位设置为0：A = A & ~(1 << (k - 1))
检测数字 A 的第 k 位：A & (1 << (k - 1)) != 0
*/
package main

import "fmt"

type bitmap []byte

// New create BitMap
func newBitMap(length uint) bitmap {
	return make([]byte, length/8+1)
}

// Set set value in bitmap
func (b *bitmap) Set(value uint) {
	byteIndex := value / 8
	if byteIndex >= uint(len(*b)) {
		return
	}
	bitIndex := value % 8
	[]byte(*b)[byteIndex] |= 1 << bitIndex
}

// Get check whether value exist or not
func (b *bitmap) Get(value uint) bool {
	byteIndex := value / 8
	if byteIndex >= uint(len(*b)) {
		return false
	}
	bitIndex := value % 8
	return []byte(*b)[byteIndex]&(1<<bitIndex) != 0
}

func main() {
	bm := newBitMap(100)
	fmt.Println("before set,bitmap is:", bm)
	bm.Set(18)
	fmt.Println("after set,bitmap is:", bm)
	fmt.Println("bitmap get 18:", bm.Get(18))
	fmt.Println("bitmap check 20:", bm.Get(20))
	fmt.Println("bitmap check 500:", bm.Get(500))
}
