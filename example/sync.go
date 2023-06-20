package example

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func AtomicAdd() {
	var ops uint64

	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			for c := 0; c < 1000; c++ {
				atomic.AddUint64(&ops, 1)
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("ops:", ops)
	// ops: 50000
}

func Mutex() {
	var state = make(map[int]int)
	var mutex = &sync.Mutex{}
	var ops atomic.Int64 = 0

	for r := 0; r < 100; r++ {
		go func() {
			var total = 0
			for {
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				ops.Add(1)

				// 让出时间片
				runtime.Gosched()
			}
		}()
	}

	for w := 0; w < 10; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				ops.Add(1)

				runtime.Gosched()
			}
		}()
	}

	time.Sleep(time.Second)

	opsFinal := ops.Load()
	fmt.Println("ops:", opsFinal)

	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()
}

// 状态协程
// 在这个例子中，state 将被一个单独的协程拥有。 这能保证数据在并行读取时不会混乱。
// 为了对 state 进行读取或者写入， 其它的协程将发送一条数据到目前拥有数据的协程中， 然后等待接收对应的回复。
// 结构体 readOp 和 writeOp 封装了这些请求，并提供了响应协程的方法。
func StateCorountine() {
	var readOps uint64
	var writeOps uint64

	reads := make(chan readOp)
	writes := make(chan writeOp)

	// 处理读写请求，用到 channel 串行处理，所以 state 不会出现并发问题
	go func() {
		var state = make(map[int]int)

		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	// 100 个协程循环发起读请求
	for r := 0; r < 100; r++ {
		go func() {
			for {
				// 构造读请求
				read := readOp{
					key:  rand.Intn(5),
					resp: make(chan int),
				}
				reads <- read
				<-read.resp
				// 读操作计数+1
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// 10 个协程循环发起写请求
	for w := 0; w < 10; w++ {
		go func() {
			write := writeOp{
				key:  rand.Intn(5),
				val:  rand.Intn(100),
				resp: make(chan bool),
			}
			// 发起写请求
			writes <- write
			<-write.resp
			atomic.AddUint64(&writeOps, 1)
			time.Sleep(time.Millisecond)
		}()
	}

	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)

	// readOps: 71708
	// writeOps: 7177
}

// 读请求
type readOp struct {
	key  int
	resp chan int // 传递读返回
}

// 写请求
type writeOp struct {
	key  int
	val  int
	resp chan bool // 传递写返回
}
