package workers

import (
	"sync"
	"testing"
	"time"
)

func Test_Workers(t *testing.T) {
	wg := New(50)
	var l sync.Mutex
	p := 0

	for i := 0; i < 1000; i++ {
		//	fmt.Println("Adding :", i)
		wg.Add(func() {
			time.Sleep(time.Second / 10)
			l.Lock()
			p++
			l.Unlock()

		})
	}
	wg.Wait()
	if p != 1000 {
		t.Log("P not reached it's goal")
		t.Fail()
	}

}
