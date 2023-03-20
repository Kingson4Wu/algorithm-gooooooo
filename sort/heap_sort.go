package sort

/**
堆排序（Heapsort）是指利用堆这种数据结构所设计的一种排序算法。堆积是一个近似完全二叉树的结构，并同时满足堆积的性质：即子结点的键值或索引总是小于（或者大于）它的父节点。

将初始待排序关键字序列(R1,R2….Rn)构建成大顶堆，此堆为初始的无序区；
将堆顶元素R[1]与最后一个元素R[n]交换，此时得到新的无序区(R1,R2,……Rn-1)和新的有序区(Rn),且满足R[1,2…n-1]<=R[n]；
由于交换后新的堆顶R[1]可能违反堆的性质，因此需要对当前无序区(R1,R2,……Rn-1)调整为新堆，然后再次将R[1]与无序区最后一个元素交换，得到新的无序区(R1,R2….Rn-2)和新的有序区(Rn-1,Rn)。不断重复此过程直到有序区的元素个数为n-1，则整个排序过程完成。
————————————————
版权声明：本文为CSDN博主「码墨」的原创文章，遵循CC 4.0 BY-SA版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/guidao13/article/details/86430483

*/

func initHead(nums []int, parent, len int) {
	temp := nums[parent]
	child := 2*parent + 1

	for child < len {
		//右子结点存在，且左子结点<右子结点
		//找出左右结点，那个最大，后续根结点只需跟这个子结点比较和交换即可
		if child+1 < len && nums[child] < nums[child+1] {
			child++
		}

		//子结点小于根结点
		if child < len && nums[child] <= temp {
			break
		}

		nums[parent] = nums[child]
		//这里子结点没重新赋值，在跳出循环后，统一处理，效果最终一样，
		//这里始终是和temp比较， temp即被交换过的值，即当前的根结点

		//交换之后，被交换的子结点，要作为父结点，重新调整对应的子结点
		parent = child
		child = parent*2 + 1

	}

	// 若进上面的循环： parent = child
	nums[parent] = temp
}

func HeadSort(nums []int) {
	/** 遍历所有的根结点，初始化堆，从 index=length/2 开始， 即从最底下的根结点到最上的根结点 */
	for i := len(nums) / 2; i >= 0; i-- {
		initHead(nums, i, len(nums))
	}

	/** 根结点和最后一个值交换，即最大值放最后一位 */
	for i := len(nums) - 1; i > 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		/** 重新调整堆，只需要调整最上的根结点的堆，因为其他根结点在初始化时已经调整过了 */
		initHead(nums, 0, i)
	}
}

//原文链接：https://blog.csdn.net/guidao13/article/details/86430483
