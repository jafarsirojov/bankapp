package card

import (
	"github.com/jafarsirojov/bankapp/pkg/bank/types"
)

//PaymentSources -
func PaymentSources(cards []types.Card) (payments []types.PaymentSource) {
	payment := types.PaymentSource{}
	for _, card := range cards {
		if card.Balance >= 0 && card.Active {
			payment.Type = "card"
			payment.Number = string(card.PAN)
			payment.Balance = card.Balance
			payments = append(payments, payment)
		}
	}
	return
}
