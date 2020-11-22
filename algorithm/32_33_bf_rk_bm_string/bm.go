/**
 * Created by lock
 * bm算法实现
 */
package main

import (
	"fmt"
	"math"
)

type bmStrSearch struct {
	s          string
	pattern    string
	badCharMap []int
	suffix     []int
	prefix     []bool
}

func (bm *bmStrSearch) generateBadCharHashMap() {
	pattern := bm.pattern
	aBadChar := make([]int, 256)
	for i, _ := range aBadChar {
		aBadChar[i] = -1
	}
	for index, char := range pattern {
		aBadChar[int(char)] = index
	}
	bm.badCharMap = aBadChar
}

func (bm *bmStrSearch) generateGoodSuffix() {
	pattern := bm.pattern
	m := len(pattern)
	suffix := make([]int, m)
	prefix := make([]bool, m)
	for i := 0; i < m; i++ {
		suffix[i] = -1
		prefix[i] = false
	}
	for i := 0; i < m-1; i++ {
		j := i
		k := 0
		for j >= 0 && pattern[j] == pattern[m-1-k] {
			j--
			k++
			suffix[k] = j + 1
		}
		if j == -1 {
			prefix[k] = true
		}
	}
	bm.suffix = suffix
	bm.prefix = prefix
	return
}

func (bm *bmStrSearch) moveByGoodSuffix(badCharStartIndex int) int {
	patternLength := len(bm.pattern)
	//length of good suffix
	k := patternLength - badCharStartIndex - 1
	//complete match
	if bm.suffix[k] != -1 {
		return badCharStartIndex + 1 - bm.suffix[k]
	}
	//partial match
	for t := patternLength - 1; t > badCharStartIndex+1; t-- {
		if bm.prefix[t] {
			return t
		}
	}
	//no match
	return patternLength
}

func (bm *bmStrSearch) findBadChar(subStr string) (int, int) {
	j := -1
	badCharMapPos := -1
	badChar := rune(0)

	for index := len(subStr) - 1; index >= 0; index-- {
		if subStr[index] != bm.pattern[index] {
			j = index
			badChar = rune(subStr[index])
			break
		}
	}
	//if bad character exist, then find it's index at pattern
	if j > 0 {
		badCharMapPos = bm.badCharMap[int(badChar)]
	}
	return badCharMapPos, j
}

func newBmSearch(s, pattern string) *bmStrSearch {
	bm := &bmStrSearch{
		s:          s,
		pattern:    pattern,
		badCharMap: nil,
	}
	bm.generateBadCharHashMap()
	bm.generateGoodSuffix()
	return bm
}

//bad char && good suffix
func (bm *bmStrSearch) search() int {
	s, pattern := bm.s, bm.pattern
	sLen := len(s)
	patternLen := len(pattern)
	if sLen == 0 || patternLen == 0 || patternLen > sLen {
		return -1
	}
	step := 1
	for i := 0; i <= sLen-patternLen; i = i + step {
		subStr := s[i : i+patternLen]
		badCharMapPos, j := bm.findBadChar(subStr)
		stepForBadChar := j - badCharMapPos
		//j is bad char occur index, badCharMapPos is last bad char index in pattern
		if j == -1 {
			return i
		}
		stepForGoodSuffix := -1
		if j < patternLen-1 {
			stepForGoodSuffix = bm.moveByGoodSuffix(j)
		}
		//select max step
		step = int(math.Max(float64(stepForBadChar), float64(stepForGoodSuffix)))
	}
	return -1
}

func main() {
	s := "aaabcdefaacbcd"
	t := "aacb"
	bm := newBmSearch(s, t)
	index := bm.search()
	fmt.Println("bm str search result:", index)
}
