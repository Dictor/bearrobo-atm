package atmcontroller

const (
	ViewerEventCardAttached   ViewerEventType  = "event_card_attached"
	ViewerEventActionSelected ViewerEventType  = "event_action_selected"
	ViewerEventEndTransaction ViewerEventType  = "event_end_transaction"
	ViewerActionBalance       ViewerActionType = "action_balance"
	ViewerActionDeposit       ViewerActionType = "action_deposit"
	ViewerActionWithdraw      ViewerActionType = "action_withdraw"
)

type (
	// Card is container of card information
	Card struct {
		Number      []int
		OwnerName   string
		ExpireYear  int
		ExpireMonth int
		CVC         int
	}

	// CardAuth is container of authorization infomation with card
	CardAuth struct {
		Token string
	}

	// Account is conatiner of account information
	Account struct {
		Number []int
	}

	// AtmModel is connector between this atm and bank.
	// AtmContainer transfer information with bank thorough this interface.
	AtmModel interface {
		Init() error
		AccountBalance(target Account, auth CardAuth) (error, int)
		AccountDeposit(target Account, auth CardAuth, amount int) error
		AccountWithdraw(target Account, auth CardAuth, amount int) error
		CardAccount(target Card, auth CardAuth) (error, []Account)
		CardVerify(target Card, pin int) (error, CardAuth)
	}

	ViewerEventType  string
	ViewerActionType string

	// ViewerEventCallbackFunc is callback function's signature definition
	ViewerEventCallbackFunc func(eventName ViewerEventType, params interface{}) interface{}

	ViewerEventCardParam struct {
		Card Card
		Pin  int
	}
	ViewerEventActionParam struct {
		Account Account
		Action  ViewerActionType
		Amount  int
	}

	// AtmViewer is part of atm which interact between atm and user.
	// Event callback is called when user operates this atm or some specific events caused.
	AtmViewer interface {
		Init()
		SetEventCallback(ViewerEventCallbackFunc)
		Panic(error)
	}

	// AtmCashBin is part of atm which maintain real cash and recieve and emit real cash between user.
	AtmCashBin interface {
		Init() error
		Balance() int
		RecieveCash() (error, int)
		EmitCash(int) error
	}

	// AtmController is controller of atm.
	// This control every other part of atm.
	AtmController interface {
		Init(viewer *AtmViewer, model *AtmModel, cashbin *AtmCashBin)
		ViewerEventCallback(eventName ViewerEventType, params interface{}) interface{}
	}
)
