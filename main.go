package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Hello, World!")

	// Доступное количество билетов
	var tickets int = 500

	wg := sync.WaitGroup{}

	// Симулируем большое количество покупателей билетов.
	// Для простоты будем считать, что покупатель может приобрести один билет
	for userId := 0; userId < 2000; userId++ {

		// добавляем единичку в wg перед запсуком ОДНОЙ горутины!
		wg.Add(1)

		// Покупаем билет с ИД пользователя
		// запускаем горутину
		go buyTicket(&wg, userId, &tickets)

	}

	// ждем окончания работы всех горутин
	wg.Wait()
}

// buyTicket Функция покупки билета определенным пользователем
func buyTicket(wg *sync.WaitGroup, userId int, remainingTickets *int) {
	// в конце сигнализируем в wg, что работа закончена
	defer wg.Done()

	// Проверяем, сколько осталось билетов
	if *remainingTickets > 0 {
		// покупка билета
		*remainingTickets--
		fmt.Printf("User %d purchased a ticket. Tickets remaining:%d\n", userId, *remainingTickets)
	} else {
		fmt.Printf("User %d found no ticket\n", userId)
	}

}
