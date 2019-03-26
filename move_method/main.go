//TODO: Source Example
//TODO: Student Excercise


//Before
stuct A {
  methodA()
  methodB()
}

(A) methodA() string {
  return "A" + methodB()
}

(A) methodB() string {
  return "B";
}

struct B {
  whatever()
}

//After
stuct A {
  methodA()
}

struct B {
  whatever()
  methodB()
}



//Excercise
stuct Excercise_A {
  overdraftCharge()
  isOverdrawn()
}

(A) methodA() string, err {
  if(isOverdrawn()) {
    return nil, fmt.Errorf("isOverdrawn")
  }
  return "Cash Money"
}

(A) isOverdrawn(accountId int) string {
  overdrawn = accountId == 3
  return overdrawn
}

struct OverdrawnService {

}
