package main

func Before(numberOfLateDeliveries int) int {
  if moreThanFiveLateDeliveries(numberOfLateDeliveries) {
    return 2
  } else {
    return 1
  }
}

func moreThanFiveLateDeliveries(numberOfLateDeliveries int) bool {
  return numberOfLateDeliveries > 5
}

func After(numberOfLateDeliveries int) int {
  if numberOfLateDeliveries > 5 {
    return 2
  } else {
    return 1
  }
}

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
