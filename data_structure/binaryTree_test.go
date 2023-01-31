package data_structure

import (
	"reflect"
	"testing"
)

func Test_preorderTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "空边界",
			args: args{root: nil},
			want: []int{},
		},
		{
			name: "根结点",
			args: args{root: &TreeNode{
				Val: 1,
			}},
			want: []int{1},
		},
		{
			name: "单一左结点",
			args: args{root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
			}},
			want: []int{1,2},
		},
		{
			name: "单一右结点",
			args: args{root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 3,
				},
			}},
			want: []int{1,3},
		},
		{
			name: "正常情况",
			args: args{root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val: 8,
						},
					},
				},
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 7,
					},
				},
			}},
			want: []int{1,2,3,4,7,8},
		},


	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := preorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("preorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_postorderTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "后序遍历",
			args: args{root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{Val: 3},
				},
			}},
			want: []int{3,2,1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := postorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}