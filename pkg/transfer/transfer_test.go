package transfer

import (
	"github.com/lizaMosiyash/bgo-1_homework-2.1/pkg/card"
	"testing"
)

func TestService_Card2Card(t *testing.T) {
	cardSvc := card.NewService("Test_Bank")
	cardSvc.IssueCard("5106210", 50000)
	cardSvc.IssueCard("111111", 10000)
	cardSvc.IssueCard("222222", 5000)
	cardSvc.IssueCard("5106211", 5000)
	cardSvc.IssueCard("5106210", 5000)



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
			name:      "CardNotFound",
			fields:    fields{
				CardSvc:      cardSvc,
				Comission:    5,
				MinComission: 1000,
			},
			args:      args{
				from:   cardSvc.Cards[0].Number,
				to:     cardSvc.Cards[1].Number,
				amount: 100,
			},
			wantTotal: 11000,
			wantErr:   false,
		},
		{
			name:      "LowBalance",
			fields:    fields{
				CardSvc:      cardSvc,
				Comission:    5,
				MinComission: 1000,
			},
			args:      args{
				from:   cardSvc.Cards[2].Number,
				to:     cardSvc.Cards[3].Number,
				amount: 100,
			},
			wantTotal: 11000,
			wantErr:   true,
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

func TestService_IsValid(t *testing.T) {
	cardSvc := card.NewService("Test_Bank")
	cardSvc.IssueCard("5106219", 50000)
	type fields struct {
		CardSvc      *card.Service
		Comission    int64
		MinComission int64
	}
	type args struct {
		number string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
		want1  bool
	}{
		{
			name:   "cardNotExist",
			fields: fields{
				CardSvc:      cardSvc,
				Comission:    5,
				MinComission: 1000,
			},
			args:   args{
				number: cardSvc.Cards[0].Number,
			},
			want:   23,
			want1:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				CardSvc:      tt.fields.CardSvc,
				Comission:    tt.fields.Comission,
				MinComission: tt.fields.MinComission,
			}
			got, got1 := s.IsValid(tt.args.number)
			if got != tt.want {
				t.Errorf("IsValid() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("IsValid() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}