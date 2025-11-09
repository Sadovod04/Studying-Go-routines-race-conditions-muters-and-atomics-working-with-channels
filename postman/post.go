package postman

import (
	"context"
	"sync"
	"time"

	"github.com/k0kubun/pp"
)

func Post(cxt context.Context, wq *sync.WaitGroup, tp chan<- string, n int, mail string) {
	defer wq.Done()
	for {
		select {
		case <-cxt.Done():
			pp.Println("Рабочий ден yномера ", n, " усе")
			return
		default:
			pp.Println("почтальон номер", n)
			time.Sleep(1 * time.Second)
			tp <- mail
			pp.Println("донес письмо -", mail)
		}
	}
}

func Postpol(cxt context.Context, postm int) <-chan string {
	tpm := make(chan string)
	wq := &sync.WaitGroup{}
	for i := 1; i <= postm; i++ {
		wq.Add(1)
		go Post(cxt, wq, tpm, i, posttomail(i))
	}
	go func() {
		wq.Wait()
		close(tpm)
	}()
	return tpm
}

func posttomail(postnum int) string {
	ptm := map[int]string{
		1: "привет семья",
		2: "Друг ха",
		3: "привет семья",
		4: "Друг ха",
	}

	mail, ok := ptm[postnum]
	if !ok {
		return "лотерея"
	}
	return mail
}
