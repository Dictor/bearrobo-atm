package atmmodel

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/dictor/bearrobo-atm/atmcontroller"
)

type (
	SimpleAtmModel struct {
		cardPin       map[string]int    // key is card number, value is pin
		accountToCard map[string]string // key is account number, value is card number (1:n association)
		account       map[string]int    //  key is account number, int is balance
	}
)

func (am *SimpleAtmModel) Init() error {
	am.accountToCard = map[string]string{
		"123123123123": "1234123412341234",
		"345345345345": "1234123412341234",
	}
	am.cardPin = map[string]int{
		"1234123412341234": 1234,
	}
	am.account = map[string]int{
		"123123123123": 100,
		"345345345345": 50,
	}
	return nil
}

func (am *SimpleAtmModel) findAccount(target atmcontroller.Account) error {
	for id, _ := range am.account {
		if id == target.Number {
			return nil
		}
	}
	return fmt.Errorf("cannot found given account number")
}

func (am *SimpleAtmModel) findCard(target atmcontroller.Card) error {
	for num, _ := range am.cardPin {
		if num == target.Number {
			return nil
		}
	}
	return fmt.Errorf("cannot found given card number")
}

func (am *SimpleAtmModel) verifyAuth(auth atmcontroller.CardAuth) error {
	if auth.Token[0:5] == "valid" && len(auth.Token) == 21 {
		return nil
	} else {
		return fmt.Errorf("invalid card auth")
	}
}

func (am *SimpleAtmModel) AccountBalance(target atmcontroller.Account, auth atmcontroller.CardAuth) (error, int) {
	if err := am.verifyAuth(auth); err != nil {
		return err, 0
	}
	if err := am.findAccount(target); err != nil {
		return err, 0
	}
	return nil, am.account[target.Number]
}

func (am *SimpleAtmModel) AccountDeposit(target atmcontroller.Account, auth atmcontroller.CardAuth, amount int) error {
	if err := am.verifyAuth(auth); err != nil {
		return err
	}
	if err := am.findAccount(target); err != nil {
		return err
	}
	am.account[target.Number] += amount
	return nil
}

func (am *SimpleAtmModel) AccountWithdraw(target atmcontroller.Account, auth atmcontroller.CardAuth, amount int) error {
	if err := am.verifyAuth(auth); err != nil {
		return err
	}
	if err := am.findAccount(target); err != nil {
		return err
	}
	if amount > am.account[target.Number] {
		return fmt.Errorf("account balance insuffcient")
	}
	am.account[target.Number] -= amount
	return nil
}

func (am *SimpleAtmModel) CardAccount(target atmcontroller.Card, auth atmcontroller.CardAuth) (error, []atmcontroller.Account) {
	if err := am.verifyAuth(auth); err != nil {
		return err, nil
	}
	if err := am.findCard(target); err != nil {
		return err, nil
	}
	ret := []atmcontroller.Account{}
	for aid, cid := range am.accountToCard {
		if cid == target.Number {
			ret = append(ret, atmcontroller.Account{
				Number: aid,
			})
		}
	}
	return nil, ret
}

func (am *SimpleAtmModel) CardVerify(target atmcontroller.Card, pin int) (error, atmcontroller.CardAuth) {
	if err := am.findCard(target); err != nil {
		return err, atmcontroller.CardAuth{}
	}
	if am.cardPin[target.Number] == pin {
		return nil, atmcontroller.CardAuth{
			Token: fmt.Sprintf("valid%s", randomString(16)),
		}
	} else {
		return fmt.Errorf("pin not correct"), atmcontroller.CardAuth{}
	}
}

func randomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	rand.Read(b)
	return fmt.Sprintf("%x", b)[:length]
}
