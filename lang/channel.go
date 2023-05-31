package lang

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
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
