package hot100

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*func BuildTree(arr []int) *TreeNode {

	if len(arr) == 0 {
		return nil
	}

	root := &TreeNode{Val: arr[0]}
	if len(arr) == 1 {
		return root
	}

	mid := len(arr) / 2
	root.Val = arr[mid]

	root.Left = BuildTree(arr[:mid])
	root.Right = BuildTree(arr[mid+1:])

	return root
}*/
