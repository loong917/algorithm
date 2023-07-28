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

// 6、桶排序
// 桶排序对待排序的数据要求非常苛刻。
// 首先，要排序的数据需要很容易就能划分成 m 个桶，并且，桶与桶之间有着天然的大小顺序。这样每个桶内的数据都排序完之后，桶与桶之间的数据不需要再进行排序。
// 其次，数据在各个桶之间的分布是比较均匀的。
func BucketSort(array []int, bucketSize int) {
	// 获取最大元素值
	var length int = len(array)
	if length <= 1 {
		return
	}
	var min int = array[0]
	var max int = array[0]
	for i := 1; i < length; i++ {
		if array[i] > max {
			max = array[i]
		}
		if array[i] < min {
			min = array[i]
		}
	}
	// 二维切片
	buckets := make([][]int, length/bucketSize+1)
	for _, n := range array {
		idx := (n - min) / bucketSize
		buckets[idx] = append(buckets[idx], n)
	}
	// 索引
	position := 0
	for index, bucket := range buckets {
		bucketLength := len(bucket)
		if bucketLength > 0 {
			// 桶内排序：快排
			QuickSort(bucket)
			copy(array[position:], buckets[index])
			position += bucketLength
		}
	}
}

// 7、计数排序
// 计数排序只能用在数据范围不大的场景中（比如高考成绩），如果数据范围 k 比要排序的数据 n 大很多，就不适合用计数排序了。
// 计数排序只能给非负整数排序，如果要排序的数据是其他类型的，要将其在不改变相对大小的情况下，转化为非负整数。
func CountingSort(array []int) []int {
	// 获取最大元素值
	var length int = len(array)
	if length <= 1 {
		return array
	}
	var max int = array[0]
	for i := 1; i < length; i++ {
		if array[i] > max {
			max = array[i]
		}
	}
	// 申请一个计数数组，下标大小[0,max]
	countArray := make([]int, max+1)
	// 计算每个元素的个数
	for i := 0; i < length; i++ {
		countArray[array[i]] += 1
	}
	// 依次累加
	for i := 1; i <= max; i++ {
		countArray[i] += countArray[i-1]
	}
	log.Println("计数数组：", countArray)
	// 从后到前一次扫描数据
	resultArray := make([]int, length)
	for i := length - 1; i >= 0; i-- {
		index := countArray[array[i]]
		resultArray[index-1] = array[i]
		countArray[array[i]] -= 1
	}
	return resultArray
}

// 8、基数排序
// 基数排序对要排序的数据是有要求的，需要可以分割出独立的“位”来比较，而且位之间有递进的关系，如果 a 数据的高位比 b 数据大，那剩下的低位就不用比较了。
// 除此之外，每一位的数据范围不能太大，要可以用线性排序算法来排序，否则，基数排序的时间复杂度就无法做到 O(n) 了。
func RadixSort(array []int) {
	var length int = len(array)
	if length <= 1 {
		return
	}
	// 获取最大元素值
	var max int = array[0]
	for i := 1; i < length; i++ {
		if array[i] > max {
			max = array[i]
		}
	}
	for place := 1; max/place > 0; place *= 10 {
		radixPartition(array, place)
	}
}

// 对每个数位上的数字进行计数排序
func radixPartition(array []int, exp int) {
	var length int = len(array)
	if length <= 1 {
		return
	}
	// 申请一个计数数组
	countArray := make([]int, 10)
	// 计算每个元素的个数
	for _, val := range array {
		category := (val / exp) % 10
		countArray[category] += 1
	}
	// 依次累加
	for i := 1; i < 10; i++ {
		countArray[i] += countArray[i-1]
	}

	log.Printf(" %d 位计数数组：%v", exp, countArray)

	// 从后到前一次扫描数据
	resultArray := make([]int, length)
	for i := length - 1; i >= 0; i-- {
		divisor := (array[i] / exp) % 10
		index := countArray[divisor] - 1
		resultArray[index] = array[i]
		countArray[divisor] -= 1

		log.Printf("第 %d 轮排序：%v", i, resultArray)
	}

	log.Printf(" %d 位排序结果：%v", exp, resultArray)

	for i := 0; i < length; i++ {
		array[i] = resultArray[i]
	}
}
