package map_concurrent

import (
	"sync"
	"testing"
)

func TestConcurrentRead(t *testing.T) {
	wg := &sync.WaitGroup{}
	a := map[string]string{"foo": "bar"}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			_ = a["foo"]
			wg.Done()
		}()
	}
	wg.Wait()
}

func TestConcurrentWrite(t *testing.T) {
	wg := &sync.WaitGroup{}
	a := map[string]string{"foo": "bar"}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			a["foo"] = "blah"
			wg.Done()
		}()
	}
	wg.Wait()
}
