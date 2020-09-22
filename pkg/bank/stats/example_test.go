package stats

import (
	"fmt"
	"bank/pkg/bank/types"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			ID:       1,
			Amount:   100,
			Category: "card",
		},
	}

	amount := Avg(payments)
	fmt.Println(amount)
	//Output:
	//100
}


func ExampleTotalInCategory() {
	payments := []types.Payment{
		{
			ID:       1,
			Amount:   100,
			Category: "card",
		},
	}

	amount := TotalInCategory(payments, "card")
	fmt.Println(amount)
	//Output:
	//100
}
