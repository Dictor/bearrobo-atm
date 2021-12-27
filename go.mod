module github.com/dictor/bearrobo-atm

go 1.17

require (
	github.com/dictor/bearrobo-atm/atmcashbin v0.0.0
	github.com/dictor/bearrobo-atm/atmcontroller v0.0.0
	github.com/dictor/bearrobo-atm/atmmodel v0.0.0
	github.com/dictor/bearrobo-atm/atmviewer v0.0.0
)

require (
	github.com/AlecAivazis/survey/v2 v2.3.2 // indirect
	github.com/kballard/go-shellquote v0.0.0-20180428030007-95032a82bc51 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/mgutz/ansi v0.0.0-20200706080929-d51e80ef957d // indirect
	golang.org/x/sys v0.0.0-20211216021012-1d35b9e2eb4e // indirect
	golang.org/x/term v0.0.0-20210927222741-03fcf44c2211 // indirect
	golang.org/x/text v0.3.7 // indirect
)

replace (
	github.com/dictor/bearrobo-atm/atmcashbin v0.0.0 => ./atmcashbin
	github.com/dictor/bearrobo-atm/atmcontroller v0.0.0 => ./atmcontroller
	github.com/dictor/bearrobo-atm/atmmodel v0.0.0 => ./atmmodel
	github.com/dictor/bearrobo-atm/atmviewer v0.0.0 => ./atmviewer
)
