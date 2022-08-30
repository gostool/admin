package tree

type Node[K any] struct {
	Id       int
	Children []*Node[K]
}

func Bfs[K any](treeMap map[int][]*Node[K]) (items []*Node[K]) {
	q := treeMap[0]
	items = q
	for len(q) > 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			node := q[i]
			if node == nil {
				continue
			}
			childList, ok := treeMap[node.Id]
			if !ok {
				// 当前菜单没有child
				continue
			}
			if node.Children == nil {
				n := len(childList)
				node.Children = make([]*Node[K], n, n)
			}
			copy(node.Children, childList)
			q = append(q, node.Children...)
		}
		q = q[size:]
	}
	return items
}

//func (n *Node[K]) GetId() int {
//	return n.Id
//}

//func (n *Node[K]) NewChild(l int) {
//	n.Children = make([]*Node[K], l, l)
//	return
//}

// New returns an empty B-tree.
