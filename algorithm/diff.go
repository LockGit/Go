/**
 * Created by Gogland.
 * User: lock
 * Date: 2017/9/21
 * Time: 15:56
 * git diff 算法
 */
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type operation uint

const INSERT operation = 1
const DELETE operation = 2
const MOVE operation = 3

var color = map[operation]string{
	INSERT: "\033[32m",
	DELETE: "\033[31m",
	MOVE:   "\033[39m",
}

// 支持使用负数作为索引
type intArrays []int

func (a intArrays) get(i int) int {
	if i < 0 {
		return a[len(a)+i]
	}
	return a[i]
}

func (a intArrays) set(i, v int) {
	if i < 0 {
		a[len(a)+i] = v
	} else {
		a[i] = v
	}
}

func prepend(s []operation, op operation) []operation {
	return append([]operation{op}, s...)
}

func (a intArrays) clone() intArrays {
	return append(intArrays{}, a...)
}

func (op operation) String() string {
	switch op {
	case INSERT:
		return "INS"
	case DELETE:
		return "DEL"
	case MOVE:
		return "MOV"
	default:
		return "UNKNOWN"
	}
}

func (a intArrays) String() string {
	k := len(a) / 2
	var result []string
	for i := -k; i <= k; i++ {
		result = append(result, strconv.Itoa(a.get(i)))
	}
	return strings.Join(result, " ")
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("usage: go run diff.go [src] [dst]")
		os.Exit(0)
	}
	src, err := readFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	dst, err := readFile(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	showDiff(src, dst)
}

/**
读取文件
*/
func readFile(p string) ([]string, error) {
	f, err := os.Open(p)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

func showDiff(src, dst []string) {
	script := findPath(src, dst)
	srcIndex, dstIndex := 0, 0
	for _, op := range script {
		switch op {
		case INSERT:
			fmt.Println(color[op] + "+ " + dst[dstIndex])
			dstIndex += 1
		case MOVE:
			fmt.Println(color[op] + "  " + src[srcIndex])
			srcIndex += 1
			dstIndex += 1
		case DELETE:
			fmt.Println(color[op] + "- " + src[srcIndex])
			srcIndex += 1
		}
	}
}

/**
寻找路径 核心算法
*/
func findPath(src, dst []string) []operation {
	n := len(src)
	m := len(dst)
	max := n + m
	v := make(intArrays, 2*max+1)
	var trace []intArrays
	var x, y int

loop:
	for d := 0; d <= max; d++ {
		trace = append(trace, v.clone())
		for k := -d; k <= d; k += 2 {
			if k == -d { //k=0,x=6
				x = v.get(k + 1)
			} else if k != d && v.get(k-1) < v.get(k+1) {
				x = v.get(k + 1)
			} else {
				x = v.get(k-1) + 1
			}

			y = x - k

			for x < n && y < m && src[x] == dst[y] {
				x, y = x+1, y+1
			}

			v.set(k, x)

			if x == n && y == m {
				break loop
			}
		}
	}

	//反向回溯路径
	var script []operation
	var k, prevK, prevX, prevY int

	for d := len(trace) - 1; d > 0; d-- {
		k = x - y
		v := trace[d]
		if k == -d || (k != d && v.get(k-1) < v.get(k+1)) {
			prevK = k + 1
		} else {
			prevK = k - 1
		}
		prevX = v.get(prevK)
		prevY = prevX - prevK

		for x > prevX && y > prevY {
			script = prepend(script, MOVE)
			x -= 1
			y -= 1
		}

		if x == prevX {
			script = prepend(script, INSERT)
		} else {
			script = prepend(script, DELETE)
		}

		x, y = prevX, prevY
	}
	return script
}
