package main

import "time"

func RateLimit(f func(), rps int) func() {
	tick := time.Tick(time.Second / time.Duration(rps))

	return func() {
		<-tick
		f()
	}
}
