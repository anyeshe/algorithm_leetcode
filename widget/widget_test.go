package widget

import "testing"

type Node struct {
	Name string
	Current int
	Widget int
}

// 加权轮询
// 请求来的时间不确定，但是要保证分给客户端的请求数是按照权重分摊的
func widget(nodes []*Node) (bestNode *Node) {

	if len(nodes) == 0 {
		return
	}

	total := 0

	for _, node := range nodes {
		total += node.Widget
		node.Current += node.Widget
		if bestNode == nil || node.Current > bestNode.Current {
			bestNode = node
		}
	}

	if bestNode == nil {
		return
	}

	bestNode.Current -= total

	return
}

func TestWidget(t *testing.T) {
	
	nodes := []*Node{
		&Node{
			Name:    "A",
			Current: 0,
			Widget:  4,
		},
		&Node{
			Name:    "B",
			Current: 0,
			Widget:  2,
		},
		&Node{
			Name:    "C",
			Current: 0,
			Widget:  1,
		},
	}
	
	
	for i := 0; i < 7; i++ {
		node := widget(nodes)
		t.Log(node.Name)
	}
}