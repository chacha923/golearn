package example

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// 处理系统信号
func ListenSignal() {
	sigs := make(chan os.Signal, 1)

	// signal.Notify 注册给定的通道，用于接收特定信号。
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	done := make(chan bool, 1)

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")

	/*
		awaiting signal
		^C
		interrupt
		exiting
	*/
}
