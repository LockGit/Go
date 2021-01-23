/**
 * Created by lock
 * Date: 2021/1/15
动态规划之~~~房间取钱
一排房间，相邻的房间不能取钱，怎么取获得的金额最大
nums = [1,2,3,1]
输出：4，取1号房间和3号房间的钱金额最大

nums = [2,7,9,3,1]
输出：12 ，  2+9+1=12
*/
package main

import "fmt"

func maxMoneyCalc(a, b int) (v int) {
	if a > b {
		return a
	}
	return b
}

var memoMoney = make(map[string]int)

func roomMoney(arr []int, start int) (maxMoney int) {
	if start >= len(arr) {
		return 0
	}
	key := fmt.Sprintf("start_%d", start)
	if val, ok := memoMoney[key]; ok {
		return val
	}
	maxMoney = maxMoneyCalc(
		//不抢，去下家
		roomMoney(arr, start+1),
		//抢，去下下家
		arr[start]+roomMoney(arr, start+2),
	)
	memoMoney[key] = maxMoney
	return maxMoney
}

//动态规划，自底向上解法
func dpRoomMoney(arr []int) (maxMoney int) {
	n := len(arr)
	//dp[i] = x 表示：从第 i 间房子开始取钱，最多能抢到的钱为 x
	dp := make([]int, n+2)
	for i := n - 1; i >= 0; i-- {
		dp[i] = maxMoneyCalc(
			dp[i+1],        //不抢，去下家
			dp[i+2]+arr[i], //抢，去下下家
		)
	}
	return dp[0]
}

//如果房间围成了一个圈
//分三种情况，1，要么第一间被抢，2，要么最后一间被抢，3，要么两个都没抢。 1，2涵盖了3
func dpRingRoomMoney(arr []int) (maxVal int) {
	n := len(arr)
	if n == 1 {
		return arr[0]
	}
	return maxMoneyCalc(
		dpRoomMoneyIndex(arr, 0, n-1),
		dpRoomMoneyIndex(arr, 1, n-2),
	)
}

func dpRoomMoneyIndex(arr []int, start, end int) (maxMoney int) {
	n := len(arr)
	//dp[i] = x 表示：从第 i 间房子开始取钱，最多能抢到的钱为 x
	dp := make([]int, n+2)
	for i := end; i >= start; i-- {
		dp[i] = maxMoneyCalc(
			dp[i+1],        //不抢，去下家
			dp[i+2]+arr[i], //抢，去下下家
		)
	}
	return dp[0]
}

type tree struct {
	val   int
	left  *tree
	right *tree
}

//转变成二叉树求解，相邻节点不同同事被取钱
//可以增加备忘录
func treeRoomMoney(root *tree) (maxV int) {
	if root == nil {
		return 0
	}
	//抢，然后去下下家
	doIt := root.val
	if root.left != nil {
		doIt = doIt + treeRoomMoney(root.left.left) + treeRoomMoney(root.left.right)
	}
	if root.right != nil {
		doIt = doIt + treeRoomMoney(root.right.left) + treeRoomMoney(root.right.right)
	}
	// 不抢，然后去下家
	notDo := treeRoomMoney(root.left) + treeRoomMoney(root.right)
	return maxMoneyCalc(doIt, notDo)
}

func treeRoomMoneyFun2(root *tree) (res []int) {
	if root == nil {
		return []int{0, 0}
	}
	left := treeRoomMoneyFun2(root.left)
	right := treeRoomMoneyFun2(root.right)
	//抢，下家不能抢了
	doIt := root.val + left[0] + right[0]
	//不抢，下家可抢可不抢，取决于收益大小
	notDoIt := maxCalc(left[0], left[1]) + maxCalc(right[0], right[1])
	return []int{doIt, notDoIt}
}

func main() {
	arr := []int{2, 7, 9, 3, 1}
	fmt.Println(arr, ",max money is:", roomMoney(arr, 0))
	fmt.Println(arr, ",dp max money is:", dpRoomMoney(arr))
	fmt.Println(arr, ",dp ring max money is:", dpRingRoomMoney(arr))
}
