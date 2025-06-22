package waitgr

import (
	"fmt"
	"sync"
	"time"
)

var mu sync.Mutex

func schetchik(id int, wg *sync.WaitGroup) {
	mu.Lock()
	fmt.Printf("наступило %d апреля, долг шляпы вырос\n", id)
	mu.Unlock()

	time.Sleep(time.Second)

	mu.Lock()
	fmt.Printf("Worker %d done\n", id)
	mu.Unlock()
	wg.Done()
}

func waitgruppa() {
	var wg sync.WaitGroup
	for i := 1; i <= 29; i++ {
		wg.Add(1)
		go schetchik(i, &wg)
	}
	wg.Wait()
}

func main() {
	waitgruppa()
}
func gracefulShutdown() {
	fmt.Println("Оставновка расчета долга шляпы")
	time.Sleep(3 * time.Second)

	fmt.Println("Программа завершена")
}
