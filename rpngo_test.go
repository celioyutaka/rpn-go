package main

import (
	"testing"
)

type TestRPN struct {
	description        string
	expression         string
	rpn_expression     string
	result             float64
	is_expected_result bool
}

func TestSomeExpression(t *testing.T) {
	list_expression := [...]TestRPN{
		{description: "Wrong result", expression: "5+(1+2)*4-3", rpn_expression: "5 1 2 + 4 * + 3 -", result: 14.1,
			is_expected_result: false},
		{description: "Wrong result", expression: "5+(1+2)*4-3", rpn_expression: "5 1 2 + 4 * + 3 -", result: 13.999999,
			is_expected_result: false},

		{expression: "5+(1+2)*4-3", rpn_expression: "5 1 2 + 4 * + 3 -", result: 14.0, is_expected_result: true},
		{expression: "10+5", rpn_expression: "10 5 +", result: 15.0, is_expected_result: true},
		{expression: "10+6", rpn_expression: "10 6 +", result: 16.0, is_expected_result: true},

		{expression: "7+8", rpn_expression: "7 8 +", result: 15, is_expected_result: true},
		{expression: "7+(8*3^2+4)", rpn_expression: "7 8 3 2 ^ * 4 + +", result: 83, is_expected_result: true},
		{expression: "7*(8+3)", rpn_expression: "7 8 3 + *", result: 77, is_expected_result: true},
		{expression: "9+1*8*(9*8+8*3)+99-3", rpn_expression: "9 1 8 * 9 8 * 8 3 * + * + 99 + 3 -", result: 873, is_expected_result: true},

		{expression: "2^2-3^2", rpn_expression: "2 2 ^ 3 2 ^ -", result: -5.000000, is_expected_result: true},
		{expression: "(2-3)*(2+3)", rpn_expression: "2 3 - 2 3 + *", result: -5.000000, is_expected_result: true},
		{expression: "(2+3)*2", rpn_expression: "2 3 + 2 *", result: 10.000000, is_expected_result: true},
		{expression: "2^2+2*2*3+3^2", rpn_expression: "2 2 ^ 2 2 * 3 * + 3 2 ^ +", result: 25.000000, is_expected_result: true},
		{expression: "2^2+3^2", rpn_expression: "2 2 ^ 3 2 ^ +", result: 13.000000, is_expected_result: true},
		{expression: "(2+3)*2-2*2*3", rpn_expression: "2 3 + 2 * 2 2 * 3 * -", result: -2.000000, is_expected_result: true},
		{expression: "(2-3)*2", rpn_expression: "2 3 - 2 *", result: -2.000000, is_expected_result: true},
		{expression: "2^2-2*2*3+3^2", rpn_expression: "2 2 ^ 2 2 * 3 * - 3 2 ^ +", result: 1.000000, is_expected_result: true},
		{expression: "(2+3+4)*2", rpn_expression: "2 3 + 4 + 2 *", result: 18.000000, is_expected_result: true},
		{expression: "2^2+3^2+4^2+2*2*3+2*3*4+2*4*2", rpn_expression: "2 2 ^ 3 2 ^ + 4 2 ^ + 2 2 * 3 * + 2 3 * 4 * + 2 4 * 2 * +", result: 81.000000, is_expected_result: true},
		{expression: "(2-3-4)*2", rpn_expression: "2 3 - 4 - 2 *", result: -10.000000, is_expected_result: true},
		{expression: "2^2+3^2+4^2-2*2*3+2*3*4-2*4*2", rpn_expression: "2 2 ^ 3 2 ^ + 4 2 ^ + 2 2 * 3 * - 2 3 * 4 * + 2 4 * 2 * -", result: 25.000000, is_expected_result: true},
		{expression: "(2+3)*3", rpn_expression: "2 3 + 3 *", result: 15.000000, is_expected_result: true},
		{expression: "2^3+3*2^2*3+3*2*3^2+3^3", rpn_expression: "2 3 ^ 3 2 2 ^ * 3 * + 3 2 * 3 2 ^ * + 3 3 ^ +", result: 125.000000, is_expected_result: true},
		{expression: "(2+3)*3", rpn_expression: "2 3 + 3 *", result: 15.000000, is_expected_result: true},
		{expression: "2^3+3^3+3*2*3*(2+3)", rpn_expression: "2 3 ^ 3 3 ^ + 3 2 * 3 * 2 3 + * +", result: 125.000000, is_expected_result: true},
		{expression: "(2-3)*3", rpn_expression: "2 3 - 3 *", result: -3.000000, is_expected_result: true},
		{expression: "2^3-3*2^2*3+3*2*3^2-3^3", rpn_expression: "2 3 ^ 3 2 2 ^ * 3 * - 3 2 * 3 2 ^ * + 3 3 ^ -", result: -1.000000, is_expected_result: true},
		{expression: "2^3-3^3-3*2*3*(2-3)", rpn_expression: "2 3 ^ 3 3 ^ - 3 2 * 3 * 2 3 - * -", result: -1.000000, is_expected_result: true},
		{expression: "2^3-3^3", rpn_expression: "2 3 ^ 3 3 ^ -", result: -19.000000, is_expected_result: true},
		{expression: "(2-3)*(2^2+2*3+3^2)", rpn_expression: "2 3 - 2 2 ^ 2 3 * + 3 2 ^ + *", result: -19.000000, is_expected_result: true},
		{expression: "2^3+3^3", rpn_expression: "2 3 ^ 3 3 ^ +", result: 35.000000, is_expected_result: true},
		{expression: "(2+3)*(2^2-2*3+3^2)", rpn_expression: "2 3 + 2 2 ^ 2 3 * - 3 2 ^ + *", result: 35.000000, is_expected_result: true},
		{expression: "(2+3)*4", rpn_expression: "2 3 + 4 *", result: 20.000000, is_expected_result: true},
		{expression: "2^4+4*2^3*3+6*2^2*3^2+4*2*3^3+3^4", rpn_expression: "2 4 ^ 4 2 3 ^ * 3 * + 6 2 2 ^ * 3 2 ^ * + 4 2 * 3 3 ^ * + 3 4 ^ +", result: 625.000000, is_expected_result: true},
		{expression: "(2-3)*4", rpn_expression: "2 3 - 4 *", result: -4.000000, is_expected_result: true},
		{expression: "2^4-4*2^3*3+6*2^2*3^2-4*2*3^3+3^4", rpn_expression: "2 4 ^ 4 2 3 ^ * 3 * - 6 2 2 ^ * 3 2 ^ * + 4 2 * 3 3 ^ * - 3 4 ^ +", result: 1.000000, is_expected_result: true},
		{expression: "2^4-3^4", rpn_expression: "2 4 ^ 3 4 ^ -", result: -65.000000, is_expected_result: true},
		{expression: "(2-3)*(2+3)*(2^2+3^2)", rpn_expression: "2 3 - 2 3 + * 2 2 ^ 3 2 ^ + *", result: -65.000000, is_expected_result: true},
		{expression: "2^5-3^5", rpn_expression: "2 5 ^ 3 5 ^ -", result: -211.000000, is_expected_result: true},
		{expression: "(2-3)*(2^4+2^3*3+2^2*3^2+2*3^3+3^4)", rpn_expression: "2 3 - 2 4 ^ 2 3 ^ 3 * + 2 2 ^ 3 2 ^ * + 2 3 3 ^ * + 3 4 ^ + *", result: -211.000000, is_expected_result: true},

		{description: "Factorial test", expression: "2!+5-3!+5", rpn_expression: "2 ! 5 + 3 ! - 5 +", result: 6, is_expected_result: true},
	}
	for i := range list_expression {
		item := list_expression[i]
		rpn := RpnGo{}
		rpn.SetDebug(false)
		rpn.CalculateExpression(item.expression)
		if rpn.rpn_expression != item.rpn_expression {
			t.Fatalf(`RPN Expression["%v"] should be ["%v"]`, rpn.rpn_expression, item.rpn_expression)
		}
		if item.is_expected_result {
			if rpn.result != item.result {
				t.Fatalf(`RPN Expression["%v"] should be ["%v"]`, rpn.result, item.result)
			}
		} else {
			if rpn.result == item.result {
				t.Fatalf(`RPN Expression["%v"] should not be ["%v"]`, rpn.result, item.result)
			}
		}

	}
}

func TestIsOperator(t *testing.T) {
	list_operator_true := [...]string{"^", "*", "/", "+", "-", "=", ")", "("}
	list_operator_false := [...]string{"@", "#", "$", "?", "&", "`", ";", ","}

	r := RpnGo{}
	for i := range list_operator_true {
		operator := list_operator_true[i]
		msg := r.IsOperator(operator)
		if msg == false {
			t.Fatalf(`IsOperator("%v") = must return true, but the return was false`, operator)
		}
	}
	for i := range list_operator_false {
		operator := list_operator_false[i]
		msg := r.IsOperator(operator)
		if msg == true {
			t.Fatalf(`IsOperator("%v") = must return false, but the return was true`, operator)
		}
	}

}
