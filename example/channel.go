package example

import (
	"fmt"
	"time"
)

func chanBlock() {
	messages := make(chan string)

	go func() {
		messages <- "ping"
	}()

	// channel 无缓冲，会阻塞
	msg := <-messages
	println(msg)
}

func master() {
	var done = make(chan bool, 1)

	var worker = func() {
		fmt.Println("working...")
		time.Sleep(time.Second)
		fmt.Println("done")

		done <- true
	}

	go worker()

	<-done
}

func selectChan() {
	var c1 = make(chan string)
	var c2 = make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}

func timeOut() {
	var c1 = make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		// 1s 后捕获一个信号
		fmt.Println("timeout 1")
	}

	var c2 = make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 3")
	}

	// timeout 1
	// result 2
}

func selectDefault() {
	messages := make(chan string)
	signals := make(chan bool) // 退出信号

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}

	// no message received
	// no message sent
	// no activity
}

// 使用协程与通道实现一个工作池
func WorkPool() {
	var numJobs = 5
	var jobs = make(chan int, numJobs)
	var results = make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		// 遍历会阻塞，除非 channel 被关闭，并且消费完！！！
		go worker(w, jobs, results)
	}

	// 发送任务
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// 接收结果
	for a := 1; a <= numJobs; a++ {
		<-results
	}

}

// 模拟一个耗时的任务
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}
