package loop

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
)

// 并发回调列队，错误不处理
func QueueConcurrentCallbackAllDone[T any](ctx context.Context, arr []T, limit int, cb func(T)) error {
	return queueConcurrentCallback(ctx, arr, limit, func(t T) error {
		cb(t)
		return nil
	}, false)
}

// 并发回调列队遇 error 结束
func QueueConcurrentCallback[T any](ctx context.Context, arr []T, limit int, cb func(T) error) error {
	return queueConcurrentCallback(ctx, arr, limit, cb, true)
}

// 并发回调列队
func queueConcurrentCallback[T any](ctx context.Context, arr []T, limit int, cb func(T) error, all_done bool) error {
	// 需要处理的任务总个数
	taskTotal := int64(len(arr))
	// 完成任务的总个数
	var doneCount int64
	// 结束标识
	doneCh := make(chan error)
	// 有长度的管道模拟列队
	ch := make(chan struct{}, limit)
	// 错误列表
	var errs []error
	var errs_lock sync.RWMutex

	for _, t := range arr {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case err := <-doneCh:
			return err
		case ch <- struct{}{}:
			go func(_t T) {
				if err := cb(_t); err != nil {
					if all_done {
						doneCh <- err
						return
					}
					func() {
						errs_lock.Lock()
						defer errs_lock.Unlock()
						errs = append(errs, err)
					}()
				}
				if newDoneCount := atomic.AddInt64(&doneCount, 1); newDoneCount == taskTotal {
					// 全部任务结束，通知结束 channel
					doneCh <- errors.Join(errs...)
				} else {
					<-ch
				}
			}(t)
		}
	}
	return <-doneCh
}
