package main

import (
	"context"
	"posmin/miner"
	"posmin/postman"
	"sync"
	"sync/atomic"
	"time"

	"github.com/k0kubun/pp"
)

func main() {
	var coal atomic.Int64
	mtx := sync.Mutex{}
	var mails []string
	minercxt, minercansel := context.WithCancel((context.Background()))
	postncxt, postcansel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(3 * time.Second)
		postcansel()
	}()
	go func() {
		time.Sleep(6 * time.Second)
		minercansel()
	}()

	caoltp := miner.Minerpool(minercxt, 10)
	mailptp := postman.Postpol(postncxt, 10)

	// красиво
	wq := &sync.WaitGroup{}
	wq.Add(1)

	go func() {
		defer wq.Done()
		for x := range caoltp {
			coal.Add(int64(x))
		}
	}()
	wq.Add(1)
	go func() {
		defer wq.Done()
		for z := range mailptp {
			mtx.Lock()
			mails = append(mails, z)
			mtx.Unlock()
		}
	}()
	wq.Wait()

	pp.Println("summ caol:", coal.Load())
	mtx.Lock()
	pp.Println("massange:", len(mails))
	mtx.Unlock()
}
