/**
 * Created by lock
 */
package round_node

import (
	"testing"
)

func Test_TestRoundRobin(t *testing.T) {
	nodeList := []*node{
		{host: "127.0.0.1", weight: 5, current: 0},
		{host: "192.0.0.1", weight: 1, current: 0},
		{host: "223.0.0.1", weight: 1, current: 0},
	}
	rr := new(roundRobin)
	rr.nodeList = nodeList
	for i := 0; i < 10; i++ {
		selectNode := rr.selectNext()
		t.Log(i, "=>", selectNode.host, ",weight=", selectNode.weight)
	}
}
