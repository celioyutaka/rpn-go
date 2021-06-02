# RPN-GO
Reverse Polish Notation or postfix notation, application to convert math expression to RPN with debugstep-by-step and solver using Golang


## What it can do
Turn a math expression like

> ( 2 + 3 ) * 2 - 2 * 2 * 3

Into Reverse Polish Notation or postfix notation 

> 2 3 + 2 * 2 2 * 3 * -

And solve it
> -2

Debug, step-by-step how this app is converting 
> ![image](https://user-images.githubusercontent.com/12768598/116594131-451fcd80-a8f8-11eb-8ee2-567f12e5d6ec.png)
> ![image](https://user-images.githubusercontent.com/12768598/116594230-5cf75180-a8f8-11eb-8711-9299893e6be4.png)

## How to use
### Using go run
> go run rpngo.go MATH_EXPRESSION [-d]
> 
> -d is for debug
> 
#### Example 1
> go run rpngo.go "10+9+9*(9+6)"
> 
> It will return 154

#### Example 2
> go run rpngo.go "10+9+9*(9+6)" -d
> 
> It will return all debug info and entire stacking process

## Testing
> In the rpngo_test.go file there are some test cases, but it is still in progress
 
### Add some test case
Example of test function **TestSomeExpression()** on file **rpngo_test.go**

> {expression: "7+8", rpn_expression: "7 8 +", result: 15, isResultCorrect: true},
 
> **expression** is the math expression
> 
> **rpn_expression** is the rpn (expected)
> 
> **result** is the result of this expression
> 
> **isResultCorrect** is a boolean to test a wrong result (use true when **rpn_expression** and **result** are OK)

Then, just run
> go test -v
