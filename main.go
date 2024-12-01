package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Мьютекс для общего ресурса - в данном случае количества билетов. Если его не будет,
var mutex sync.Mutex

// счетчик проданных билетов (потокобезопасный)
var count = int32(0)

func main() {

	// Доступное количество билетов - ЭТО ОБЩИЙ РЕСУРС!
	var tickets int = 500

	wg := sync.WaitGroup{}
	// Симулируем большое количество покупателей билетов.
	// Для простоты будем считать, что покупатель может приобрести один билет
	for userId := 0; userId < 1000; userId++ {

		// добавляем единичку в wg перед запсуком ОДНОЙ горутины!
		wg.Add(1)

		// Покупаем билет с ИД пользователя
		// запускаем горутину
		go buyTicket(&wg, userId, &tickets)

	}

	// ждем окончания работы всех горутин
	wg.Wait()
	fmt.Println(count)
}

// buyTicket Функция покупки билета определенным пользователем
func buyTicket(wg *sync.WaitGroup, userId int, remainingTickets *int) {
	// в конце сигнализируем в wg, что работа закончена
	defer wg.Done()
	// блокируем фрагмент кода, в котором используются общие ресурысы, с помощью mutex
	mutex.Lock()

	// Проверяем, сколько осталось билетов
	if *remainingTickets > 0 {
		// покупка билета
		*remainingTickets--
		// альтернативный счетчик проданных билетов (чтобы не считать с помощью grep)
		atomic.AddInt32(&count, 1)
		fmt.Printf("User %d purchased a ticket. Tickets remaining:%d\n", userId, *remainingTickets)
	} else {
		fmt.Printf("User %d found no ticket\n", userId)
	}
	// отпускаем блокировку
	mutex.Unlock()

}
