package main

import (
	"fmt"
	"sync"
	"time"
)

// Расскажи подробно что происходит
// Сколько времени будет выполняться?
const numRequests = 10000

var count int

var m sync.Mutex

func networkRequest() {
	time.Sleep(time.Millisecond) // Эмуляция сетевого запроса.
	m.Lock()
	count++
	m.Unlock()
}

func main() {
	var wg sync.WaitGroup

	wg.Add(numRequests)
	for i := 0; i < numRequests; i++ {
		go func() {
			defer wg.Done()
			networkRequest()
		}()
	}

	wg.Wait()
	fmt.Println(count)
}

/*
Объявили константу numRequests = 10000, объявили переменную типа int,
объявили переменную типа sync.Mutex, объявили функцию networkRequest(), которая эмулирует сетевой запрос и
увеличивает счетчик в мьютексе, в main() объявили переменную wg типа sync.WaitGroup, в wg.Add(numRequests)
указали количество горутин (10000), создали цикл, который каждую итерацию создает горутину и вызывает функцию networkRequest(),
ждем завершение всех горутин и выводим счетчик равный 10000

Будет выполняться не дольше 0.1 секунды
*/
