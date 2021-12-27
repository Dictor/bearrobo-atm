package main

import (
	"runtime"

	"github.com/dictor/bearrobo-atm/atmcashbin"
	"github.com/dictor/bearrobo-atm/atmcontroller"
	"github.com/dictor/bearrobo-atm/atmmodel"
	"github.com/dictor/bearrobo-atm/atmviewer"
)

func main() {
	var (
		controller atmcontroller.SimpleAtmController = atmcontroller.SimpleAtmController{}
		viewer     atmcontroller.AtmViewer           = &atmviewer.ConsoleAtmViewer{}
		model      atmcontroller.AtmModel            = &atmmodel.SimpleAtmModel{}
		cashbin    atmcontroller.AtmCashBin          = &atmcashbin.SimpleAtmCashbin{}
	)
	controller.Init(viewer, model, cashbin)

	for {
		runtime.Gosched()
	}
}
