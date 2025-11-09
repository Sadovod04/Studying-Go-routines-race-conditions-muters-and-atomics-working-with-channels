package miner

import (
	"context"
	"sync"
	"time"

	"github.com/k0kubun/pp"
)

func Miner(cxt context.Context,
	wq *sync.WaitGroup,
	tp chan<- int, num int,
	power int) {
	defer wq.Done()
	for {
		select {
		case <-cxt.Done():
			pp.Println("Рабочий день  ", num, " усе")
			return
		default:
			pp.Println("шахтер номер ", num)
			time.Sleep(1 * time.Second)
			tp <- power * 10
			pp.Println("закончил добычу номер", num, "количевство:", power*10)
		}
	}
}
func Minerpool(cxt context.Context, count int) <-chan int {
	caoltp := make(chan int)
	wq := &sync.WaitGroup{}
	for i := 1; i < count+1; i++ {
		wq.Add(1)
		go Miner(cxt, wq, caoltp, i, i*2)
	}
	go func() {
		wq.Wait()
		close(caoltp)
	}()
	return caoltp
}
