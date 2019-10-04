package main

import "fmt"

func main() {
  assert := func (description string, expected int, actual int) {
    if expected != actual {
      panic(fmt.Sprintf("%s\nExpected:\n[%v] but was:\n[%v]", description, expected, actual))
    }
  }

  fmt.Println("BEFORE")
  assert("less than five", 1, Before(1))
  assert("more than five", 2, Before(6))

  fmt.Println("AFTER")
  assert("less than five", 1, After(1))
  assert("more than five", 2, After(6))

  fmt.Println("EXCERCISE")
  assert("adds one", 2, Excercise(1))
}

func Before(numberOfLateDeliveries int) int {
  moreThanFiveLateDeliveries := func(numberOfLateDeliveries int) bool {
    return numberOfLateDeliveries > 5
  }

  if moreThanFiveLateDeliveries(numberOfLateDeliveries) {
    return 2
  } else {
    return 1
  }
}

func After(numberOfLateDeliveries int) int {
  if numberOfLateDeliveries > 5 {
    return 2
  } else {
    return 1
  }
}

//Excercise: inline everything. eventually will have a simple statement
func Excercise(start int) int {
  CreateAdder := func() func(int,int) int {
    return func(a,b int) int {
      return a+b
    }
  }
  
  addToTotal := func(a,b int) int {
    adder := CreateAdder()
    return adder(a,b)
  }
  
  subtractFromTotal := func(a,b int) int{
    adder := CreateAdder()
    return adder(a, b*-1)
  }
  
  newTotal := start
  newTotal = addToTotal(newTotal, 2)
  newTotal = subtractFromTotal(newTotal, 1)
  
  return newTotal
}
