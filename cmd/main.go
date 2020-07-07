package main

import (
	"fmt"
	"github.com/lizaMosiyash/bgo-1_homework-2.1/pkg/card"
	"github.com/lizaMosiyash/bgo-1_homework-2.1/pkg/transfer"
)

func main() {
	svc := card.NewService("MyBank")
	card := svc.IssueCard("11111111", 555500)
	card = svc.IssueCard("22222222", 300)

	a := transfer.NewService(svc, 5, 1000)
	fmt.Println(svc, card, a)
	svf := transfer.NewService(svc, 5, 1000)

	_, err := svf.Card2Card("2", "11111111", 1000)
	if err != nil {
		switch err {
		case transfer.ErrLowBalance:
			fmt.Print("Недостаточно средств для перевода")
		case transfer.ErrCardNotFound:
			fmt.Println("Карта не найдена")
		default:
			fmt.Print("Перевод осуществлен успешно")
		}
	}


}


