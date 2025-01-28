// Channels are used to transfer data inBetween go-routines
package main

import "fmt"

func main() {

	// messageChan := make(chan string) // making channel
	// NOTE: in-normal channels both reciving and sending parts are blocking in behaviour

	// messageChan <- "Ping..." // Adding data to the channel

	// channelMessage := <-messageChan // Reciving data from channel

	// fmt.Println(channelMessage)

	// numChan := make(chan int)
	// go processNum(numChan)

	// for {
	// 	numChan <- rand.Intn(100)
	// 	time.Sleep(time.Second)
	// }

	// resultChan := make(chan int)

	// go sum(resultChan, 4, 5)
	// fmt.Println("Result...", <-resultChan) // reading is also blocking in behaviour thats why no sleep or waitgroup used

	// waitgroup behaviour using channel
	// done := make(chan bool)
	// go task(done)

	// <-done // awating / holding the program till something is added to channel

	// BUFFERED_CHANNELS
	// In this we can send limited amount of data without having blocking behaviour

	emailChan := make(chan string, 100)
	done := make(chan bool)

	go emailSender(emailChan, done)

	for i := range 100 {
		emailChan <- fmt.Sprintf("%d@example.com", i)
	}

	close(emailChan) // closing channel in case of buffered channel to avoid deadlock error
	<-done

	// Listening to multiple channels

	// intChan := make(chan int)
	// strChan := make(chan string)

	// go func() {
	// 	intChan <- 3
	// }()
	// go func() {
	// 	strChan <- "golang"
	// }()

	// for i := 0; i < 2; i++ {
	// 	select {
	// 	case chan1Value := <-intChan:
	// 		fmt.Println("Recived Data from intChan->", chan1Value)

	// 	case chan2Value := <-strChan:
	// 		fmt.Println("Recived Data from intChan->", chan2Value)
	// 	}

	// }

}

// by using <- befor chan it make the channel as recive only.  From this func data cannot be sent to that particular channel. It is mostly for type-safety

//  And to make chan send only <- after chan

func emailSender(emailChan <-chan string, done chan<- bool) {
	defer func() { done <- true }()

	// deadlock error because of emailchan is used in range which is working as infinite loop. So closing of channel is required
	for email := range emailChan {
		fmt.Println("Sending email to ->", email)
	}
}

// func processNum(numChan chan int) {
// 	for num := range numChan {
// 		fmt.Println("Processing number.....", num)
// 	}
// }

// func sum(rc chan int, num1, num2 int) {
// 	result := num1 + num2
// 	rc <- result
// }

func task(done chan bool) {
	defer func() { done <- true }()
	fmt.Println("Processing...")
}
