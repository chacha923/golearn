package lang

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

// channel 模拟一个消息队列，单生产者多消费者

var ch = make(chan int, 10)

var exitCh = make(chan bool) // 这里用于发出退出信号

var sum atomic.Uint32

var wg sync.WaitGroup // 控制主线程退出

type MQ struct {
	producer      *Producer
	consumers     []*Consumer
	consumerCount int
}

func NewMQ() *MQ {
	var mq = &MQ{
		producer:  NewProducer(),
		consumers: make([]*Consumer, 0),
	}
	return mq
}

func (m *MQ) Run() {
	wg.Add(2)
	m.AddConsumer(2)
	for i := 0; i < 10; i++ {
		m.producer.Produce()
	}
	wg.Wait()
}

func (m *MQ) Close() {
	// 停止生产
	m.producer.Close()
	// 关队列
	close(ch)
	// 发退出信号
	for i := 0; i < m.consumerCount; i++ {
		exitCh <- true
	}
	fmt.Println("sum:", sum.Load())
}

func (m *MQ) AddConsumer(count int) {
	for i := 0; i < count; i++ {
		var consumer = NewConsumer()
		m.consumers = append(m.consumers, consumer)
		go consumer.Consume()
	}
	m.consumerCount = count
}

type Producer struct {
}

func NewProducer() *Producer {
	return &Producer{}
}

func (p *Producer) Produce() {
	// 生产一个随机数
	// 将随机数放入 channel
	var randInt = rand.Int() % 100
	ch <- randInt
}

func (p *Producer) Close() {
	// 停止生产
}

type Consumer struct {
}

func NewConsumer() *Consumer {
	return &Consumer{}
}

// 这里需要异步执行，否则会阻塞
func (c *Consumer) Consume() {
	defer func() {
		wg.Done()
	}()
	for {
		select {
		case v, ok := <-ch:
			if ok {
				fmt.Println("consumered:", v)
				sum.Add(uint32(v))
			}
		case <-exitCh:
			fmt.Println("exit signal")
			return
		default:
		}
	}
}

// 1. 使用三个协程，每秒钟打印cat dog fish（要求：顺序不能变化，协程1打印cat，协程2打印dog，协程3打印fish）
// 打印 3 次
func printCatDogFish() {
	var wg = sync.WaitGroup{}
	wg.Add(3)

	// 3 个通道触发打印
	chCat := make(chan struct{}, 1)
	chDog := make(chan struct{}, 1)
	chFish := make(chan struct{}, 1) // 留缓冲，否则最后一个信号可能阻塞（也可以统计全局的打印次数）

	go printer(&wg, "cat", chCat, chDog)
	go printer(&wg, "dog", chDog, chFish)
	go printer(&wg, "fish", chFish, chCat)

	chCat <- struct{}{}

	wg.Wait()

	close(chCat)
	close(chDog)
	close(chFish)
}

func printer(wg *sync.WaitGroup, target string, cur chan struct{}, next chan struct{}) {
	var count = 3 // 计数

	defer wg.Done()
	for {
		<-cur
		fmt.Println(target)
		time.Sleep(time.Second)
		count--
		next <- struct{}{}
		if count == 0 {
			// 退出
			return
		}
	}
}

// N个 Goroutine循环打印数字 min - max
// ch1: 1 => ch2: 2 => ch3: 3 => ch1: 4 => ch2: 5 => ch3: 6
func PrintByNGoroutine() {
	var wg = sync.WaitGroup{}
	// 多少个Goroutine (1 - n)
	var n int = 3
	// 打印数字，从min - max
	var min = 1
	var max = 100

	var ch = make(chan int, 1)
	ch <- min

	wg.Add(n)
	// 正常可以考虑分 group 分组消费，但是 channel 原生不支持查看待消费元素
	// 考虑用一个全局信号控制谁消费
	var currentChIdx atomic.Int64
	currentChIdx.Store(0)

	var printer = func(wg *sync.WaitGroup, idx int64) {
		go func() {
			defer wg.Done()
			for {
				if currentChIdx.Load() != idx {
					continue
				}
				// 到我消费了
				if v, ok := <-ch; ok && v <= max {
					fmt.Println("ch:", idx, "print:", v)
					v++
					ch <- v
				} else {
					return
				}
				// 结束了，通知下一个 ch
				currentChIdx.Add(1)
				currentChIdx.CompareAndSwap(int64(n), 0)
			}
		}()
	}

	for i := 0; i < n; i++ {
		printer(&wg, int64(i))
	}

	wg.Wait()
}
