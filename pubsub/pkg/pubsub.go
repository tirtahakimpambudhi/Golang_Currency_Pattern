package pkg

import (
	"errors"
	"fmt"
	"sync"
)

var ErrNotFoundTopic = errors.New("topic not found")
type MapPubSub[T any] struct {
	pubsubs map[string]*PubSub[T]
	wg sync.WaitGroup
	mu sync.RWMutex
}
func NewMapPubSub[T any]() *MapPubSub[T] {
	return &MapPubSub[T]{
		pubsubs: make(map[string]*PubSub[T]),
	}
}
func (mps *MapPubSub[T]) Subscribe(topic string) {
	mps.mu.RLock() // Menggunakan RWMutex untuk read lock
	defer mps.mu.RUnlock()
	ps , ok := mps.pubsubs[topic]
	if !ok {
		ps = NewPubSub[T]()
		mps.pubsubs[topic] = ps
	}
	mps.wg.Add(1)
	s := ps.Subscribe()
	go func() {
		for {
			select {
			case value , ok := <- s:
				if !ok {
					fmt.Println("exiting")
					mps.wg.Done()
					return
				}
				fmt.Println(value)
			}
		}
	}()

}
func (mps *MapPubSub[T]) Publish(topic string, value T) {
	mps.mu.Lock() // Write lock untuk memodifikasi map
	defer mps.mu.Unlock()

	ps, ok := mps.pubsubs[topic]
	if !ok {
		ps = NewPubSub[T]()
		mps.pubsubs[topic] = ps
	}

	ps.Publish(value) // Publikasikan nilai
}

func (mps *MapPubSub[T]) Close() {
	mps.mu.Lock() // Write lock untuk memastikan map tidak dimodifikasi
	defer mps.mu.Unlock()

	for _, ps := range mps.pubsubs {
		ps.Close() // Tutup semua PubSubs
	}

	mps.wg.Wait() // Tunggu semua goroutine selesai
	fmt.Println("All goroutines have finished.")
}

type PubSub [T any] struct {
	subscribers []chan T
	close bool
	mu sync.RWMutex
}

func NewPubSub[T any]() *PubSub[T] {
	return &PubSub[T]{
		mu: sync.RWMutex{},
	}
}

func (ps *PubSub[T]) Subscribe() <-chan T  {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	if ps.close {
		return nil
	}
	r := make(chan T)
	ps.subscribers = append(ps.subscribers, r)
	return r
}

func (ps *PubSub[T]) Publish(value T)  {
	ps.mu.RLock()
	defer ps.mu.RUnlock()
	if ps.close {
		return
	}
	for _ , ch := range ps.subscribers {
		ch <- value
	}
}

func (ps *PubSub[T]) Close() {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	if ps.close {
		return
	}
	for _ , ch := range ps.subscribers {
		close(ch)
	}
	ps.close = true
}

