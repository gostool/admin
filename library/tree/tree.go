package tree

type INode interface {
	GetId() int
	ChildIsNil() bool
	NewChild(n int)
	SetChild(childList []INode)
	GetChildList() []INode
}

func Bfs(treeMap map[int][]INode) (items []INode) {
	q := treeMap[0]
	items = q
	for len(q) > 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			node := q[i]
			if node == nil {
				continue
			}
			childList, ok := treeMap[node.GetId()]
			if !ok {
				// 当前菜单没有child
				continue
			}
			if node.ChildIsNil() {
				n := len(childList)
				node.NewChild(n)
			}
			node.SetChild(childList)
			q = append(q, node.GetChildList()...)
		}
		q = q[size:]
	}
	return items
}

// Convenience types for common cases
type Node struct {
	Id    int
	Child []*Node
}

func (x *Node) GetId() int {
	return x.Id
}

func (x *Node) ChildIsNil() bool {
	return x.Child == nil
}
func (x *Node) NewChild(n int) {
	x.Child = make([]*Node, n, n)
	return
}

func (x *Node) SetChild(childList []*Node) {
	copy(x.Child, childList)
	return
}

func (x *Node) GetChildList() (childList []*Node) {
	return x.Child
}
