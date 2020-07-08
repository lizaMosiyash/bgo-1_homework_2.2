package transfer

import (
	"errors"
	"github.com/lizaMosiyash/bgo-1_homework-2.1/pkg/card"
	"strconv"
	"strings"
)

var (
	ErrLowBalance = errors.New("Transfer impossible, low balance")
	ErrSourceCardNotFound = errors.New("Source card not found")
	ErrSourceCardNotExist = errors.New("Card not exist")
	ErrInvalidCard = errors.New("Invalid card")
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
	cComission := (amount * s.Comission) / 10
	if cComission < s.MinComission {
		cComission = s.MinComission
	}
	total = cComission + amount*100
	_, ok := s.IsValid(from)
	_, ok1 := s.IsValid(to)
	if (ok || ok1) != true  {
		return total, ErrInvalidCard
	}
	sourceOk := strings.HasPrefix(from, "4456 61")
	targetOk := strings.HasPrefix(to, "4456 61")
	source := s.SearchByNumber(from)
	target := s.SearchByNumber(to)
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

func (s *Service) IsValid(number string) (int, bool) {
	check := strings.ReplaceAll(number, " ", "")
	var ar []string
	var newAr []int
	sum := 0
	ar = strings.Split(check, "")
	for _, num := range ar{
		if _, err := strconv.Atoi(num); err != nil {
			return 0, false
		}
		if f, err := strconv.Atoi(num); err == nil {
			newAr = append(newAr, f)
		}
	}
	l := len(newAr)
	if (l % 2) == 1 {
		for i, g := range newAr {
			if (i % 2) == 1 {
				s := g *2
				if s>9{
					s -= 9
				}
				newAr[i] = s
			}
		}
	}
	if (l % 2) == 0 {
		for i, g := range newAr {
			if (i % 2) == 0 {
				s := g *2
				if s>9{
					s -= 9
				}
				newAr[i] = s
			}
		}
	}
	for _, i := range newAr {
		sum += i
	}
	if (sum % 10) != 0 {
		return sum, false
	}
	return sum, true
}
