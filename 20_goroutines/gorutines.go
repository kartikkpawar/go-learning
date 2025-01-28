package main

import (
	"fmt"
	"sync"
)

func task(id int, wg *sync.WaitGroup) {

	// defer keyword: whatever return after defer will execute after the completion of the function
	defer wg.Done()
	fmt.Println("Doing Task", id)
}

// NOTE: If the main fuction is exited before the goroutine completes its execution then in such cases all the goroutines are cleared ir-respective of whether or not the go-routines funtion is executed
// To avoid this senario we use waitGroup
// 1. Add to waitGroup 2. Mark waitgroup as done as required 3. wait for waitgroup to execute fully

func main() {

	var waitGroup sync.WaitGroup

	for i := 0; i <= 10; i++ {
		waitGroup.Add(1)
		go task(i, &waitGroup)
	}

	// time.Sleep(time.Second * 2) // Simulating delay
	waitGroup.Wait() // waits till all the functions inside waitgroup are executed/done

}
