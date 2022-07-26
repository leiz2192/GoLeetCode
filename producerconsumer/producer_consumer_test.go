package producerconsumer_test

import (
	"fmt"
	"sync"
	"testing"
)

func TestProducerConsumer(t *testing.T) {
	wg := sync.WaitGroup{}
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 28; i++ {
			<-ch1
			fmt.Printf("%d", i)
			ch2 <- struct{}{}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 'A'; i <= 'Z'; i++ {
			<-ch2
			fmt.Printf("%c", i)
			ch1 <- struct{}{}
		}
		<-ch2
		ch1 <- struct{}{}
		<-ch2
	}()
	ch1 <- struct{}{}
	wg.Wait()
}
