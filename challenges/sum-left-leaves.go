package challenges

// TreeNode -
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	var sum int

	getSum(root, &sum)

	return sum
}

func isleftLeaf(root *TreeNode) bool {
	if root != nil && root.Left != nil && root.Left.Left == nil && root.Left.Right==nil {
		return true
	}
	return false
}

func getSum(root *TreeNode, sum *int) {
	if root == nil {
		return
	}
	if isleftLeaf(root) {
		*sum += root.Left.Val
	}
	getSum(root.Left, sum)
	getSum(root.Right, sum)

}
