package transfer

import (
	"errors"
	"fmt"
	"github.com/lizaMosiyash/bgo-1_homework-2.1/pkg/card"
)

var (
	ErrLowBalance = errors.New("Transfer impossible, low balance")
	ErrCardNotFound = errors.New("Card not found")
)

type Service struct {
	CardSvc      *card.Service
	Comission    int64
	MinComission int64
}

func NewService(cardSvc *card.Service, comission int64, minComission int64) *Service {
	return &Service{
		CardSvc:      cardSvc,
		Comission:    comission,
		MinComission: minComission,
	}
}

func (s *Service) Card2Card(from, to string, amount int64) (total int64, err error) {
	a,_ := s.SearchByNumber(from)
	b,_ := s.SearchByNumber(to)
	cComission := (amount * s.Comission) / 10
	if cComission < s.MinComission {
		cComission = s.MinComission
	}
	total = cComission + amount*100
	if a==nil || b==nil {
		return total, ErrCardNotFound
	}
	if a.Balance < total {
		return total, ErrLowBalance
	}
	a.Balance -= total
	fmt.Println(b)
	return total, nil


}

func (s *Service) SearchByNumber(number string) (c *card.Card, err error) {
	for _, card := range s.CardSvc.Cards {
		if card.Number == number {
			return card, nil
		}
	}
	return nil, ErrCardNotFound
}

func (s *Service) CheckBalance(number string, sum int64) error {
	c,_ := s.SearchByNumber(number)
	if c.Balance < sum {
		return ErrLowBalance
	}
	c.Balance -= sum
	return nil
}
