//TODO: Book Example
//TODO: Student Excercise




Before:
string printOwing(double amount, string name) {
  banner = printBanner()
  
  //print details
  banner += "name:" + _name
  banner += "amount:" + _amount
  
  return banner
}

After:

string printOwing(double amount, string name) {
  banner = printBanner()
  details = printDetails()
  
  return banner + details
}

void printDetails() {
  print("name:" + _name)
  print("amount:" + _amount)
}


Excercise:
string printOwing() {
  owingStr = ""
  var print = (str) => {
    owingStr += str
  }
  
  //banner
  print("****Banner")
  
  //calculate outstanding
  while(element in elements) {
    print("Element: " + element)
  }
  
  //print details
  print("name:" + name)
  print("amount:" + amount)
  
  return owingStr;
}
