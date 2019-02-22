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
	leverNodes := []*TreeNode{root}
	valueParentNodeMap := make(map[int]*TreeNode, 0)

	for {
		if len(leverNodes) <= 0 {
			return false
		}
		valueParentNodeMap, leverNodes = getLevelNodeMap(leverNodes)
		parentX, findX := valueParentNodeMap[x]
		parentY, findY := valueParentNodeMap[y]
		if findX && findY && parentX != parentY {
			return true
		} else if !findX && !findY {
			continue
		} else {
			return false
		}
	}
}

func getLevelNodeMap(nodeSlice []*TreeNode) (map[int]*TreeNode, []*TreeNode) {
	levelNodes := make([]*TreeNode, 0)
	valueParentNodeMap := make(map[int]*TreeNode, 0)
	for _, v := range nodeSlice {
		if v.Left != nil {
			valueParentNodeMap[v.Left.Val] = v
			levelNodes = append(levelNodes, v.Left)
		}
		if v.Right != nil {
			valueParentNodeMap[v.Right.Val] = v
			levelNodes = append(levelNodes, v.Right)

		}
	}
	return valueParentNodeMap, levelNodes
}

//func isCousins(root *TreeNode, x int, y int) bool {
//	depthX, parentX := lookup(root, x)
//	depthY, parentY := lookup(root, y)
//	return depthX == depthY && parentX != parentY
//}
//
//func lookup(root *TreeNode, n int) (depth int, parent int) {
//    if root.Val == n {
//        return 0, 0
//    }
//    if root.Left != nil {
//        if root.Left.Val == n {
//            return 1, root.Val
//        }
//        dep, par := lookup(root.Left, n)
//        if dep != 0 {
//            return dep+1, par
//        }
//    }
//    if root.Right != nil {
//        if root.Right.Val == n {
//            return 1, root.Val
//        }
//        dep, par := lookup(root.Right, n)
//        if dep != 0 {
//            return dep+1, par
//        }
//    }
//    return 0, 0
//
//}
