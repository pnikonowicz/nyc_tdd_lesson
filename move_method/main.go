package main

import "fmt"

func assert(description string, expected string, actual string) {
  if expected != actual {
    panic(fmt.Sprintf("%s\nExpected:\n[%v] but was:\n[%v]", description, expected, actual))
  }
}

//Before
type BeforeA struct {
  bString string
}

func (a *BeforeA) PrintString() string {
  a.bString="B"
  return "A" + a.printB()
}

func (a *BeforeA) printB() string {
  return a.bString
}

//After
type AfterA struct {
}

type AfterB struct {
}

func (a *AfterA) PrintString() string {
  return "A" + AfterB{}.printB("B")
}

func (AfterB) printB(s string) string {
  return s
}

func main() {
  fmt.Println("Before")
  assert("a returns a + b", "AB", (&BeforeA{}).PrintString())
  assert("b returns b", "B", (&BeforeA{bString:"B"}).printB())

  fmt.Println("After")
  assert("a returns a + b", "AB", (&AfterA{}).PrintString())
  assert("b returns b", "B", (&AfterB{}).printB("B"))

  fmt.Println("Excercise")
  assertExcercise := func(description string, expectedString string, expectedError error, excercise Excercise) {
    actualString, actualError := excercise.GetMoney()

    if expectedString != actualString {
      panic(fmt.Sprintf("%s\nExpectedString:\n[%v]\nbut was:\n[%v]", description, expectedString, actualString))
    }

    if expectedError == nil || actualError == nil {
      if expectedError != actualError {
        panic(fmt.Sprintf("%s\nExpectedError:\n[%v]\nbut was:\n[%v]", description, expectedError, actualError))
      }
    } else {
      if expectedError.Error() != actualError.Error() {
        panic(fmt.Sprintf("%s\nExpectedError:\n[%v]\nbut was:\n[%v]", description, expectedError.Error(), actualError.Error()))
      }
    }
  }
  assertExcercise("when overdrawn", "", fmt.Errorf("isOverdrawn"), Excercise{charge:3, balance: 2})
  assertExcercise("when not overdrawn", "Cash Money", nil, Excercise{charge:3, balance: 5})
}

//Excercise
//Move isOverdrawn to BalanceService
type Excercise struct {
  charge int
  balance int
  balanceService BalanceService
}

func (a Excercise) GetMoney() (string, error) {
  if a.isOverdrawn() {
    return "", fmt.Errorf("isOverdrawn")
  }
  return "Cash Money", nil
}

func (a Excercise) isOverdrawn() bool {
  overdrawn := a.charge > a.balance
  return overdrawn
}

type BalanceService struct {

}
