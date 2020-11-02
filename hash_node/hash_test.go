/**
 * Created by lock
 * Date: 2020/10/27
 */
package hash_node

import (
	"strconv"
	"testing"
)

func Test_TestNewHash(t *testing.T) {
	//hn := New(3, crc32.ChecksumIEEE)
	hn := New(3, func(key []byte) uint32 {
		i, _ := strconv.Atoi(string(key))
		return uint32(i)
	})
	// 2, 4, 6, 12, 14, 16, 22, 24, 26
	hn.Add("6", "4", "2")
	t.Logf("real node count:%d,all node cout:%d,node keys:%+v", 3, len(hn.Keys), hn.Keys)
	testCases := map[string]string{
		"2":  "2",
		"11": "2",
		"23": "4",
		"27": "2",
	}
	for k, v := range testCases {
		if hn.Get(k) != v {
			t.Errorf("request for %s, should return %s", k, v)
		}
	}

	// Adds 8, 18, 28
	hn.Add("8")
	testCases["27"] = "8"

	for k, v := range testCases {
		if hn.Get(k) != v {
			t.Errorf("request for %s, should return %s", k, v)
		}
	}

}
