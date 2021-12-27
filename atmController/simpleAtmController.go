package atmcontroller

import "fmt"

type SimpleAtmController struct {
	viewer      AtmViewer
	model       AtmModel
	cashbin     AtmCashBin
	currentAuth *CardAuth
}

func (c *SimpleAtmController) Init(viewer AtmViewer, model AtmModel, cashbin AtmCashBin) {
	c.viewer = viewer
	c.model = model
	c.cashbin = cashbin

	c.viewer.Init()
	c.viewer.SetEventCallback(c.ViewerEventCallback)
	if err := c.model.Init(); err != nil {
		c.viewer.Panic(err)
	}
	if err := c.cashbin.Init(); err != nil {
		c.viewer.Panic(err)
	}
}

func (c *SimpleAtmController) ViewerEventCallback(eventType ViewerEventType, params interface{}) interface{} {
	var (
		assertError error = fmt.Errorf("params assert error")
		typeError   error = fmt.Errorf("event type error")
	)

	switch eventType {
	case ViewerEventCardAttached:
		/*
			In this event,
			- callback expects params is ViewerEventCardParam
			- caller expects return is error or []Account
		*/

		// assert params to expected type
		cp, ok := params.(ViewerEventCardParam)
		if !ok {
			return assertError
		}

		// validate card through model
		err, auth := c.model.CardVerify(cp.Card, cp.Pin)
		if err != nil {
			return err
		}

		// retrieve accounts
		err, accounts := c.model.CardAccount(cp.Card, auth)
		if err != nil {
			return err
		}

		return accounts

	case ViewerEventEndTransaction:
		/*
			In this event, no params and return expected
		*/
		c.currentAuth = nil

	case ViewerEventActionSelected:
		/*
			In this event,
			- callback expects params Account
			- caller expects return is error or []Account
		*/

		// assert params to expected type
		ap, ok := params.(ViewerEventActionParam)
		if !ok {
			return assertError
		}

		switch ap.Action {
		case ViewerActionBalance:
			err, bal := c.model.AccountBalance(ap.Account, *c.currentAuth)
			if err != nil {
				return err
			}
			return bal
		case ViewerActionDeposit:
			return c.model.AccountDeposit(ap.Account, *c.currentAuth, ap.Amount)
		case ViewerActionWithdraw:
			return c.model.AccountWithdraw(ap.Account, *c.currentAuth, ap.Amount)
		}
	}
	return typeError
}
