package main

import (
	"fmt"
	"sync"
)

func main() {
	var myMap sync.Map // Используем sync map

	arr1 := [...]string{"2", "4", "6", "8", "10"}
	arr2 := [...]int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup
	for i := range arr1 {
		wg.Add(1)
		go func(key string, value int) { // создаем горутину и записываем значение в map
			defer wg.Done()
			myMap.Store(arr1[i], arr2[i])
		}(arr1[i], arr2[i])
	}
	wg.Wait() // ожидаем завершения работы всех горутин
	myMap.Range(func(key, value interface{}) bool {
		fmt.Printf("Ключ: %v, Значение: %v\n", key, value)
		return true // Возврат true продолжает перебор
	})
}
