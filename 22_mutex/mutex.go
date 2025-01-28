package main

import (
	"fmt"
	"sync"
)

type post struct {
	views int
	mu    sync.Mutex
}

func (p *post) viewsInc(w *sync.WaitGroup) {
	defer func() {
		w.Done()
		p.mu.Unlock()
	}()

	p.mu.Lock()
	p.views += 1
}

func main() {
	var wg sync.WaitGroup

	myPost := post{
		views: 0,
	}

	for range 10000 {
		wg.Add(1)
		go myPost.viewsInc(&wg)
	}

	wg.Wait()
	fmt.Println(myPost.views)
}
