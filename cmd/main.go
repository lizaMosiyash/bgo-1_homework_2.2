package main

import (
	"fmt"
	"github.com/lizaMosiyash/bgo-1_homework-2.1/pkg/card"
	"github.com/lizaMosiyash/bgo-1_homework-2.1/pkg/transfer"
)

func main() {
	svc := card.NewService("MyBank")
	card := svc.IssueCard("11111111", 555500)
	card = svc.IssueCard("22222222", 30000)

	a := transfer.NewService(svc, 5, 1000)
	fmt.Println(svc, card, a)
	svf := transfer.NewService(svc, 5, 1000)
	total, ok := svf.Card2Card("11111111", "22222222", 1000)
	fmt.Println(total, ok)

}
