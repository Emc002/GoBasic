//note that if using using gorouitnes to make something like this from the
// the unordered list and you manipulate the the flow so you can get the
// ordered list it makes the goroutines benefit it lost cause we now is
// running this code in synchronus pattern using mutex rw the choice is yours

package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}

func main() {
	runtime.GOMAXPROCS(100)
	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock()
		go sayHello()
		m.Lock()
		go increment()
	}
	wg.Wait()
}

func sayHello() {
	fmt.Println("hello", counter)
	m.RUnlock()
	wg.Done()

}

func increment() {
	counter++
	m.Unlock()
	wg.Done()
}
