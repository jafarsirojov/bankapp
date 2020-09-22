package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExamplePaymentSources() {
	cards := []types.Card{
		{
			PAN:     "5058 xxxx xxxx 8888",
			Balance: 10,
			Active:  true,
		}, {
			PAN:     "5058 xxxx xxxx 8888",
			Balance: 10,
			Active:  false,
		}, {
			PAN:     "5058 xxxx xxxx 8888",
			Balance: 0,
			Active:  true,
		},
	}

	payments := PaymentSources(cards)
	for _, payment := range payments {
		fmt.Println(payment.Number)
	}
	//Output:
	//5058 xxxx xxxx 8888
	//5058 xxxx xxxx 8888
}
