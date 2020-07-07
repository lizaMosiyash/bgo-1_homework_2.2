package transfer

import (
	"github.com/lizaMosiyash/bgo-1_homework-2.1/pkg/card"
	"testing"
)

func TestService_Card2Card(t *testing.T) {
	cardSvc := card.NewService("TestBank")
	cardSvc.IssueCard("111111", 100)
	cardSvc.IssueCard("222222", 1000000)

	type fields struct {
		CardSvc      *card.Service
		Comission    int64
		MinComission int64
	}
	type args struct {
		from   string
		to     string
		amount int64
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantTotal int64
		wantErr   bool
	}{
		{
			name: "LowBalanceOnSource",
			fields: fields{
				CardSvc:      cardSvc,
				MinComission: 0,
			},
			args: args{
				cardSvc.Cards[1].Number,
				cardSvc.Cards[0].Number,
				1000,
			},
			wantTotal: 100000,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				CardSvc:      tt.fields.CardSvc,
				Comission:    tt.fields.Comission,
				MinComission: tt.fields.MinComission,
			}
			gotTotal, err := s.Card2Card(tt.args.from, tt.args.to, tt.args.amount)
			if (err != nil) != tt.wantErr {
				t.Errorf("Card2Card() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotTotal != tt.wantTotal {
				t.Errorf("Card2Card() gotTotal = %v, want %v", gotTotal, tt.wantTotal)
			}
		})
	}
}