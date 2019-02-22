package tree

import "testing"

func Test_isCousins(t *testing.T) {
	type args struct {
		root *TreeNode
		x    int
		y    int
	}
	node := TreeNode{4, nil, nil}
	node2 := TreeNode{5, nil, nil}
	left := TreeNode{2, &node, nil}
	right := TreeNode{3, nil, nil}
	right2 := TreeNode{3, nil, &node2}
	root := TreeNode{1, &left, &right}
	root2 := TreeNode{1, &left, &right2}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{&root, 3, 4}, false},
		{"2", args{&root2, 4, 5}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isCousins(tt.args.root, tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("isCousins() = %v, want %v", got, tt.want)
			}
		})
	}
}
