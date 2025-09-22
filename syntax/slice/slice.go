package slice

import "fmt"

func NewSclice[T any](t T) []T {
	var slice []T
	return slice
}

// DeleteAt 删除指定索引的元素，支持泛型和自动缩容
// 要求一：基本删除功能
// 要求二：高性能实现（使用copy避免逐个移动）
// 要求三：泛型方法
// 要求四：支持缩容机制
func DeleteAt[T any](slice []T, index int) ([]T, error) {
	if index < 0 || index >= len(slice) {
		return slice, fmt.Errorf("index %d out of range [0,%d)", index, len(slice))
	}

	// 高性能删除：使用copy批量移动元素
	copy(slice[index:], slice[index+1:])

	// 缩短长度
	slice = slice[:len(slice)-1]

	// 缩容机制：当长度小于容量的1/4时，重新分配内存
	if shouldShrink(len(slice), cap(slice)) {
		return shrinkSlice(slice), nil
	}

	return slice, nil
}

// DeleteAtOptimized 高性能删除（不保证元素顺序）
// 将最后一个元素移到删除位置，避免大量数据移动
func DeleteAtOptimized[T any](slice []T, index int) ([]T, error) {
	if index < 0 || index >= len(slice) {
		return slice, fmt.Errorf("index %d out of range [0,%d)", index, len(slice))
	}

	// 如果不是最后一个元素，用最后一个元素替换
	lastIndex := len(slice) - 1
	if index != lastIndex {
		slice[index] = slice[lastIndex]
	}

	// 缩短长度
	slice = slice[:lastIndex]

	// 缩容机制
	if shouldShrink(len(slice), cap(slice)) {
		return shrinkSlice(slice), nil
	}

	return slice, nil
}

// DeleteRange 删除指定范围的元素 [start, end)
func DeleteRange[T any](slice []T, start, end int) ([]T, error) {
	if start < 0 || end > len(slice) || start > end {
		return slice, fmt.Errorf("invalid range [%d,%d) for slice length %d", start, end, len(slice))
	}

	if start == end {
		return slice, nil // 没有元素需要删除
	}

	// 高性能批量删除
	copy(slice[start:], slice[end:])
	slice = slice[:len(slice)-(end-start)]

	// 缩容机制
	if shouldShrink(len(slice), cap(slice)) {
		return shrinkSlice(slice), nil
	}

	return slice, nil
}

// shouldShrink 判断是否需要缩容
// 缩容策略：当长度小于容量的1/4且容量大于4时进行缩容
func shouldShrink(length, capacity int) bool {
	return capacity > 4 && length < capacity/4
}

// shrinkSlice 执行缩容操作
// 将容量缩减到长度的2倍，保证一定的增长空间
func shrinkSlice[T any](slice []T) []T {
	newCap := len(slice) * 2
	if newCap < 4 {
		newCap = 4 // 最小容量
	}

	newSlice := make([]T, len(slice), newCap)
	copy(newSlice, slice)
	return newSlice
}

// SliceInfo 获取切片的详细信息，用于调试
func SliceInfo[T any](slice []T) string {
	return fmt.Sprintf("len=%d, cap=%d", len(slice), cap(slice))
}
