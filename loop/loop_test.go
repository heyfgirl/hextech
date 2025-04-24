package loop

import (
	"context"
	"errors"
	"sync/atomic"
	"testing"
)

func TestQueueConcurrentCallbackAllDone(t *testing.T) {
	// 测试 QueueConcurrentCallbackAllDone 函数
	ctx := context.Background()
	arr := []int{1, 2, 3, 4, 5}
	limit := 2

	var result int64
	cb := func(t int) {
		// 模拟处理
		atomic.AddInt64(&result, int64(t))
	}
	err := QueueConcurrentCallbackAllDone(ctx, arr, limit, cb)
	if err != nil {
		t.Errorf("QueueConcurrentCallbackAllDone failed: %v", err)
	}
	// 检查结果
	expected := int64(15) // 1 + 2 + 3 + 4 + 5
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func TestQueueConcurrentCallback(t *testing.T) {
	// 测试 QueueConcurrentCallback 函数
	ctx := context.Background()
	arr := []int{1, 2, 3, 4, 5}
	limit := 2

	var result int64
	// 模拟处理函数，返回错误
	e := errors.New("error")
	cb := func(t int) error {
		if t == 3 {
			return e
		}
		// 模拟处理
		atomic.AddInt64(&result, int64(t))
		return nil
	}
	err := QueueConcurrentCallback(ctx, arr, limit, cb)
	if err != nil {
		if errors.Is(err, e) {
			// 预期错误
		} else {
			t.Errorf("QueueConcurrentCallback failed: %v", err)
		}
	}
	// 检查结果
	expected := 15 // 1 + 2 + 3 + 4 + 5
	if result == int64(expected) {
		t.Errorf("expected %d, got %d", expected, result)
	}
}
