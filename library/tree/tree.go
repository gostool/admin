package tree

type Node struct {
	Id       int
	Pid      int
	Children []*Node
}

//
func bfs(treeMap map[int][]*Node) (items []*Node) {
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
				node.Children = make([]*Node, n, n)
			}
			copy(node.Children, childList)
			q = append(q, node.Children...)
		}
		q = q[size:]
	}
	return
}
