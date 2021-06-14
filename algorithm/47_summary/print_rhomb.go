/**
 * Created by lock
 * Date: 2021/3/12
	打印菱形

n=1
       *

n=3
       *       1

	 * * *     3

   * * * * *   5

	 * * *     3

	   *       1
*/
package main

import (
	"fmt"
	"strings"
)

func getShowStr(num int, show string) string {
	var showArr []string
	for i := 0; i < num; i++ {
		showArr = append(showArr, show)
	}
	return strings.Join(showArr, " ")
}

func printRhomb(n int, show string) {
	if n%2 != 1 {
		fmt.Println("菱形必须为奇数")
		return
	}
	//var arr []string
	a := 2 * n
	for i := 1; i < a; i++ {
		space := getShowStr((n+2-i)/2+1, " ")
		if i%2 == 0 {
			fmt.Println()
		} else {
			fmt.Println(fmt.Sprintf("%s%s%s", space, getShowStr(i, show), space))
		}
	}
	for i := a - 2; i >= 1; i-- {
		space := getShowStr((n+2-i)/2+1, " ")
		if i%2 == 0 {
			fmt.Println()
		} else {
			fmt.Println(fmt.Sprintf("%s%s%s", space, getShowStr(i, show), space))
		}
	}
}

//如果要在中间隔一行，line append 到数组 最后统一输出处理就可以
func printBestRhomb(n int, show string) {
	for i := 0; i <= n; i++ {
		line := fmt.Sprintf("%s%s", getShowStr(n-i+1, " "), getShowStr(i, show))
		if i > 0 {
			line += fmt.Sprintf(" %s", getShowStr(i-1, show))
		}
		fmt.Println(line)
	}
	for i := n - 1; i >= 1; i-- {
		line := fmt.Sprintf("%s%s", getShowStr(n-i+1, " "), getShowStr(i, show))
		if i > 1 {
			line += fmt.Sprintf(" %s", getShowStr(i-1, show))
		}
		fmt.Println(line)
	}
}

func main() {
	n := 7
	printRhomb(n, "*")
	fmt.Println("++++++++")
	printBestRhomb(5, "*")
}
