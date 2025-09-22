package slice

import (
	"reflect"
	"testing"
)

func TestDeleteAt(t *testing.T) {
	// 测试整数切片
	intSlice := []int{1, 2, 3, 4, 5}
	result, err := DeleteAt(intSlice, 2)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expected := []int{1, 2, 4, 5}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	// 测试字符串切片
	strSlice := []string{"a", "b", "c", "d"}
	result2, err := DeleteAt(strSlice, 0)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expected2 := []string{"b", "c", "d"}
	if !reflect.DeepEqual(result2, expected2) {
		t.Errorf("Expected %v, got %v", expected2, result2)
	}

	// 测试边界情况
	_, err = DeleteAt([]int{1}, 1)
	if err == nil {
		t.Error("Expected error for out of range index")
	}

	_, err = DeleteAt([]int{1}, -1)
	if err == nil {
		t.Error("Expected error for negative index")
	}
}

func TestDeleteAtOptimized(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	result, err := DeleteAtOptimized(slice, 1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// 检查长度
	if len(result) != 4 {
		t.Errorf("Expected length 4, got %d", len(result))
	}

	// 检查第1个位置现在是原来的最后一个元素
	if result[1] != 5 {
		t.Errorf("Expected result[1] to be 5, got %d", result[1])
	}
}

func TestDeleteRange(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6, 7}
	result, err := DeleteRange(slice, 2, 5)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expected := []int{1, 2, 6, 7}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	// 测试删除空范围
	result2, err := DeleteRange(slice, 2, 2)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !reflect.DeepEqual(result2, slice) {
		t.Errorf("Expected no change, but got %v", result2)
	}
}

func TestShrinkMechanism(t *testing.T) {
	// 创建一个大容量的切片
	slice := make([]int, 100, 100)
	for i := 0; i < 100; i++ {
		slice[i] = i
	}

	t.Logf("Initial: %s", SliceInfo(slice))

	// 删除大部分元素，触发缩容
	for i := 0; i < 90; i++ {
		var err error
		slice, err = DeleteAt(slice, 0)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
	}

	t.Logf("After deletion: %s", SliceInfo(slice))

	// 验证缩容是否发生
	if cap(slice) >= 100 {
		t.Error("Expected capacity to shrink, but it didn't")
	}

	if len(slice) != 10 {
		t.Errorf("Expected length 10, got %d", len(slice))
	}
}

func TestShouldShrink(t *testing.T) {
	tests := []struct {
		length   int
		capacity int
		expected bool
	}{
		{1, 8, true},    // 1 < 8/4
		{2, 8, false},   // 2 == 8/4
		{3, 8, false},   // 3 > 8/4
		{1, 4, false},   // capacity <= 4
		{10, 20, false}, // 10 > 20/4
	}

	for _, test := range tests {
		result := shouldShrink(test.length, test.capacity)
		if result != test.expected {
			t.Errorf("shouldShrink(%d, %d) = %v, expected %v",
				test.length, test.capacity, result, test.expected)
		}
	}
}

func TestSliceInfo(t *testing.T) {
	slice := make([]int, 5, 10)
	info := SliceInfo(slice)
	expected := "len=5, cap=10"
	if info != expected {
		t.Errorf("Expected %s, got %s", expected, info)
	}
}

// 性能基准测试
func BenchmarkDeleteAt(b *testing.B) {
	slice := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		slice[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		testSlice := make([]int, len(slice))
		copy(testSlice, slice)
		DeleteAt(testSlice, 500)
	}
}

func BenchmarkDeleteAtOptimized(b *testing.B) {
	slice := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		slice[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		testSlice := make([]int, len(slice))
		copy(testSlice, slice)
		DeleteAtOptimized(testSlice, 500)
	}
}
