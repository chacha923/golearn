package lang

import (
	"context"
	"errors"
	"fmt"
	"go-common/library/sync/errgroup"
	"runtime"
	"time"
)

func UseGroup() {
	fmt.Println("use group start")
	g := errgroup.Group{}
	g.GOMAXPROCS(runtime.NumCPU())
	g.Go(foo1)
	g.Go(foo1WithErr)

	if err := g.Wait(); err != nil {
		fmt.Println("g.wait catch error:", err)
	}

	time.Sleep(10 * time.Second)
}

var (
	rootCtx = context.Background()
	cv      = context.WithValue(rootCtx, "key", "value")
	c, _    = context.WithCancel(cv)
)

func UseWithContext() {
	fmt.Println("use with context start")
	g, ctx := errgroup.WithContext(c)
	g.Go(foo1WithErr)
	g.Go(foo2)
	g.Go(foo3)
	if err := g.Wait(); err != nil {
		fmt.Println("g.wait catch error:", err)
		fmt.Println("ctx error:", ctx.Err())
	}

	time.Sleep(10 * time.Second)
}

func foo1() error {
	fmt.Println("in foo1")
	return nil
}

func foo1WithErr() error {
	fmt.Println("in foo1WithErr")
	return errors.New("foo1 error")
}

func foo2() error {
	fmt.Println("in foo2")
	time.Sleep(3 * time.Second)
	fmt.Println("end foo2")
	return nil
}

func foo3() error {
	fmt.Println("in foo3")
	time.Sleep(3 * time.Second)
	fmt.Println("kv:", c.Value("key"))
	fmt.Println("end foo3")
	return nil
}
