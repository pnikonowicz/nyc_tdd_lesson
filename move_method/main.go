package main

import "fmt"

func assert(description string, expected string, actual string) {
  if expected != actual {
    panic(fmt.Sprintf("%s\nExpected:\n[%v] but was:\n[%v]", description, expected, actual))
  }
}

//Before
type BeforeA struct {
}

func (a BeforeA) funcA() string {
  return "A" + a.funcB()
}

func (BeforeA) funcB() string {
  return "B"
}

//After
type AfterA struct {
  b AfterB
}

type AfterB struct {
}

func (a AfterA) funcA() string {
  return "A" + a.b.funcB()
}

func (AfterB) funcB() string {
  return "B"
}

func main() {
  fmt.Println("Before")
  assert("a returns a + b", "AB", BeforeA{}.funcA())
  assert("b returns b", "B", BeforeA{}.funcB())

  fmt.Println("After")
  assert("a returns a + b", "AB", AfterA{}.funcA())
  assert("b returns b", "B", AfterB{}.funcB())

  fmt.Println("Excercise")
  assertExcercise := func(description string, expectedString string, expectedError error, excercise Excercise) {
    actualString, actualError := excercise.methodA()

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
type Excercise struct {
  charge int
  balance int
  balanceService BalanceService
}

func (a Excercise) methodA() (string, error) {
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
