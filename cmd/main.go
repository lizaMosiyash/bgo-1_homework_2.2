package main

import (
	"fmt"
	"github.com/lizaMosiyash/bgo-1_homework-2.1/pkg/card"
	"github.com/lizaMosiyash/bgo-1_homework-2.1/pkg/transfer"
)

func main() {
	svc := card.NewService("MyBank")
	card := svc.IssueCard("4456 6180 5182 7953", 555500)
	card = svc.IssueCard("4716 6942 5741 4323", 300)

	a := transfer.NewService(svc, 5, 1000)
	fmt.Println(svc, card, a)
	svf := transfer.NewService(svc, 5, 1000)

	_, err := svf.Card2Card("4456 6180 5182 7954", "4716 6942 5741 4323", 1000)
	if err != nil {
		switch err {
		case transfer.ErrLowBalance:
			fmt.Print("Недостаточно средств для перевода")
		case transfer.ErrSourceCardNotFound:
			fmt.Println("Для перевода необходимо указать номер карты нашего банка")
		case transfer.ErrSourceCardNotExist:
			fmt.Print("Карты с указанным номером не выпущено")
		case transfer.ErrInvalidCard:
			fmt.Println("Некорректные данные карты")
		default:
			fmt.Print("Что-то пошло не так...")
		}
	}
	svf.IsValid(svc.Cards[0].Number)


}


