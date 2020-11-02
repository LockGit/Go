/**
 * Created by lock
 * Date: 2020/10/26
 */
package hash_node

import (
	"hash/crc32"
	"sort"
	"strconv"
	"sync"
)

type HashFunc func(data []byte) uint32 //(0 to 4294967295)

type HashMap struct {
	HashFunc HashFunc
	Replicas int
	Keys     []int
	HashMap  map[int]string
	mu       sync.RWMutex
}

func New(replicas int, hashFunc HashFunc) (hm *HashMap) {
	if hashFunc == nil {
		hashFunc = crc32.ChecksumIEEE
	}
	hm = &HashMap{
		HashFunc: hashFunc,             //Hash 函数 (可以自定定义)
		Replicas: replicas,             //虚拟节点倍数
		Keys:     nil,                  //哈希环
		HashMap:  make(map[int]string), //虚拟节点与真实节点的映射表,值是真实节点的名称
	}
	return hm
}

func (hm *HashMap) Add(keys ...string) {
	for _, key := range keys {
		for i := 0; i < hm.Replicas; i++ {
			hashKey := strconv.Itoa(i) + key
			hashVal := int(hm.HashFunc([]byte(hashKey)))
			hm.Keys = append(hm.Keys, hashVal)
			hm.mu.Lock()
			hm.HashMap[hashVal] = key
			hm.mu.Unlock()
		}
	}
	sort.Ints(hm.Keys) //环上的哈希值排序
}

func (hm *HashMap) Get(key string) (realNode string) {
	nodeCount := len(hm.Keys)
	if nodeCount <= 0 {
		return
	}
	hashVal := int(hm.HashFunc([]byte(key)))
	idx := sort.Search(nodeCount, func(i int) bool {
		return hm.Keys[i] >= hashVal
	})
	hm.mu.RLock()
	defer hm.mu.RUnlock()
	return hm.HashMap[hm.Keys[idx%nodeCount]]
}
