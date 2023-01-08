package pkg

import (
	"context"
	"sync"
)

func FanIn(ctx context.Context, out chan<- int, in ...<-chan int) {
	wg := sync.WaitGroup{}
	wg.Add(len(in))
	defer wg.Wait()
	for _, inChan := range in {
		go func(c <-chan int, waitGroup *sync.WaitGroup) {
			defer waitGroup.Done()
			for v := range c {
				out <- v
			}
		}(inChan, &wg)
	}
}
