package main

import (
	"fmt"
	"github.com/lizaMosiyash/bgo-1_homework-2.1/pkg/card"
	"github.com/lizaMosiyash/bgo-1_homework-2.1/pkg/transfer"
)

func main() {
	svc := card.NewService("MyBank")
	card := svc.IssueCard("51062122", 555500)
	card = svc.IssueCard("51062128", 300)

	a := transfer.NewService(svc, 5, 1000)
	fmt.Println(svc, card, a)
	svf := transfer.NewService(svc, 5, 1000)

	_, err := svf.Card2Card("51062122", "11111111", 1000)
	if err != nil {
		switch err {
		case transfer.ErrLowBalance:
			fmt.Print("Недостаточно средств для перевода")
		case transfer.ErrSourceCardNotFound:
			fmt.Println("Для перевода необходимо указать номер карты нашего банка")
		case transfer.ErrSourceCardNotExist:
			fmt.Print("Карты с указанным номером не выпущено")
		default:
			fmt.Print("Что-то пошло не так...")
		}
	}
	fmt.Println("Перевод осуществлен")


}


