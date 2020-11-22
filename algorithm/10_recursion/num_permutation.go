/**
 * Created by lock
 * 一组数据集合的全排列 + (递归+回溯)
 */
package main

import (
	"fmt"
)

type numPermutation struct {
	aData  []interface{}
	Length int
}

func newNumPermutation(aData []interface{}) *numPermutation {
	return &numPermutation{
		aData:  aData,
		Length: len(aData),
	}
}

func (n *numPermutation) GetAllRange(index int) {
	if index == n.Length-1 {
		fmt.Println(n.aData)
		return
	}
	for i := index; i < n.Length; i++ {
		if i == index || (n.aData[i] != n.aData[index]) {
			n.aData[i], n.aData[index] = n.aData[index], n.aData[i]
			//fmt.Println(fmt.Sprintf("before--%d--,value:%v", index, n.aData))
			n.GetAllRange(index + 1)
			//fmt.Println(fmt.Sprintf("after--%d--,value:%v", index, n.aData))
			//回溯
			n.aData[i], n.aData[index] = n.aData[index], n.aData[i]
		}
	}
}

func main() {
	aData := make([]interface{}, 0)
	aData = append(aData, "a")
	aData = append(aData, "b")
	aData = append(aData, "c")
	o := newNumPermutation(aData)
	o.GetAllRange(0)
}
