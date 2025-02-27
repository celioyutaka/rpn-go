# RPN-GO
Reverse Polish Notation or postfix notation, application to convert math expression to RPN with step-by-step and solver using Golang


## What it can do
Turn a **math expression** like:
`( 2 + 3 ) * 2 - 2 * 2 * 3`

Into **Reverse Polish Notation** or **postfix notation** 
`2 3 + 2 * 2 2 * 3 * -`

And **solve it**:
`-2`

### Step-by-step:
#### Verbose:
```
[INFO] Sanitized Expression: ((2+3)*2-2*2*3)
[INFO] Expression Stack:
[INFO] [( ( 2 + 3 ) * 2 - 2 * 2 * 3 ) )]
[INFO] STARTING STACKING PROCESS
[INFO] Final RPN Expression: 2 3 + 2 * 2 2 * 3 * -
[INFO] START CALCULATING RPN
[INFO] Calculating: 2.000000 + 3.000000 = 5.000000
[INFO] Calculating: 5.000000 * 2.000000 = 10.000000
[INFO] Calculating: 2.000000 * 2.000000 = 4.000000
[INFO] Calculating: 4.000000 * 3.000000 = 12.000000
[INFO] Calculating: 10.000000 - 12.000000 = -2.000000
[INFO] EXPRESSION: ((2+3)*2-2*2*3)
[INFO] RPN EXPRESSION: 2 3 + 2 * 2 2 * 3 * -
[INFO] RESULT: -2.000000
-2
```

#### Debug:
```
[INFO] Sanitized Expression: ((2+3)*2-2*2*3)
[INFO] Expression Stack:
[INFO] [( ( 2 + 3 ) * 2 - 2 * 2 * 3 ) )]
[INFO] STARTING STACKING PROCESS
[DEBUG] ----------------------------------------------------------
[DEBUG] # Item: `(` at position: 0
[DEBUG] ((2+3)*2-2*2*3))
[DEBUG] RPN Expression:
[DEBUG] AppendRPNOperatorItem: `(`
[DEBUG] RPN Stack:
[DEBUG] []
[DEBUG] Operator Stack: ↓
[DEBUG]
[DEBUG] |     |
[DEBUG] |  (  |
[DEBUG] |_____|
[DEBUG]
[DEBUG] ----------------------------------------------------------
[DEBUG] # Item: `(` at position: 1
[DEBUG] ((2+3)*2-2*2*3))
[DEBUG] RPN Expression:
[DEBUG] AppendRPNOperatorItem: `(`
[DEBUG] RPN Stack:
[DEBUG] []
[DEBUG] Operator Stack: ↓
[DEBUG]
[DEBUG] |     |
[DEBUG] |  (  |
[DEBUG] |  (  |
[DEBUG] |_____|
[DEBUG]
[DEBUG] ----------------------------------------------------------
[DEBUG] # Item: `2` at position: 2
[DEBUG] ((2+3)*2-2*2*3))
[DEBUG] RPN Expression:
[DEBUG] AppendRPNItem: `2`
[DEBUG] RPN Stack:
[DEBUG] [2]
[DEBUG] Operator Stack: ↓
[DEBUG]
[DEBUG] |     |
[DEBUG] |  (  |
[DEBUG] |  (  |
[DEBUG] |_____|
[DEBUG]
[DEBUG] ----------------------------------------------------------
[DEBUG] # Item: `+` at position: 3
[DEBUG] ((2+3)*2-2*2*3))
[DEBUG] RPN Expression: 2
[DEBUG] AppendRPNOperatorItem: `+`
[DEBUG] RPN Stack:
[DEBUG] [2]
[DEBUG] Operator Stack: ↓
[DEBUG]
[DEBUG] |     |
[DEBUG] |  +  |
[DEBUG] |  (  |
[DEBUG] |  (  |
[DEBUG] |_____|
[DEBUG]
[DEBUG] ----------------------------------------------------------
[DEBUG] # Item: `3` at position: 4
[DEBUG] ((2+3)*2-2*2*3))
[DEBUG] RPN Expression: 2
[DEBUG] AppendRPNItem: `3`
[DEBUG] RPN Stack:
[DEBUG] [2 3]
[DEBUG] Operator Stack: ↓
[DEBUG]
[DEBUG] |     |
[DEBUG] |  +  |
[DEBUG] |  (  |
[DEBUG] |  (  |
[DEBUG] |_____|
[DEBUG]
[DEBUG] ----------------------------------------------------------
[DEBUG] # Item: `)` at position: 5
[DEBUG] ((2+3)*2-2*2*3))
[DEBUG] RPN Expression: 2 3
[DEBUG] AppendRPNItem: `+`
[DEBUG] RPN Stack:
[DEBUG] [2 3 +]
[DEBUG] Operator Stack: ↓
[DEBUG]
[DEBUG] |     |
[DEBUG] |  +  |
[DEBUG] |  (  |
[DEBUG] |  (  |
[DEBUG] |_____|
[DEBUG]
[DEBUG] PopOperatorFromStack
[DEBUG] RPN Stack:
[DEBUG] [2 3 +]
[DEBUG]
[DEBUG]  |     |     |     |
[DEBUG]  |  +  |     |     |
[DEBUG]  |  (  |     |  (  |
[DEBUG]  |  (  |     |     |
[DEBUG]  |_____|     |_____|
[DEBUG]   BEFORE  →   AFTER
[DEBUG] PopOperatorFromStack
[DEBUG] RPN Stack:
[DEBUG] [2 3 +]
[DEBUG]
[DEBUG]  |     |     |     |
[DEBUG]  |  (  |     |     |
[DEBUG]  |  (  |     |     |
[DEBUG]  |_____|     |_____|
[DEBUG]   BEFORE  →   AFTER
[DEBUG] ----------------------------------------------------------
[DEBUG] # Item: `*` at position: 6
[DEBUG] ((2+3)*2-2*2*3))
[DEBUG] RPN Expression: 2 3 +
[DEBUG] AppendRPNOperatorItem: `*`
[DEBUG] RPN Stack:
[DEBUG] [2 3 +]
[DEBUG] Operator Stack: ↓
[DEBUG]
[DEBUG] |     |
[DEBUG] |  *  |
[DEBUG] |  (  |
[DEBUG] |_____|
[DEBUG]
[DEBUG] ----------------------------------------------------------
[DEBUG] # Item: `2` at position: 7
[DEBUG] ((2+3)*2-2*2*3))
[DEBUG] RPN Expression: 2 3 +
[DEBUG] AppendRPNItem: `2`
[DEBUG] RPN Stack:
[DEBUG] [2 3 + 2]
[DEBUG] Operator Stack: ↓
[DEBUG]
[DEBUG] |     |
[DEBUG] |  *  |
[DEBUG] |  (  |
[DEBUG] |_____|
[DEBUG]
[DEBUG] ----------------------------------------------------------
[DEBUG] # Item: `-` at position: 8
[DEBUG] ((2+3)*2-2*2*3))
[DEBUG] RPN Expression: 2 3 + 2
[DEBUG] AppendRPNItem: `*`
[DEBUG] RPN Stack:
[DEBUG] [2 3 + 2 *]
[DEBUG] Operator Stack: ↓
[DEBUG]
[DEBUG] |     |
[DEBUG] |  *  |
[DEBUG] |  (  |
[DEBUG] |_____|
[DEBUG]
[DEBUG] PopOperatorFromStack
[DEBUG] RPN Stack:
[DEBUG] [2 3 + 2 *]
[DEBUG]
[DEBUG]  |     |     |     |
[DEBUG]  |  *  |     |     |
[DEBUG]  |  (  |     |     |
[DEBUG]  |_____|     |_____|
[DEBUG]   BEFORE  →   AFTER
[DEBUG] AppendRPNOperatorItem: `-`
[DEBUG] RPN Stack:
[DEBUG] [2 3 + 2 *]
[DEBUG] Operator Stack: ↓
[DEBUG]
[DEBUG] |     |
[DEBUG] |  -  |
[DEBUG] |  (  |
[DEBUG] |_____|
[DEBUG]
[DEBUG] ----------------------------------------------------------
[DEBUG] # Item: `2` at position: 9
[DEBUG] ((2+3)*2-2*2*3))
[DEBUG] RPN Expression: 2 3 + 2 *
[DEBUG] AppendRPNItem: `2`
[DEBUG] RPN Stack:
[DEBUG] [2 3 + 2 * 2]
[DEBUG] Operator Stack: ↓
[DEBUG]
[DEBUG] |     |
[DEBUG] |  -  |
[DEBUG] |  (  |
[DEBUG] |_____|
[DEBUG]
[DEBUG] ----------------------------------------------------------
[DEBUG] # Item: `*` at position: 10
[DEBUG] ((2+3)*2-2*2*3))
[DEBUG] RPN Expression: 2 3 + 2 * 2
[DEBUG] AppendRPNOperatorItem: `*`
[DEBUG] RPN Stack:
[DEBUG] [2 3 + 2 * 2]
[DEBUG] Operator Stack: ↓
[DEBUG]
[DEBUG] |     |
[DEBUG] |  *  |
[DEBUG] |  -  |
[DEBUG] |  (  |
[DEBUG] |_____|
[DEBUG]
[DEBUG] ----------------------------------------------------------
[DEBUG] # Item: `2` at position: 11
[DEBUG] ((2+3)*2-2*2*3))
[DEBUG] RPN Expression: 2 3 + 2 * 2
[DEBUG] AppendRPNItem: `2`
[DEBUG] RPN Stack:
[DEBUG] [2 3 + 2 * 2 2]
[DEBUG] Operator Stack: ↓
[DEBUG]
[DEBUG] |     |
[DEBUG] |  *  |
[DEBUG] |  -  |
[DEBUG] |  (  |
[DEBUG] |_____|
[DEBUG]
[DEBUG] ----------------------------------------------------------
[DEBUG] # Item: `*` at position: 12
[DEBUG] ((2+3)*2-2*2*3))
[DEBUG] RPN Expression: 2 3 + 2 * 2 2
[DEBUG] AppendRPNItem: `*`
[DEBUG] RPN Stack:
[DEBUG] [2 3 + 2 * 2 2 *]
[DEBUG] Operator Stack: ↓
[DEBUG]
[DEBUG] |     |
[DEBUG] |  *  |
[DEBUG] |  -  |
[DEBUG] |  (  |
[DEBUG] |_____|
[DEBUG]
[DEBUG] PopOperatorFromStack
[DEBUG] RPN Stack:
[DEBUG] [2 3 + 2 * 2 2 *]
[DEBUG]
[DEBUG]  |     |     |     |
[DEBUG]  |  *  |     |     |
[DEBUG]  |  -  |     |  -  |
[DEBUG]  |  (  |     |     |
[DEBUG]  |_____|     |_____|
[DEBUG]   BEFORE  →   AFTER
[DEBUG] AppendRPNOperatorItem: `*`
[DEBUG] RPN Stack:
[DEBUG] [2 3 + 2 * 2 2 *]
[DEBUG] Operator Stack: ↓
[DEBUG]
[DEBUG] |     |
[DEBUG] |  *  |
[DEBUG] |  -  |
[DEBUG] |  (  |
[DEBUG] |_____|
[DEBUG]
[DEBUG] ----------------------------------------------------------
[DEBUG] # Item: `3` at position: 13
[DEBUG] ((2+3)*2-2*2*3))
[DEBUG] RPN Expression: 2 3 + 2 * 2 2 *
[DEBUG] AppendRPNItem: `3`
[DEBUG] RPN Stack:
[DEBUG] [2 3 + 2 * 2 2 * 3]
[DEBUG] Operator Stack: ↓
[DEBUG]
[DEBUG] |     |
[DEBUG] |  *  |
[DEBUG] |  -  |
[DEBUG] |  (  |
[DEBUG] |_____|
[DEBUG]
[DEBUG] ----------------------------------------------------------
[DEBUG] # Item: `)` at position: 14
[DEBUG] ((2+3)*2-2*2*3))
[DEBUG] RPN Expression: 2 3 + 2 * 2 2 * 3
[DEBUG] AppendRPNItem: `*`
[DEBUG] RPN Stack:
[DEBUG] [2 3 + 2 * 2 2 * 3 *]
[DEBUG] Operator Stack: ↓
[DEBUG]
[DEBUG] |     |
[DEBUG] |  *  |
[DEBUG] |  -  |
[DEBUG] |  (  |
[DEBUG] |_____|
[DEBUG]
[DEBUG] PopOperatorFromStack
[DEBUG] RPN Stack:
[DEBUG] [2 3 + 2 * 2 2 * 3 *]
[DEBUG]
[DEBUG]  |     |     |     |
[DEBUG]  |  *  |     |     |
[DEBUG]  |  -  |     |  -  |
[DEBUG]  |  (  |     |     |
[DEBUG]  |_____|     |_____|
[DEBUG]   BEFORE  →   AFTER
[DEBUG] AppendRPNItem: `-`
[DEBUG] RPN Stack:
[DEBUG] [2 3 + 2 * 2 2 * 3 * -]
[DEBUG] Operator Stack: ↓
[DEBUG]
[DEBUG] |     |
[DEBUG] |  -  |
[DEBUG] |  (  |
[DEBUG] |_____|
[DEBUG]
[DEBUG] PopOperatorFromStack
[DEBUG] RPN Stack:
[DEBUG] [2 3 + 2 * 2 2 * 3 * -]
[DEBUG]
[DEBUG]  |     |     |     |
[DEBUG]  |  -  |     |     |
[DEBUG]  |  (  |     |     |
[DEBUG]  |_____|     |_____|
[DEBUG]   BEFORE  →   AFTER
[DEBUG] PopOperatorFromStack
[DEBUG] RPN Stack:
[DEBUG] [2 3 + 2 * 2 2 * 3 * -]
[DEBUG]
[DEBUG]  |     |     |     |
[DEBUG]  |  (  |     |     |
[DEBUG]  |_____|     |_____|
[DEBUG]   BEFORE  →   AFTER
[DEBUG] ----------------------------------------------------------
[DEBUG] # Item: `)` at position: 15
[DEBUG] ((2+3)*2-2*2*3))
[DEBUG] RPN Expression: 2 3 + 2 * 2 2 * 3 * -
[DEBUG] AppendRPNOperatorItem: `)`
[DEBUG] RPN Stack:
[DEBUG] [2 3 + 2 * 2 2 * 3 * -]
[DEBUG] Operator Stack: ↓
[DEBUG]
[DEBUG] |     |
[DEBUG] |  )  |
[DEBUG] |_____|
[DEBUG]
[DEBUG] RPN Stack:
[DEBUG] [2 3 + 2 * 2 2 * 3 * -]
[DEBUG] Operator Stack: ↓
[DEBUG]
[DEBUG] |     |
[DEBUG] |  )  |
[DEBUG] |_____|
[DEBUG]
[DEBUG] PopOperatorFromStack
[DEBUG] RPN Stack:
[DEBUG] [2 3 + 2 * 2 2 * 3 * -]
[DEBUG]
[DEBUG]  |     |     |     |
[DEBUG]  |  )  |     |     |
[DEBUG]  |_____|     |_____|
[DEBUG]   BEFORE  →   AFTER
[INFO] Final RPN Expression: 2 3 + 2 * 2 2 * 3 * -
[INFO] START CALCULATING RPN
[DEBUG]
[DEBUG] [2 3 + 2 * 2 2 * 3 * -]
[DEBUG]
[DEBUG] # Current item: 2 isNumeric
[DEBUG] # Current item: 3 isNumeric
[DEBUG] Current item: + isOperator
[DEBUG] [2 3 + 2 * 2 2 * 3 * -]
[INFO] Calculating: 2.000000 + 3.000000 = 5.000000
[DEBUG] # Current item: 2 isNumeric
[DEBUG] Current item: * isOperator
[DEBUG] [5.000000 2 * 2 2 * 3 * -]
[INFO] Calculating: 5.000000 * 2.000000 = 10.000000
[DEBUG] # Current item: 2 isNumeric
[DEBUG] # Current item: 2 isNumeric
[DEBUG] Current item: * isOperator
[DEBUG] [10.000000 2 2 * 3 * -]
[INFO] Calculating: 2.000000 * 2.000000 = 4.000000
[DEBUG] # Current item: 4.000000 isNumeric
[DEBUG] # Current item: 3 isNumeric
[DEBUG] Current item: * isOperator
[DEBUG] [10.000000 4.000000 3 * -]
[INFO] Calculating: 4.000000 * 3.000000 = 12.000000
[DEBUG] # Current item: 12.000000 isNumeric
[DEBUG] Current item: - isOperator
[DEBUG] [10.000000 12.000000 -]
[INFO] Calculating: 10.000000 - 12.000000 = -2.000000
[INFO] EXPRESSION: ((2+3)*2-2*2*3)
[INFO] RPN EXPRESSION: 2 3 + 2 * 2 2 * 3 * -
[INFO] RESULT: -2.000000
-2
```


## How to use
### Using go run
`go run rpngo.go MATH_EXPRESSION [-v|-d]`

> -v is for verbose - show macro steps
> -d is for debug - show macro and micro steps

#### Example 1
`go run rpngo.go "10+9+9*(9+6)"`

It will return 154

#### Example 2
`go run rpngo.go "10+9+9*(9+6)" -d`

It will return all debug info and entire stacking process

## Testing
In the rpngo_test.go file there are some test cases, but it is still in progress
 
### Add some test case
Example of test function **TestSomeExpression()** on file **rpngo_test.go**

```
{description: "Test simple SUM", expression: "7+8", rpn_expression: "7 8 +", result: 15, isResultCorrect: true}
```

- **description** is the description of test
- **expression** is the math expression
- **rpn_expression** is the rpn (expected)
- **result** is the result of this expression
- **is_expected_result** is a boolean to test a wrong result (use true when **rpn_expression** and **result** are OK)

Then, just run
`go test -v`
