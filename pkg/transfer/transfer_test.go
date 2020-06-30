package transfer

import (
	"github.com/lizaMosiyash/bgo-1_homework-2.1/pkg/card"
	"testing"
)

func TestService_Card2Card(t *testing.T) {
	cardSvc := card.NewService("TestBank")
	cardSvc.IssueCard("111111", 555500)
	cardSvc.IssueCard("222222", 30000)
	println(cardSvc)

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
		wantOk    bool
	}{
		{
			name:      "myBank=>myBank,ok",
			fields:    fields{},
			args:      args{
				from:   cardSvc.Cards[0].Number,
				to:     cardSvc.Cards[1].Number,
				amount: 1000,
			},
			wantTotal: 101000,
			wantOk:    true,
		},
		{
			name:      "myBank=>MyBank,notOk",
			fields:    fields{},
			args:      args{
				from:   "222222",
				to:     "111111",
				amount: 1000,
			},
			wantTotal: 101000,
			wantOk:    false,
		},
		{
			name:      "myBank=>notMyBank,ok",
			fields:    fields{},
			args:      args{
				from:   "111111",
				to:     "333",
				amount: 1000,
			},
			wantTotal: 101000,
			wantOk:    true,
		},
		{
			name:      "myBank=>notMyBank,notOk",
			fields:    fields{},
			args:      args{
				from:   "222222",
				to:     "333",
				amount: 1000,
			},
			wantTotal: 101000,
			wantOk:    false,
		},
		{
			name:      "notMyBank=>myBank",
			fields:    fields{},
			args:      args{
				from:   "333",
				to:     "111111",
				amount: 1000,
			},
			wantTotal: 101000,
			wantOk:    true,
		},
		{
			name:      "notMyBank=>notMyBank",
			fields:    fields{},
			args:      args{
				from:   "333",
				to:     "3333",
				amount: 1000,
			},
			wantTotal: 101000,
			wantOk:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				CardSvc:      tt.fields.CardSvc,
				Comission:    tt.fields.Comission,
				MinComission: tt.fields.MinComission,
			}
			gotTotal, gotOk := s.Card2Card(tt.args.from, tt.args.to, tt.args.amount)
			if gotTotal != tt.wantTotal {
				t.Errorf("Card2Card() gotTotal = %v, want %v", gotTotal, tt.wantTotal)
			}
			if gotOk != tt.wantOk {
				t.Errorf("Card2Card() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}