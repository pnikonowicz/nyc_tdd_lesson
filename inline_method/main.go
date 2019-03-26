//TODO: Teacher Example
//TODO: Student Excercise

//From p. 117


//Before:
int getRating() {
  return moreThanFiveLateDeliveries() ? 2 : 1;
}

bool moreThanFiveLateDeliveries() {
  return numberOfLateDeliveries > 5;
}

//After:

int getRating() {
  return (numberOfLateDeliveries > 5) ? 2 : 1
}


//Excercise:
int getTotal(start int) {
  addToTotal = func(a,b) {
    CreateAdder = func() {
      return func(a,b) {
        return a+b
      }
    }
    adder = CreateAdder()
    adder(a,b)
  }
  
  subtractFromTotal = func(a,b) {
    return addToTotal(a, b*-1)
  }
  
  newTotal = start
  newTotal = addToTotal(newTotal, 2)
  newTotal = subtractFromTotal(newTotal, 1)
}
