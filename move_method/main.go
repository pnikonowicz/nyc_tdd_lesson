package main

import "fmt"

//Before
type BeforeA struct {
}

type BeforeB struct {
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
