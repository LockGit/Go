/**
 * Created by lock
基于二分搜索的问题---最低吃香蕉速度
有一堆香蕉，[5,2,4,3,1,8,9,10,4,5]
吃的速度为: k根/h
限制为每小时最多吃一堆香蕉，如果吃不下的话留到下一小时再吃
如果吃完了这一堆还有胃口，也只会等到下一小时才会吃下一堆
现在有H小时，求吃的最小速度为多少（根/小时 才能保证在H小时内全部吃完所有香蕉
*/
package main

import "fmt"

func getMaxPiles(arr []int) (maxVal int) {
	maxVal = arr[0]
	for i := 1; i <= len(arr)-1; i++ {
		if arr[i] > maxVal {
			maxVal = arr[i]
		}
	}
	return maxVal
}

func canEatFinish(arr []int, speed, h int) (b bool) {
	useTime := 0
	for i := 0; i < len(arr); i++ {
		useTime = useTime + timeOf(arr[i], speed)
	}
	return useTime <= h
}

//吃某一堆香蕉，速度为speed，所需时间
func timeOf(num, speed int) (needTime int) {
	last := 0
	if num%speed > 0 {
		last = 1
	}
	needTime = (num / speed) + last
	return needTime
}

func getMinEatSpeed(arr []int, h int) (minSpeed int) {
	//speed in min[1~max(arr)]
	maxSpeed := getMaxPiles(arr)
	for speed := 1; speed <= maxSpeed; speed++ {
		if canEatFinish(arr, speed, h) {
			return speed
		}
	}
	return maxSpeed
}

//二分法最小吃香蕉速度
func getMinEatSpeedByBinary(arr []int, h int) (minSpeed int) {
	right := getMaxPiles(arr)
	left := 1
	//搜索左侧边界的二分查找来代替线性搜索
	for left < right {
		midSpeed := left + ((right - left) >> 1)
		if canEatFinish(arr, midSpeed, h) {
			right = midSpeed
		} else {
			left = midSpeed + 1
		}
	}
	return left
}

func main() {
	arr := []int{3, 6, 7, 11}
	h := 8
	fmt.Println(arr, ",h=", h, ",getMinEatSpeed min speed is:", getMinEatSpeed(arr, h))
	fmt.Println(arr, ",h=", h, ",getMinEatSpeedByBinary min speed is:", getMinEatSpeedByBinary(arr, h))
	fmt.Println([]int{10}, ",h=", h, ",getMinEatSpeedByBinary min speed is:", getMinEatSpeedByBinary([]int{10}, 1))
}
