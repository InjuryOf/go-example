package slide_window

import (
	"errors"
	"sync"
	"time"
)

type Bucket struct {
	Total   int
	Success int
	Fail    int
	next    *Bucket
	mux     sync.RWMutex
}

type BucketRing struct {
	buckets []*Bucket
	curent  *Bucket
	size    int
}

type Window struct {
	ring     *BucketRing
	size     int
	duration time.Duration
	interval time.Duration
	startTime time.Time
}

// 初始化滑动窗口
func NewWindow(size int, duration time.Duration) (*Window, error) {
	if size <= 0 {
		return nil, errors.New("size value error")
	}
	win := &Window{
		duration: duration,
		size:     size,
		ring:     NewBucketRing(size),
		interval: time.Duration(duration.Nanoseconds() / int64(size)),
		startTime: time.Now(),
	}
	return win, nil
}

// 初始化环形桶列表
func NewBucketRing(size int) *BucketRing {
	br := &BucketRing{
		buckets: make([]*Bucket, size),
		size:    size,
	}

	for i := 0; i < size; i++ {
		br.buckets[i] = &Bucket{}
	}
	br.curent = br.buckets[0]
	for i := 0; i < size; i++ {
		nextIndex := i + 1
		// 最后一个bucket指向第一个bucket，形成环形桶
		if nextIndex == size {
			nextIndex = 0
		}
		br.buckets[i].next = br.buckets[nextIndex]
	}
	return br
}

func (b *Bucket) IncrSuccess(value int) {
	b.mux.Lock()
	defer b.mux.Unlock()
	b.Success += value
	b.Total += value
}

func (b *Bucket) IncrFail(value int) {
	b.mux.Lock()
	defer b.mux.Unlock()
	b.Fail += value
	b.Total += value
}

// 向右轮转bucket
func (br *BucketRing) RightRotate() {
	br.curent.next.Reset()
	br.curent = br.curent.next
}

// 重置桶内计数
func (b *Bucket) Reset() {
	b.Success = 0
	b.Fail = 0
	b.Total = 0
}

func (w *Window) Roll(f RollFunc) {
	// 如果当前时间减去启动时间大于间隔时间，则滚动到下一个桶计数
	if time.Now().Sub(w.startTime) >= w.interval {
		w.ring.RightRotate()
	}
	f(w.ring.curent)
}

// 统计
func (w *Window) Sum(field string) int {
	var sum int
	for i := 0; i < len(w.ring.buckets); i++ {
		switch field {
		case "Success":
			sum += w.ring.buckets[i].Success
		case "Fail":
			sum += w.ring.buckets[i].Fail
		default:
			sum += w.ring.buckets[i].Total
		}
	}
	return sum
}

func IncrOutSuccess(bucket *Bucket) {
	bucket.IncrSuccess(1)
}

func IncrOutFail(bucket *Bucket) {
	bucket.IncrFail(1)
}

type RollFunc func(bucket *Bucket)