package main

import (
	"sync"
)

func MergeChannels[T any](channels ...chan T) chan T {
	merged := make(chan T)

	go func() {
		var wg sync.WaitGroup
		wg.Add(len(channels))

		for _, ch := range channels {
			go func(ch chan T) {
				defer wg.Done()

				for val := range ch {
					merged <- val
				}
			}(ch)
		}

		wg.Wait()
		close(merged)
	}()

	return merged
}
