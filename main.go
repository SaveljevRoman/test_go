package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go say(&wg, 1, "go is awesome")

	wg.Add(2)
	go say(&wg, 2, "cats are cute")
	wg.Wait()
}

func say(wg *sync.WaitGroup, id int, phrase string) {
	for _, word := range strings.Fields(phrase) {
		fmt.Printf("Worker #%d says: %s...\n", id, word)
		dur := time.Duration(rand.Intn(100)) * time.Microsecond
		time.Sleep(dur)
	}
	wg.Done()
}

//Горутины
//ф-ия main() является горутиной!!!
//При завершении main завершаются и все остальные горутины, поэтому в примере пишем sleep
//Чтобы вызвать метод асинхронно необходимо перед вызовом метода написать go
//func main() {
//	go say(1, "go is awesome")
//	go say(2, "cats are cute")
//	time.Sleep(500 * time.Microsecond)
//}

//Группа ожидания
//Чтобы объявить счетчик в группе ожидания нужно - var wg sync.WaitGroup
//	увеличивает счетчик - wg.Add(1)
//	уменьшает счетчик на единицу - wg.Done()
//  блокирование горутины, до тех пор пока счетчик не обнулится - wg.Wait()
// Пример
// defer и анонимная ф-ия - используется для отделения асинхронности и бизнес логики
//func main() {
//	var wg sync.WaitGroup
//	wg.Add(2)
//
//	go func() {
//		defer wg.Done()
//		say(1, "go is awesome")
//	}()
//
//	go func() {
//		defer wg.Done()
//		say(2, "cats are cute")
//	}()
//
//	wg.Wait()
//}
//
//func say(id int, phrase string) {
//	for _, word := range strings.Fields(phrase) {
//		fmt.Printf("Worker #%d says: %s...\n", id, word)
//		dur := time.Duration(rand.Intn(100)) * time.Millisecond
//		time.Sleep(dur)
//	}
//}
