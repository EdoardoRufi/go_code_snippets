package datastructures

import (
	"fmt"
	"sync"
)

type Stack[T any] []T

type Flight struct {
	Origin      string
	Destination string
	Price       int
}

func (s *Stack[T]) Push(x T) {
	*s = append(*s, x)
}

func (s *Stack[T]) IsEmpty() bool {
	return len(*s) == 0
}

// Pop removes the last pushed element and returns it;
// it returns false if the stack is empty.
func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	if s.IsEmpty() {
		return zero, false
	}
	last := len(*s) - 1
	v := (*s)[last]
	*s = (*s)[:last] // :last means that last index is excluded
	return v, true
}

// Peek returns the top element of the stack without removing it.
// It returns false if the stack is empty.
func (s *Stack[T]) Peek() (T, bool) {
	var zero T
	if s.IsEmpty() {
		return zero, false
	}
	return (*s)[len(*s)-1], true
}

func (s *Stack[T]) Len() int {
	return len(*s)
}

func LaunchStack() {
	var flights Stack[Flight]

	flights.Push(Flight{
		Origin:      "ROM",
		Destination: "NYC",
		Price:       500,
	})
	flights.Push(Flight{
		Origin:      "MIL",
		Destination: "PAR",
		Price:       150,
	})

	top, ok := flights.Peek()
	if ok {
		fmt.Println("Top:", top)
	}

	popped, ok := flights.Pop()
	if ok {
		fmt.Println("Popped:", popped)
	}

	popped, ok = flights.Pop()
	if ok {
		fmt.Println("Popped:", popped)
	}

	//nothing to pop
	popped, ok = flights.Pop()
	if ok {
		fmt.Println("Popped:", popped)
	}

	fmt.Println("Len:", flights.Len())
}

type ConcurrentStack[T any] struct {
	mu   sync.Mutex
	data []T
}

func (s *ConcurrentStack[T]) Push(v T) {
	s.mu.Lock()
	s.data = append(s.data, v)
	s.mu.Unlock()
}

func (s *ConcurrentStack[T]) Pop() (T, bool) {
	var zero T
	if len(s.data) == 0 {
		return zero, false
	}
	last := len(s.data) - 1
	v := s.data[last]
	s.data = s.data[:last] // :last means that last index is excluded
	return v, true
}

func (s *ConcurrentStack[T]) ConcurrentPop() (T, bool) {
	var zero T

	s.mu.Lock()
	defer s.mu.Unlock()

	n := len(s.data)
	if n == 0 {
		return zero, false
	}
	v := s.data[n-1]
	s.data = s.data[:n-1]
	return v, true
}

func (s *ConcurrentStack[T]) Peek() (T, bool) {
	var zero T

	s.mu.Lock()
	defer s.mu.Unlock()

	n := len(s.data)
	if n == 0 {
		return zero, false
	}
	return s.data[n-1], true
}

func (s *ConcurrentStack[T]) Len() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.data)
}

func LaunchConcurrentStack() {
	var st ConcurrentStack[int]

	// --- 1) 4 concurrent go routines push from 1 to 20

	var wgPush sync.WaitGroup

	// each routine gets a range: [1..5], [6..10], [11..15], [16..20]
	ranges := [][2]int{
		{1, 5},
		{6, 10},
		{11, 15},
		{16, 20},
	}

	for _, r := range ranges {
		wgPush.Add(1)
		start, end := r[0], r[1]

		go func(a, b int) {
			defer wgPush.Done()
			for i := a; i <= b; i++ {
				st.Push(i)
				fmt.Printf("PUSH: %d\n", i)
			}
		}(start, end)
	}

	// wait fo all the go routines to finish
	wgPush.Wait()

	// --- 2) print everything in the stack ---

	snap := st.Snapshot()
	fmt.Println("\nStack after all pushes (order is LIFO, non deterministic):")
	fmt.Println(snap)

	// --- 3) go routines pop everything

	// (1)
	//problem: second go routine pops and blocks the mutex. when the mutex is free, is faster than the other to take another one and relock it
	//So the other ones are alwyas stuck at the first item. One go routine does all the rest of the job.

	// (2)
	//possible solution: pop without mutex and only one go routine popping. each thing popped is sent in a channel, where a workerpool perform his action
	//problem: the pop will be in order, but sending it to a worker pool of go routine will make me lose the pop order.
	// so soultion 1 remains kind of the best, even tough the multiple worker popping is useless and wrong

	// but it's been fun to do these tryings :)

	/* start (1)*/
	var wgPop sync.WaitGroup
	const numPoppers = 3

	for id := 0; id < numPoppers; id++ {
		wgPop.Add(1)
		go func(id int) {
			defer wgPop.Done()
			for {
				v, ok := st.ConcurrentPop()
				if !ok {
					fmt.Printf("ending go routine %d because there is nothing to pop\n", id)
					return
				}
				fmt.Printf("POP (goroutine %d): %d\n", id, v)
			}
		}(id)
	}

	wgPop.Wait()
	/* end (1) */

	/* start (2) */
	// 2) Worker pool
	// const workers = 3
	// toPerform := make(chan int)
	// var wg sync.WaitGroup

	// for id := 0; id < workers; id++ {
	// 	wg.Add(1)
	// 	go func(id int, toPerform <-chan int) {
	// 		defer wg.Done()
	// 		for v := range toPerform {
	// 			fmt.Printf("worker %d processing %d\n", id, v)
	// 		}
	// 	}(id, toPerform)
	// }

	// // 3) Only one go routine doing pop in LIFO order
	// go func() {
	// 	for {
	// 		v, ok := st.Pop()
	// 		if !ok {
	// 			close(toPerform) // nothing to do anymore
	// 			return
	// 		}
	// 		toPerform <- v
	// 	}
	// }()

	// wg.Wait()

	/* end (2) */

	fmt.Println("\nDone, remaining len:", st.Len())
}

// Snapshot returns a copy of the current stack contents.
func (s *ConcurrentStack[T]) Snapshot() []T {
	s.mu.Lock()
	defer s.mu.Unlock()

	out := make([]T, len(s.data))
	copy(out, s.data)
	return out
}
