package atmcashbin

import "fmt"

type (
	SimpleAtmCashbin struct {
		balance int
	}
)

func (cb *SimpleAtmCashbin) Init() error {
	// in real cash bin, hardward initialize code is wrote here.
	cb.balance = 1000
	return nil
}

func (cb *SimpleAtmCashbin) Balance() int {
	return cb.balance
}

func (cb *SimpleAtmCashbin) RecieveCash() (error, int) {
	fmt.Println("[SimpleAtmCashbin] cashbin opened")
	fmt.Println("[SimpleAtmCashbin] enter amount of deposit dollars : ")
	amount := 0
	fmt.Scanln(&amount)
	fmt.Println("[SimpleAtmCashbin] cashbin closed")
	return nil, amount
}

func (cb *SimpleAtmCashbin) EmitCash(amount int) error {
	fmt.Println("[SimpleAtmCashbin] cashbin opened")
	fmt.Printf("[SimpleAtmCashbin] enter after get %d dollars\n", amount)
	fmt.Scanln()
	fmt.Println("[SimpleAtmCashbin] cashbin closed")
	return nil
}
