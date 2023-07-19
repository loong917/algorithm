package algorithm

import (
	"fmt"
	"log"
)

// 原地排序算法，就是特指空间复杂度是 O(1) 的排序算法。
// 稳定排序算法，如果待排序的序列中存在值相等的元素，经过排序之后，相等元素之间原有的先后顺序不变。

// 1、冒泡排序
// 冒泡的过程只涉及相邻数据的交换操作，只需要常量级的临时空间，它的空间复杂度为 O(1)，所以冒泡排序是一个原地排序算法。
// 冒泡的过程当相邻的两个元素大小相等的时候，我们不做交换，相同大小的数据在排序前后不会改变顺序，所以冒泡排序是稳定的排序算法。
func BubbleSort(array []int, n int) []int {
	for i := 0; i < n; i++ {
		// 提前退出冒泡循环的标志位
		var flag bool = false
		for j := 0; j < n-i-1; j++ {
			if array[j] > array[j+1] {
				// 相邻数据交换
				array[j], array[j+1] = array[j+1], array[j]
				// 数据交换
				flag = true
			}
		}
		fmt.Printf("第 %d 次冒泡，冒泡后的结果：%v\n", i+1, array)
		if !flag {
			break
		}
	}
	return array
}

// 2、插入排序
// 在插入排序中，只需要常量级的临时空间，所以空间复杂度是 O(1)，所以插入排序是一个原地排序算法。
// 在插入排序中，对于值相同的元素，可以选择将后面出现的元素插入到前面出现元素的后面，这样就可以保持原有的前后顺序不变，所以插入排序是稳定的排序算法。
func InsertionSort(array []int, n int) []int {
	for i := 1; i < n; i++ {
		value := array[i]
		j := i - 1
		for j >= 0 && array[j] > value {
			// 数据移动
			array[j+1] = array[j]
			fmt.Printf("第 %d 次插入前移动，排序后的结果：%v\n", i, array)
			j--
		}
		// 插入数据
		array[j+1] = value
		fmt.Printf("第 %d 次插入，排序后的结果：%v\n", i, array)
	}
	return array
}

// 3、选择排序
// 在选择排序中，只需要常量级的临时空间，所以空间复杂度是 O(1)，所以选择排序是一个原地排序算法。
// 选择排序每次都要找剩余未排序元素中的最小值，并和前面的元素交换位置，这样破坏了稳定性，所以选择排序不是稳定的排序算法。
func SelectionSort(array []int, n int) []int {
	for i := 0; i < n; i++ {
		// 查找最小值
		min := i
		for j := i + 1; j < n; j++ {
			if array[j] < array[min] {
				min = j
			}
		}
		// 交换数据
		array[min], array[i] = array[i], array[min]
		fmt.Printf("第 %d 次交换，排序后的结果：%v\n", i+1, array)
	}
	return array
}

// 4、归并排序
// 在归并排序中，尽管每次合并操作都需要申请额外的内存空间，但在合并完成之后，临时开辟的内存空间就被释放掉了。临时内存空间最大也不会超过 n 个数据的大小，所以空间复杂度是 O(n)。归并排序不是一个原地排序算法。
// 如果 array[start...mid] 和 array[mid...end] 之间有值相同的元素，可以先把 array[start...mid] 中的元素放入 temp 数组。这样就保证了值相同的元素，在合并前后的先后顺序不变。所以归并排序是一个稳定的排序算法。
func MergeSort(array []int) []int {
	// 递归每一次函数调用都会有一次返回。当程序执行到某一级递归的结尾处时，它会转移到前一级递归继续执行。
	// 递归函数中，位于递归调用语句后的语句的执行顺序和各个被调用函数的顺序相反。
	if len(array) < 2 {
		return array
	}
	var middle = len(array) / 2
	a := MergeSort(array[:middle])
	log.Printf("mergeSort full = %v，left = %v", array, a)
	b := MergeSort(array[middle:])
	log.Printf("mergeSort full = %v，right = %v", array, b)
	log.Printf("mergeArray: a = %v，b = %v", a, b)
	c := mergeArray(a, b)
	return c
}

// 定义两个索引 i 和 j，分别指向 array[start...mid] 和 array[mid...end] 的第一个元素。
// 比较两个元素 array[i] 和 array[j]，如果 array[i] <= array[j]，array[i]放入到临时数组temp，并且 i 后移一位，
// 否则将 array[j] 放入到数组 temp，将j 后移一位。
// 直到其中一个子数组中的所有数据都放入临时数组中，再把另一个数组中的数据依次加入到临时数组的末尾。
func mergeArray(a []int, b []int) []int {
	var temp = make([]int, len(a)+len(b))
	var i = 0
	var j = 0
	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			temp[i+j] = a[i]
			i++
		} else {
			temp[i+j] = b[j]
			j++
		}
	}
	for i < len(a) {
		temp[i+j] = a[i]
		i++
	}
	for j < len(b) {
		temp[i+j] = b[j]
		j++
	}
	return temp
}

// 5、快速排序
// 快速排序通过设计巧妙的原地分区函数，不要额外申请内存空间，空间复杂度是 O(1)，所以快速排序是一个原地排序算法。
// 快速排序分区的过程涉及交换操作，如果数组中有两个相同的元素，在经过分区操作之后，两个相同的元素相对先后顺序就会改变。所以，快速排序并不是一个稳定的排序算法。
func QuickSort(array []int) {
	quickSortRange(array, 0, len(array)-1)
}

func quickSortRange(array []int, start int, end int) {
	if start >= end {
		return
	}
	partition := quickPartition(array, start, end)
	quickSortRange(array, start, partition-1)
	quickSortRange(array, partition+1, end)
}

func quickPartition(array []int, start int, end int) int {
	// 选取最后一位元素作为对比
	pivot := array[end]
	var pivotIndex int = start
	log.Printf("before array：%v，start：%d，end：%d，pivot：%d，index：%d", array, start, end, pivot, pivotIndex)
	for i := start; i < end; i++ {
		if array[i] < pivot {
			array[pivotIndex], array[i] = array[i], array[pivotIndex]
			pivotIndex++
		}
	}
	array[pivotIndex], array[end] = array[end], array[pivotIndex]
	log.Printf("after array：%v，start：%d，end：%d，pivot：%d，index：%d", array, start, end, pivot, pivotIndex)
	return pivotIndex
}
