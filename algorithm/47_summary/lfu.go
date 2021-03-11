/**
 * Created by lock
lfu算法
在一段时间内,数据被使用频次最少的,优先被淘汰
如果多个数据拥有相同的访问频次，我们就得删除最早插入的那个数据。
LFU 算法是淘汰访问频次最低的数据，如果访问频次最低的数据有多条，需要淘汰最旧的数据。
*/
package main

import (
	"fmt"
	"github.com/pkg/errors"
	"sync"
)

//链表
type linkLfu struct {
	Val  interface{}
	next *linkLfu
}

func (l *linkLfu) add(item *linkLfu) {
	l.next = item
}

func (l *linkLfu) remove(key interface{}) {

}

type LFUCache struct {
	//key value
	Kv map[interface{}]interface{}
	//key freq(使用频次)计数
	Kf map[interface{}]int
	//freq keys 列表映射
	freqToKeys map[int]*linkLfu //hash+link
	MinFreq    int
	Cap        int //容量
	Size       int //当前大小
	mu         *sync.Mutex
}

func NewLFUCache(capacity int) *LFUCache {
	return &LFUCache{
		Kv:         make(map[interface{}]interface{}),
		Kf:         make(map[interface{}]int),
		freqToKeys: make(map[int]*linkLfu),
		MinFreq:    0,
		Cap:        capacity,
		mu:         new(sync.Mutex),
	}
}

func (lfu *LFUCache) get(k interface{}) interface{} {
	lfu.mu.Lock()
	defer lfu.mu.Unlock()
	v, ok := lfu.Kv[k]
	if !ok {
		return nil
	}
	//key的freq次数增加
	//lfu.Kf[k] = lfu.Kf[k] + 1
	lfu.increaseFreq(k)
	return v
}

func (lfu *LFUCache) put(k, v interface{}) (err error) {
	lfu.mu.Lock()
	defer lfu.mu.Unlock()
	if lfu.Cap <= 0 {
		return errors.New("no cap")
	}
	//key 是否存在
	_, ok := lfu.Kv[k]
	if ok {
		lfu.Kv[k] = v
		//lfu.Kf[k]++
		lfu.increaseFreq(k)
		return
	}
	//key 不存在，需要插入
	if lfu.Size >= lfu.Cap {
		//淘汰使用次数少的key
		lfu.removeMinFreqKey()
	}
	lfu.Kv[k] = v
	//lfu.Kf[k] = 1
	lfu.increaseFreq(k)

	linkItem := &linkLfu{
		Val:  k,
		next: nil,
	}
	link, have := lfu.freqToKeys[1]
	if !have {
		lfu.freqToKeys[1] = linkItem
	} else {
		link.add(linkItem)
	}
	lfu.Size++
	lfu.MinFreq = 1
	return
}

func (lfu *LFUCache) increaseFreq(key interface{}) {
	freq, ok := lfu.Kf[key]
	if ok {
		lfu.Kf[key] = freq + 1
		//将 key 从 freq 对应的列表中删除
		lfu.freqToKeys[freq].remove(key)

		linkItem := &linkLfu{
			Val:  key,
			next: nil,
		}

		link, have := lfu.freqToKeys[freq+1]
		if !have {
			lfu.freqToKeys[1] = linkItem
		} else {
			link.add(linkItem)
		}

		if lfu.freqToKeys[freq] == nil {
			delete(lfu.freqToKeys, freq)
			if freq == lfu.MinFreq {
				lfu.MinFreq++
			}
		}
	}
}

//淘汰使用次数少的key,最最老的key删除
func (lfu *LFUCache) removeMinFreqKey() {
	oldMinLinkList := lfu.freqToKeys[lfu.MinFreq]
	//其中最先被插入的那个 key 就是该被淘汰的 key
	lfu.freqToKeys[lfu.MinFreq] = oldMinLinkList.next

	deleteKey := oldMinLinkList.Val
	//remove from Kv
	delete(lfu.Kv, deleteKey)
	//remove from Kf
	delete(lfu.Kf, deleteKey)
	lfu.Size--
}

func main() {
	lfu := NewLFUCache(4)
	lfu.put(1, "one")
	lfu.put(2, "two")
	lfu.put(3, "three")
	lfu.put(4, "four")
	lfu.put(5, "five")
	fmt.Println("get,", lfu.get(1))
	fmt.Println("get,", lfu.get(2))
	fmt.Println("get,", lfu.get(3))
	fmt.Println("get,", lfu.get(4))
	fmt.Println("get,", lfu.get(5))
	fmt.Println("get count,", lfu.Kf[1])
}
