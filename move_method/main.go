//Before
type A struct {
}

func (a A) funcA() string {
  return "A" + a.funcB()
}

func (A) funcB() string {
  return "B"
}

struct B {
}

//After
stuct A {
}

struct B {
}

func (a A) funcA() string {
  b := B{}
  return "A" + b.funcB()
}

func (B) funcB() string {
  return "B"
}


//Excercise
struct Excercise {
  charge int
  balance int
}

func (a A) methodA() string, err {
  if(a.isOverdrawn()) {
    return nil, fmt.Errorf("isOverdrawn")
  }
  return "Cash Money", nil
}

func (a A) isOverdrawn() bool {
  overdrawn := a.charge > a.balance
  return overdrawn
}

type OverdrawnService struct {

}
