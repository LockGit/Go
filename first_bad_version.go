/**
 * Created by GoLand.
 * User: lock
 * Date: 2018/8/28
 * Time: 10:43
 * 二分查找
	你是一个产品经理当前正带领团队开发一个新的产品。不幸的是，你最新版本的产品没有通过质量审核。因为每一个版本的开发都基于前一个版本，一旦有一个版本是坏的，所有之后的版本都是坏的。
	假设你有 n 个版本 [1, 2, ..., n] 并且你想要找出第一个坏的从而导致后面的版本都坏了的那个。
	你被给定一个接口 bool isBadVersion(version) 能够返回 Version 是否是坏的。实现一个函数来找到第一个坏的版本。你应该最小化调用接口的次数。
*/
package main

import "fmt"

func isBadVer(index int) bool {
	//第二个版本是正确版本，那么第3个版本就是第一个坏了的版本
	if index == 2 {
		return false
	}
	return true
}

func main() {
	left := 1
	right := 6
	for left+1 < right {
		mid := left + (right-left)/2
		if isBadVer(mid) {
			right = mid
		} else {
			left = mid
		}
	}
	if isBadVer(left) {
		fmt.Println("first bad version is", left)
	} else if isBadVer(right) {
		fmt.Println("first bad version is", right)
	}
}
