package tree

/*

Copy

b = make([]T, len(a))
copy(b, a) // or b = append([]T(nil), a...)
Cut

a = append(a[:i], a[j:]...)
Delete

a = append(a[:i], a[i+1:]...) // or a = a[:i+copy(a[i:], a[i+1:])]
Delete without preserving order

a[i], a = a[len(a)-1], a[:len(a)-1]
Pop

x, a = a[len(a)-1], a[:len(a)-1]
Push

a = append(a, x)

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCousins(root *TreeNode, x int, y int) bool {

	if root.Val == x || root.Val == y {
		return false
	}
	nodeSlice := []*TreeNode{root}
	valueParentNodeMap := make(map[int]*TreeNode, 0)

	for {
		if len(nodeSlice) <= 0 {
			return false
		}
		valueParentNodeMap, nodeSlice = getLevelNodeMap(nodeSlice)
		valX, okx := valueParentNodeMap[x]
		valY, oky := valueParentNodeMap[y]
		if okx && oky && valX != valY {
			return true
		} else if !okx && !oky {
			continue
		} else {
			return false
		}
	}
}

func getLevelNodeMap(nodeSlice []*TreeNode) (map[int]*TreeNode, []*TreeNode) {
	nodeLevel := make([]*TreeNode, 0)
	valueParentNodeMap := make(map[int]*TreeNode, 0)
	for _, v := range nodeSlice {
		if v.Left != nil {
			valueParentNodeMap[v.Left.Val] = v
			nodeLevel = append(nodeLevel, v.Left)
		}
		if v.Right != nil {
			valueParentNodeMap[v.Right.Val] = v
			nodeLevel = append(nodeLevel, v.Right)

		}
	}
	return valueParentNodeMap, nodeLevel
}
