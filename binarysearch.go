package algorithm

// 二分查找
// 1）二分查找只能用在数据是通过顺序表（数组）来存储的数据结构上。
// 2）二分查找针对的是有序数据。
// 3）数据量太小不适合二分查找。如果要处理的数据量很小，顺序遍历就足够了。（如果数据之间的比较非常耗时，无论数据量大小，推荐使用二分查找。）
// 4）数据量太大不适合二分查找。二分查找是作用在数组这种数据结构上，数组需要连续的内存空间，太大的数据不适合使用数组来存储，也就不能用二分查找。
func BinarySearch(array []int, query int) int {
	var length int = len(array)
	if length == 0 {
		return -1
	}
	return binarySearchRecursive(array, query, 0, length-1)
}

func binarySearchRecursive(array []int, value int, low int, high int) int {
	if low > high {
		return -1
	}
	// 实际上，mid = (low + high) / 2 这种写法是有问题的。因为如果 low 和 high 比较大的话，两者之和就有可能会溢出。
	// 改进的方法是将 mid 的计算方式写成 low + (high-low)/2。
	// 这里可以将除以 2 操作转化成位运算 low + ((high - low) >> 1)。因为相比除法运算来说，计算机处理位运算要快得多。
	mid := low + ((high - low) >> 1)
	if array[mid] == value {
		return mid
	} else if array[mid] > value {
		return binarySearchRecursive(array, value, low, mid-1)
	} else {
		return binarySearchRecursive(array, value, mid+1, high)
	}
}
