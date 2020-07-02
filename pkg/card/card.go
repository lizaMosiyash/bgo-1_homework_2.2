package card

type Card struct {
	BankName string
	Number   string
	Balance  int64
	Currency string
}

type Service struct {
	BankName string
	Cards    []*Card
}

func (s *Service) IssueCard(number string, balance int64) *Card {
	card := &Card{
		BankName: "MyBank",
		Number:   number,
		Balance:  balance,
		Currency: "RUB",
	}
	s.Cards = append(s.Cards, card)

	return card
}

func NewService(bankName string) *Service {
	return &Service{BankName: bankName}
}
