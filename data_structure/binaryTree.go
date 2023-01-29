package data_structure

//### 二叉树遍历
//
//**前序遍历**：**先访问根节点**，再前序遍历左子树，再前序遍历右子树
//**中序遍历**：先中序遍历左子树，**再访问根节点**，再中序遍历右子树
//**后序遍历**：先后序遍历左子树，再后序遍历右子树，**再访问根节点**
//
//注意点
//
//- 以根访问顺序决定是什么遍历
//- 左子树都是优先右子树

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

//前序非递归 利用栈的思想存储
func preorderTraversal(root *TreeNode) []int  {
	res := make([]int,0)
	if root == nil {
		return res
	}

	stackList := make([]*TreeNode,0)
	stackList = append(stackList, root)
	for len(stackList) > 0  {
		currentNode := stackList[0]
		res = append(res, currentNode.Val)
		if currentNode.Left != nil {
			stackList = append(stackList, currentNode.Left)
		}

		if currentNode.Right != nil {
			stackList = append(stackList, currentNode.Right)
		}

		stackList = stackList[1:]
	}

	return res
}

//中序非递归 利用栈的思想存储
func inorderTraversal(root *TreeNode) []int  {
	res := make([]int,0)
	stackList := make([]*TreeNode,0)
	for root != nil || len(stackList) > 0  {
		//先获取左子树,
		for root != nil {
			stackList =append(stackList, root)
			root = root.Left
		}

		//拿到最左子树的值放到结果中
		root = stackList[len(stackList)-1]
		stackList = stackList[:len(stackList)-1]
		res = append(res, root.Val)
		//遍历右子树
		root = root.Right

	}

	return res
}

//后序非递归 利用栈的思想存储
func postorderTraversal(root *TreeNode) []int  {
	res := make([]int,0)
	if root == nil {
		return res
	}

	stackList := make([]*TreeNode,0)
	stackList = append(stackList, root)
	for len(stackList) > 0  {
		currentNode := stackList[0]
		res = append(res, currentNode.Val)
		if currentNode.Left != nil {
			stackList = append(stackList, currentNode.Left)
		}

		if currentNode.Right != nil {
			stackList = append(stackList, currentNode.Right)
		}

		stackList = stackList[1:]
	}

	return res
}

//前序递归
func preorderTraversalRecursion(root *TreeNode)  {

}


