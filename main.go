package main



import (
	"fmt"
	"sync"



	
)

var x = 0
func adder() {
	x = x + 1
}

func main() {
	fmt.Println("asd")
	//x := 0
	wg := sync.WaitGroup{}
	m := sync.RWMutex{}
	for i:=0 ; i<100 ; i++ {
		wg.Add(1)
		go func() {
			m.Lock()
			adder()
			wg.Done()
			m.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(x)



}

