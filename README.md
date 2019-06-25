Question: what should be done to optimize pairing rotation? Should the class decide?
https://www.ascd.org/ASCD/pdf/books/mctighe2004_intro.pdf
https://images.template.net/wp-content/uploads/2016/05/06055421/Music-Lesson-Plan-Template-Word-2010-Free-Download.jpg

# Problem
It has been observed that TDD proficiency and level of adherence could use improvement. 

# Objective
Increase TDD proficiency across the NYC office via a course. Progress measured via feedback at the end of class. Another after a 3 month period. Attendance will be voluntary. 
Students will self assess: <link for feedback>



# Requirements - Instructor
a projector or tv, internet for github, and send copies of relevant pages for reference from Martin Fowlers “Refactoring”  to students https://books.google.com/books/about/Refactoring.html?id=HmrDHwgkbPsC&printsec=frontcover&source=kp_read_button#v=onepage&q&f=false

# Requirements - Student
A computer with golang (current version). Read over materials sent by instructor. 

It is ok to use an IDE, but you are not allowed to use the refactoring plugins (even rename!)

# Test Coverage
Ask these questions to the group. Open discussion as there is no “correct” answer.
Try to cover topics of feature testing, behaviour testing, and unit testing. As well as the pros and cons of each. Lots of teams define these testing categories differently, that is the takeaway from this discussion. 


## How to measure test coverage?
Cyclomatic complexity - The sad truth is that there probably is no single simple measurement that can express an abstract concept such as complexity in a single number. But this does not imply that we cannot measure and control complexity. It just has to be done with multiple metrics and checks that cover the various aspects of complexity. - https://www.cqse.eu/en/blog/mccabe-cyclomatic-complexity/
Test coverage tools - 

For example, the following simple Java code snippet is reported with complexity 2 by the Eclipse Metrics Plugin, with 4 by GMetrics, and with complexity 5 by SonarQube:
```
int foo (int a, int b) {
        if (a > 17 && b < 42 && a+b < 55) {
                return 1;
        }
        return 2;
} - https://www.cqse.eu/en/blog/mccabe-cyclomatic-complexity/
```

## Is this code fully tested? 
Strive for 100% code coverage? - this is highly debatable
How many tests is enough? - this is highly debatable

# Refactoring

## What is refactoring? 
Code refactoring is the process of restructuring existing computer code—changing the factoring—without changing its external behavior. - https://en.wikipedia.org/wiki/Code_refactoring

## Why refactor? 
Refactoring is usually motivated by noticing a code smell.[2] For example, the method at hand may be very long, or it may be a near duplicate of another nearby method. Once recognized, such problems can be addressed by refactoring the source code, or transforming it into a new form that behaves the same as before but that no longer "smells". - https://en.wikipedia.org/wiki/Code_refactoring


## When to do it? 
   When It is easier to extend the capabilities of the application if it uses recognizable design patterns, and it provides some flexibility where none before may have existed. https://en.wikipedia.org/wiki/Code_refactoring

## When is refactoring risky?
   When there are no tests

Talk about testing frameworks. How many have everyone used. What are common functions of testing frameworks (ie assertEquals, assertTrue, assertThat(x, Is(anything())) etc..)


# First test framework
## What do we really need in a testing framework?
  A function, maybe a print statement

Have students write Assert for equality
Explain this is probably all we will need, but feel free to develop the tools that will help you go faster ™ 
Try not to be strict here. 

Rotate pairs.





# Plan
Introduce technique by summarize from book. The book is quite detailed, put it into your own words. 
Demonstrate with an example in front of class. 
Give class their own example to try with their pair. Rotate pairs after every technique.
Have students demonstrate technique in front of class?

Refactoring legacy systems techniques:
Show code without tests. 
Write tests
Comment code and ensure tests fail
Refactor using the technique. 


Techniques: (4 hours) 
Extract variable (do this one first because it is very easy)
Inline Method
Extract method
Extract struct
Move method
Extract Interface
Add parameter
Remove parameter
(Additional examples from book)

Using techniques (2-4 hours)
Refactor gilded rose (change pairs every X minutes) https://github.com/pnikonowicz/gildedrose
 






# Techniques

Extract Variable: (p 124)
https://github.com/pnikonowicz/nyc_tdd_lesson/blob/master/extract_variable.go


Extract Method: (p 110)
https://github.com/pnikonowicz/nyc_tdd_lesson/blob/master/extract_method.go

Inline Function: (p 117)
https://github.com/pnikonowicz/nyc_tdd_lesson/blob/master/inline_method.go

In Student excercise, pay attention to subtractFromTotal. Point out that it is coupled to AddToTotal and cannot be moved withouth moving add to total. 


Move Function: (p 117)
https://github.com/pnikonowicz/nyc_tdd_lesson/blob/master/move_method.go

In Student excercise, Talk about how there might be feature envy, how the overdraft function belongs in the service. 







# Email Template:

Hello and thank you for your interest in attending NYC working in legacy systems TDD workshop. This course will help prepare you for working in legacy code that has no tests. The format will basically be, write tests, ensure tests pass, then refactor. The refactor part is what we will be focusing on. We will be using the refactoring methods outlined in Martin Fowler’s “Refactoring” book. We will be learning very small refactoring steps. 

In order to get the most out of this course, please do not use any IDE refactoring tools. 

There are a couple things that you should check out before showing up.

Read the Introduce Explaining Variable from Martin Fowler’s “Refactoring” links here:
https://github.com/pnikonowicz/nyc_tdd_lesson/blob/master/extract_variable/extract_variable_1.JPG
https://github.com/pnikonowicz/nyc_tdd_lesson/blob/master/extract_variable/extract_variable_2.JPG

Make sure your laptop has Go installed.The version will not matter for what we are going to be doing.

You do not need an IDE, but you can use one. (no using refactoring tools please)

Generally, the structure for this course will be:

Intro general questions and warming up (30 min) 

I will explain the refactoring technique. (15-30 min) 

We will break off into pairs (5 min) 

With your pair, you will create your own custom test framework and begin to write tests for the example function I will provide you.(30 min) 

Rotate Pairs (5 min) Refactor / continue to write tests if needed (15 min)

Feedback/Retrospective (15 min)

Looking forward to seeing you!

