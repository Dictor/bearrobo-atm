# bearrobo-atm
`Implement a simple ATM controller` assignment repository

## Instruction
### Requirements
- git
- [golang 1.17](https://go.dev/dl/)
- Internet connection (for downloading dependencies)

### Procedure
1. clone this project
```
git clone https://github.com/Dictor/bearrobo-atm
```
2. move to project directory and build
```
cd bearrobo-atm
go build
```
3. run executive
```
./bearrobo-atm
```
4. test dataset has assigned as below. When program turned on, 
you can input card number and PIN. If you give correct input then you can select accounts and action.
- 1 card : 1234123412341234 (PIN is 1234)
- 2 accounts on above card : 123123123123, 345345345345
- account 123123123123 has 100 dollars and 345345345345 has 50 dollars on intial

## Structure
There are four major components which is that `atmController`, `atmViewer`, `atmCashbin`, `atmModel` in this project.
Every component's specifications are defined as `interface` on `atmcontroller/definition.go`. 
Engineer can implement another three components as variable ways, just need to satisfy interface. 
Controller got their components through dependency injection on initialize process.

`atmViewer` represents atm frontend like keypad, screen, speaker. 
This fires user input event (like card attached) to controller and gets input and displays output.
There are two example implementation of `atmViewer` on `atmviewer/simpleAtmViewer.go`. `SimpleAtmViewer` is ideal implementation of viewer.
This has three lower components which is that `cardReader`, `numberReader` and `textWriter`. Those are defined as interface so, Engineer can use any hardware or implementation
for atm frontend method. However, `ConsoleAtmViewer` has no lower components definition because for make it simple. 

`atmCashbin` represents atm cash bin frontend and it's similar with Viewer but seperated from Viewer because of importantation of this.
`atmModel` represents database of cards and accounts. This can be API wrapper between real bank.  
