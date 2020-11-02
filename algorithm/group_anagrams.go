/**
 * Created by GoLand.
 * User: lock
 * Date: 2018/8/28
 * Time: 11:20
 * 群组错位词分类
For example, given: ["eat", "tea", "tan", "ate", "nat", "bat"],
Return:

[
  ["ate", "eat","tea"],
  ["nat","tan"],
  ["bat"]
]

*/
package main

import (
	"fmt"
	"sort"
	"strings"
)

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func groupAnagrams(strArr []string) [][]string {
	dict := make(map[string]string)
	for _, item := range strArr {
		k := makeKeys(item)
		if dict[k] == "" {
			dict[k] = item
		} else {
			dict[k] = dict[k] + "," + item
		}
	}
	var res [][]string
	for _, v := range dict {
		arr := strings.Split(v, ",")
		var tmp []string
		for _, val := range arr {
			tmp = append(tmp, val)
		}
		res = append(res, tmp)
	}
	return res
}

func makeKeys(item string) string {
	return SortString(item)
}

func main() {
	strArr := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strArr))
}
