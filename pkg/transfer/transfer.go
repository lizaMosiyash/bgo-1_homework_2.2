package transfer

import (
	"errors"
	"github.com/lizaMosiyash/bgo-1_homework-2.1/pkg/card"
	"strings"
)

var (
	ErrLowBalance = errors.New("Transfer impossible, low balance")
	ErrSourceCardNotFound = errors.New("Source card not found")
	ErrSourceCardNotExist = errors.New("Card not exist")
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
	source := s.SearchByNumber(from)
	target := s.SearchByNumber(to)
	cComission := (amount * s.Comission) / 10
	if cComission < s.MinComission {
		cComission = s.MinComission
	}
	total = cComission + amount*100
	sourceOk := strings.HasPrefix(from, "510621")
	targetOk := strings.HasPrefix(to, "510621")
	if sourceOk == false {
		return total, ErrSourceCardNotFound
	}
	if source == nil {
		return total, ErrSourceCardNotExist
	}
	if source.Balance < total {
		return total, ErrLowBalance
	}
	source.Balance -= total
	if targetOk == false {
		return total, nil
	}
	if target == nil {
		return total, ErrSourceCardNotExist
	}
	target.Balance += total
	return total, nil
}

func (s *Service) SearchByNumber(number string) *card.Card {
	for _, card := range s.CardSvc.Cards {
		if card.Number == number {
			return card
		}
	}
	return nil
}

func (s *Service) CheckBalance(number string, sum int64) error {
	c := s.SearchByNumber(number)
	if c.Balance < sum {
		return ErrLowBalance
	}
	c.Balance -= sum
	return nil
}
