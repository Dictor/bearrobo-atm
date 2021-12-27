package atmviewer

import (
	"fmt"

	survey "github.com/AlecAivazis/survey/v2"
	"github.com/dictor/bearrobo-atm/atmcontroller"
)

// Below SimpleAtmViewer and three components are desired structure.
// However, those structure too complex and big for testing on console,
// so, ignore those on console viewer and handling stdin and stdout direct.
type (
	cardReader interface {
		SetCardAttachedCallback(func())
		CardNumber() []int
		CardAttached() bool
	}

	numberReader interface {
		ClearBuffer()
		Buffer() []string
	}

	textWriter interface {
		Print(string)
	}

	SimpleAtmViewer struct {
		cardReader
		numberReader
		textWriter
		callback atmcontroller.ViewerEventCallbackFunc
	}

	ConsoleAtmViewer struct {
		cardAttached bool
		cardNumber   string
		cardPin      int
	}
)

func (av *ConsoleAtmViewer) Init() {
	// if real atm viewer, hardware and software initialize logic is wrote on here.
}

func (av *ConsoleAtmViewer) SetEventCallback(cb atmcontroller.ViewerEventCallbackFunc) {
	go func() {
		for {
			survey.AskOne(&survey.Input{
				Message: "Enter your card number : ",
			}, &av.cardNumber)
			survey.AskOne(&survey.Input{
				Message: "Enter your card PIN : ",
			}, &av.cardPin)

			if av.cardNumber == "" {
				fmt.Println("[ConsoleAtmViewer] Card detached.")
				continue
			}
			fmt.Println("[ConsoleAtmViewer] Card attached.")

			ret := cb(atmcontroller.ViewerEventCardAttached, atmcontroller.ViewerEventCardParam{
				Card: atmcontroller.Card{
					Number:      av.cardNumber,
					OwnerName:   "TEST OWNER",
					ExpireYear:  2022,
					ExpireMonth: 01,
					CVC:         123,
				},
				Pin: av.cardPin,
			})

			switch ret.(type) {
			case error:
				fmt.Printf("[ConsoleAtmViewer] Error caused during process: %s\n", ret)
				continue
			case []atmcontroller.Account:
				var (
					accountSel        int
					actionSel         int
					selectionToAction []atmcontroller.ViewerActionType = []atmcontroller.ViewerActionType{
						atmcontroller.ViewerActionBalance,
						atmcontroller.ViewerActionWithdraw,
						atmcontroller.ViewerActionDeposit,
					}
				)
				accounts := ret.([]atmcontroller.Account)
				prompt := &survey.Select{
					Message: "Choose a account :",
					Options: []string{},
				}
				for _, a := range accounts {
					prompt.Options = append(prompt.Options, a.Number)
				}
				survey.AskOne(prompt, &accountSel)
				survey.AskOne(&survey.Select{
					Message: "Choose an action for account :",
					Options: []string{"balance", "withdraw", "deposit"},
				}, &actionSel)

				actionParam := atmcontroller.ViewerEventActionParam{
					Account: accounts[accountSel],
					Action:  selectionToAction[actionSel],
					Amount:  0,
				}
				if actionSel == 1 {
					survey.AskOne(&survey.Input{
						Message: "Enter amount for withdraw : ",
					}, &actionParam.Amount)
				}
				ret = cb(atmcontroller.ViewerEventActionSelected, actionParam)

				switch ret.(type) {
				case error:
					fmt.Printf("[ConsoleAtmViewer] Error caused during process: %s\n", ret)
				case int:
					fmt.Printf("[ConsoleAtmViewer] Result: %d dollar\n", ret)
				}
			}
		}
	}()
}

func (av *ConsoleAtmViewer) Panic(err error) {
	fmt.Printf("!!!PANIC:%s", err.Error())
}
