package transfer

import (
	"github.com/lizaMosiyash/bgo-1_homework-2.1/pkg/card"
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

func (s *Service) Card2Card(from, to string, amount int64) (total int64, ok bool) {
	a := s.SearchByNumber(from)
	b := s.SearchByNumber(to)
	cComission := (amount * s.Comission) / 10
	if cComission < s.MinComission {
		cComission = s.MinComission
	}
	total = cComission + amount*100
	if a != nil && b != nil {
		if a.Balance >= total {
			a.Balance -= total
			b.Balance += total
			return total, true
		}
		if a.Balance < total {
			return total, false
		}
		return total, ok
	}
	if a != nil && b == nil {
		if a.Balance >= total {
			a.Balance -= total
		}
	}
	if a == nil && b != nil {
		b.Balance += total
	}
	return total, true
}

func (s *Service) SearchByNumber(number string) *card.Card {
	for _, card := range s.CardSvc.Cards {
		if card.Number == number {
			return card
		}
	}
	return nil
}
