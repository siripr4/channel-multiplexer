package pkg

import (
	"context"
	"sync"
)

func FanOut(ctx context.Context, in <-chan int, out ...chan<- int) {
	wg := sync.WaitGroup{}
	wg.Add(len(in))
	for v := range in {
		for _, outChan := range out {
			go func(c chan<- int, waitGroup *sync.WaitGroup) {
				defer waitGroup.Done()
					c <- v
			}(outChan, &wg)
		}
	}
}
