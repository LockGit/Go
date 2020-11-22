/**
 * Created by lock
 * 基数排序：
	需要可以分割出独立的“位”来比较，
    而且位之间有递进的关系，如果 a 数据的高位比 b 数据大，那剩下的低位就不用比较了。
	除此之外，每一位的数据范围不能太大，要可以用线性排序算法来排序，否则，基数排序的时间复杂度就无法做到 O(n) 了
*/
package main

import (
	"fmt"
	"sort"
	"strconv"
)

type phoneDataList struct {
	PhoneNumStr string
	BitNum      int
}

func baseSort(arr []string) {
	for i := 10; i >= 0; i-- {
		var aPhone []phoneDataList
		for _, v := range arr {
			bitNumStr := v[i : i+1]
			bitNum, _ := strconv.Atoi(bitNumStr)
			item := phoneDataList{
				PhoneNumStr: v,
				BitNum:      bitNum,
			}
			aPhone = append(aPhone, item)
		}
		sort.Slice(aPhone, func(i, j int) bool {
			if aPhone[i].BitNum < aPhone[j].BitNum {
				return true
			}
			return false
		})
		var aPhoneNum []string
		for _, v := range aPhone {
			aPhoneNum = append(aPhoneNum, v.PhoneNumStr)
		}
		copy(arr, aPhoneNum)
	}
	return
}

func main() {
	//手机号排序
	arr := []string{"13800001234", "13800001235", "13800001034", "13800002289"}
	baseSort(arr)
	fmt.Println("base sort after:", arr)
}
