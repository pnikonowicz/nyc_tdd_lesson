//TODO: Book Example
//TODO: Student Example




Before:
void printOwing(double amount) {
  printBanner()
  
  //print details
  print("name:" + _name)
  print("amount:" + _amount)
}

After:

void printOwing(double amount) {
  printBanner()
  printDetails()
  
}

void printDetails() {
  print("name:" + _name)
  print("amount:" + _amount)
}
