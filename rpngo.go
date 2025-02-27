package main

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	ErrorColorFont     = "\033[1;31m%v\033[0m"
	DebugColorFont     = "\033[1;32m%v\033[0m"
	WarningColorFont   = "\033[1;33m%v\033[0m"
	InfoColorFont      = "\033[1;34m%v\033[0m"
	HighlightColorFont = "\033[1;35m%v\033[0m"
)

type RpnGoer interface {
	SetDebug()
	IsDebugOn()
	printDebugHighlightAt()
	printDebug()
	SetExpression()
	GetExpression()
	GetResult()
	AppendRPNItem()
	GetRPNExpression()
	GetRPNStack()
	AppendRPNOperatorItem()
	GetLastOperatorFromStack()
	PopOperatorFromStack()
	PrintStacks()
	PrintOperatorStackBeforeAfter()
	GetOperatorStackLength()
	ConvertExpressionToStack()
	ConvertToRPN()
	CheckPrecedence()
	GetIndexOfStringList()
	IsNumericString()
	IsOperator()
	SimpleCalculate()
	CalculateRPN()
	RemoveFromStackByIndex()
	ShowVerboseResult()
	CalculateExpression()
}

type RpnGo struct {
	debug            bool
	verbose          bool
	expression       string
	expression_stack []string
	operator_stack   []string

	rpn_expression string
	rpn_stack      []string

	result        float64
	result_string string
}

func (r *RpnGo) SetDebug(debug bool) {
	r.debug = debug
}
func (r *RpnGo) IsDebugOn() bool {
	return r.debug
}
func (r *RpnGo) SetVerbose(verbose bool) {
	r.verbose = verbose
}
func (r *RpnGo) IsVerboseOn() bool {
	return r.verbose
}

func (r *RpnGo) TextColor(text string, color string) string {
	return fmt.Sprintf(color, text)
}

func (r *RpnGo) printDebugHighlightAt(list []string, index int) {
	r.printDebug("# Item: `" + fmt.Sprintf(HighlightColorFont, list[index]) + fmt.Sprintf(DebugColorFont, "` at position: "+strconv.Itoa(index)))
	aux_debug := ""
	for i := 0; i < len(list); i++ {
		if i == index {
			aux_debug = aux_debug + fmt.Sprintf(HighlightColorFont, list[i])
		} else {
			aux_debug = aux_debug + fmt.Sprintf(DebugColorFont, list[i])
		}
	}
	r.printDebug(aux_debug)
}
func (r *RpnGo) printDebug(i interface{}) {

	if r.IsDebugOn() {

		switch v := i.(type) {
		case string:
			fmt.Printf(DebugColorFont, fmt.Sprintf("[DEBUG] %s", v))
			fmt.Println()
		case int:
			fmt.Printf(DebugColorFont, fmt.Sprintf("[DEBUG] %d", v))
			fmt.Println()
		case float64:
			fmt.Printf(DebugColorFont, fmt.Sprintf("[DEBUG] %f", v))
			fmt.Println()
		case []string:
			fmt.Printf(DebugColorFont, fmt.Sprintf("[DEBUG] %v", v))
			fmt.Println()
		default:
			s, _ := json.MarshalIndent(i, "", "\t")
			fmt.Println(string(s))
		}
	}
}
func (r *RpnGo) printInfo(i interface{}) {

	if r.IsVerboseOn() || r.IsDebugOn() {
		switch v := i.(type) {
		case string:
			fmt.Printf(InfoColorFont, fmt.Sprintf("[INFO] %s", v))
			fmt.Println()
		case int:
			fmt.Printf(InfoColorFont, fmt.Sprintf("[INFO] %d", v))
			fmt.Println()
		case float64:
			fmt.Printf(InfoColorFont, fmt.Sprintf("[INFO] %f", v))
			fmt.Println()
		case []string:
			fmt.Printf(InfoColorFont, fmt.Sprintf("[INFO] %v", v))
			fmt.Println()
		default:
			s, _ := json.MarshalIndent(i, "", "\t")
			fmt.Println(string(s))
		}
	}

}

func (r *RpnGo) SetExpression(expression string) {
	expression = strings.ReplaceAll(expression, " ", "")
	expression = strings.ReplaceAll(expression, " ", "")
	expression = "(" + expression + ")"
	expression = strings.TrimSpace(expression)
	r.expression = expression

	r.printInfo("Sanitized Expression: " + r.expression)
}

func (r *RpnGo) GetExpression() string {
	return r.expression
}

func (r *RpnGo) GetResult() float64 {
	return r.result
}

func (r *RpnGo) AppendRPNItem(item string) {
	if item != "(" && item != ")" {
		r.printDebug("AppendRPNItem: `" + item + "`")
		r.rpn_expression = r.rpn_expression + item + " "
		r.rpn_stack = append(r.rpn_stack, item)
	}
	r.PrintStacks()
}

func (r *RpnGo) GetRPNExpression() string {
	return r.rpn_expression
}
func (r *RpnGo) GetRPNStack() []string {
	return r.rpn_stack
}

func (r *RpnGo) AppendRPNOperatorItem(item string) {
	r.printDebug("AppendRPNOperatorItem: `" + item + "`")
	r.operator_stack = append(r.operator_stack, item)

	r.PrintStacks()
}

func (r *RpnGo) GetLastOperatorFromStack() string {
	if len(r.operator_stack) > 0 {
		return r.operator_stack[len(r.operator_stack)-1]
	}
	return ""

}

func (r *RpnGo) PopOperatorFromStack() []string {
	r.printDebug("PopOperatorFromStack")
	aux_op_stack := r.operator_stack
	if len(aux_op_stack) > 0 {
		aux_op_stack = aux_op_stack[:len(aux_op_stack)-1]
	}
	if len(r.operator_stack) > 0 || len(aux_op_stack) > 0 {
		r.PrintOperatorStackBeforeAfter(r.operator_stack, aux_op_stack)
	}

	if len(r.operator_stack) > 0 {
		r.operator_stack = r.operator_stack[:len(r.operator_stack)-1]
	}

	return r.operator_stack
}

func (r *RpnGo) PrintStacks() {
	r.printDebug("RPN Stack: ")
	r.printDebug(r.rpn_stack)
	r.printDebug("Operator Stack: ↓")
	if len(r.operator_stack) > 0 {
		r.printDebug("     ")
		r.printDebug("|     |")
		for i := len(r.operator_stack) - 1; i >= 0; i-- {
			item := r.operator_stack[i]
			//r.printDebug(item)
			r.printDebug("|  " + item + "  |")
		}
		r.printDebug("|_____|")
		r.printDebug("       ")
	} else {
		r.printDebug("     ")
		r.printDebug("|     |")
		r.printDebug("|     |")
		r.printDebug("|_____|")
		r.printDebug("       ")
	}
}

func (r *RpnGo) PrintOperatorStackBeforeAfter(before []string, after []string) {
	r.printDebug("RPN Stack: ")
	r.printDebug(r.rpn_stack)

	var aux_debug_before []string
	var aux_debug_after []string

	max_len := max(len(before), len(after))

	if max_len == 0 {
		return
	}

	if len(before) > 0 {
		for i := max_len - 1; i >= 0; i-- {
			if i >= len(before) {
				aux_debug_before = append(aux_debug_before, "|     |")
				continue
			}
			item := before[i]
			aux_debug_before = append(aux_debug_before, " |  "+item+"  |")
		}
	}

	if len(after) > 0 {
		for i := max_len - 1; i >= 0; i-- {
			if i >= len(after) {
				aux_debug_after = append(aux_debug_after, "|     |")
				continue
			}
			item := after[i]
			aux_debug_after = append(aux_debug_after, "|  "+item+"  |")
		}
	}

	r.printDebug("                    ")
	r.printDebug(" |     |     |     |")
	for i := 0; i < max_len; i++ {
		aux_before := " |     |"
		aux_after := "|     |"

		if i < len(before) {
			aux_before = aux_debug_before[i]
		}
		if i < len(after) {
			aux_after = aux_debug_after[i]
		}
		r.printDebug(aux_before + "     " + aux_after)

	}

	if len(aux_debug_before) == 0 && len(aux_debug_after) == 0 {
		r.printDebug(" |EMPTY|     |EMPTY|")
	}

	r.printDebug(" |_____|     |_____|")
	r.printDebug("  BEFORE  →   AFTER ")
}

func (r *RpnGo) GetOperatorStackLength() int {
	return len(r.operator_stack)
}

func (r *RpnGo) ConvertExpressionToStack() []string {
	expression := r.expression
	var list []string
	tempStr := ""
	isLastCharNumeric := false

	//walk through expression, check every char
	for i := 0; i < len(expression); i++ {
		//get byte as char
		tempChar := fmt.Sprintf("%c", expression[i])

		//r.printDebug(tempChar)
		//check if char is numeric or dot "." / doest check if is "e"
		if r.IsNumericString(tempChar) {
			//if previous char is numeric OR list is empty
			if isLastCharNumeric || len(list) == 0 {
				tempStr = tempStr + tempChar
			} else {
				tempStr = tempStr + tempChar
			}
			isLastCharNumeric = true
		} else {
			if isLastCharNumeric {
				//add number to list
				list = append(list, tempStr)
			}

			tempStr = ""
			//add char to list
			list = append(list, tempChar)

			//set "previous char is numeric" as false
			isLastCharNumeric = false

		}

		//if is the last char of string
		if i == (len(expression) - 1) {
			//check if it is numeric
			if r.IsNumericString(tempChar) {
				//add number to list
				list = append(list, tempStr)
			} else {
				//add char to list
				list = append(list, tempChar)
			}

		}

	}

	/* for i := range list {
		item := list[i]
		r.printDebug(item)
	} */

	r.printInfo("Expression Stack: ")
	r.printInfo(list)

	r.expression_stack = list
	return list

}

func (r *RpnGo) ConvertToRPN() string {
	/*
		Walk through list, add number to rpn string,
			add operator to stack,
			check operator precedence
	*/

	expression_list := r.expression_stack
	//operator_precedence := []string{"=", "-", "+", "/", "*", "^"}

	first_i := true

	r.printInfo("STARTING STACKING PROCESS")

	for i := range expression_list {
		item := expression_list[i]

		r.printDebug("----------------------------------------------------------")
		r.printDebugHighlightAt(expression_list, i)
		r.printDebug("RPN Expression: " + r.rpn_expression)

		if !r.IsOperator(item) {
			//if its not operator, add item to rpn expression, then go to next item
			r.AppendRPNItem(item)
			continue
		}

		//if stack of operator is empty, just add operator to stack
		if r.GetOperatorStackLength() == 0 || first_i {
			first_i = false
			r.AppendRPNOperatorItem(item)
		} else {
			//check operator precedence, if actual operator is greater than last operator of stack, add actual operator to stack
			//check operator precedente, while actual operator is equal or  lower than last operator of stack, add last operator of stack to rpn_expression, should use while() go -> for
			//------------------------------------------------------------------------------------ add actual operator to stack
			// important, when hit left bracket, just ignore
			// important, when hit right bracket, remove everything of stack (until hit left bracked in stack), then add to rpn_expression

			//When item is "(", should add to stack, and go to next item"
			if item == "(" || item == " " {
				r.AppendRPNOperatorItem(item)
				continue
			}

			//WHEN HIT A ")", POP AS STACKS UNTIL FIND A "("
			if r.GetOperatorStackLength() > 0 && item == ")" {

				for r.GetOperatorStackLength() > 0 && r.GetLastOperatorFromStack() != "(" {
					r.AppendRPNItem(r.GetLastOperatorFromStack())

					//pop from stack
					r.PopOperatorFromStack()
					//r.GetLastOperatorFromStack()

				}
				//WHEN FIND A "(" POP IT, AND GO TO NEXT CHAR OF EXPRESSION
				if r.GetOperatorStackLength() > 0 && r.GetLastOperatorFromStack() == "(" {
					r.PopOperatorFromStack()
				}
				continue
			}

			poped_loop := false
			//check operator precedente, while actual operator is equal or  lower than last operator of stack, add last operator of stack to rpn_expression, should use while() go -> for
			for r.GetOperatorStackLength() > 0 && (r.CheckPrecedence(item) <= r.CheckPrecedence(r.GetLastOperatorFromStack())) {
				r.AppendRPNItem(r.GetLastOperatorFromStack())

				//pop from stack
				r.PopOperatorFromStack()

				poped_loop = true
			}

			if poped_loop {
				r.AppendRPNOperatorItem(item)

				poped_loop = false
			} else if r.GetOperatorStackLength() > 0 && (r.CheckPrecedence(item) > r.CheckPrecedence(r.GetLastOperatorFromStack())) {
				//check operator precedence, if actual operator is bigger than last operator of stack, add actual operator to stack
				r.AppendRPNOperatorItem(item)

			}
		}

	}

	for r.GetOperatorStackLength() > 0 {

		r.AppendRPNItem(r.GetLastOperatorFromStack())
		//pop from stack
		r.PopOperatorFromStack()

	}

	/* rpn_expression = strings.ReplaceAll(rpn_expression, "(", "")
	rpn_expression = strings.ReplaceAll(rpn_expression, ")", "")

	r.printDebug(rpn_expression) */

	//r.rpn_expression = rpn_expression

	r.printInfo("Final RPN Expression: " + r.rpn_expression)

	r.rpn_expression = strings.Trim(r.rpn_expression, " ")
	r.rpn_expression = strings.TrimRight(r.rpn_expression, " ")
	return r.rpn_expression
}

func (r *RpnGo) CheckPrecedence(item string) int {
	switch item {
	case "^":
		return 40
	case "!":
		return 40
	case "*":
		return 30
	case "/":
		return 30
	case "+":
		return 20
	case "-":
		return 20
	}
	return 0
}

func (r *RpnGo) GetIndexOfStringList(stringList []string, search string) int {
	for i := 0; i < len(stringList); i++ {
		if stringList[i] == search {
			return i
		}
	}
	return -1
}

func (r *RpnGo) IsNumericString(value string) bool {
	if value == "0" || value == "1" || value == "2" || value == "3" || value == "4" || value == "5" || value == "6" || value == "7" || value == "8" || value == "9" || value == "." {
		return true
	}
	return false
}

func (r *RpnGo) IsOperator(value string) bool {
	if value == "^" || value == "!" || value == "*" || value == "/" || value == "+" || value == "-" || value == "=" || value == ")" || value == "(" {
		return true
	}
	return false
}

func (r *RpnGo) SimpleCalculate(value1 float64, value2 float64, operator string) float64 {
	aux_result := 0.0

	switch operator {
	case "^":
		aux_result = math.Pow(value1, value2)
	case "*":
		aux_result = value1 * value2
	case "/":
		aux_result = value1 / value2
	case "+":
		aux_result = value1 + value2
	case "-":
		aux_result = value1 - value2
	case "!": // FACTORIAL
		aux_result = 1
		for i := 1; i <= int(value1); i++ {
			aux_result = aux_result * float64(i)
		}
	}

	aux_debug_string := fmt.Sprintf("Calculating: %f %s %f = %f", value1, operator, value2, aux_result)
	r.printInfo(aux_debug_string)

	return aux_result
}

func (r *RpnGo) CalculateRPN() float64 {
	r.printInfo("START CALCULATING RPN")
	r.printDebug("")
	r.printDebug(r.rpn_stack)

	r.printDebug("")
	aux_stack := r.rpn_stack
	for len(aux_stack) > 1 {
		for i := 0; i < len(aux_stack); i++ {
			item := aux_stack[i]
			if r.IsOperator(item) {
				r.printDebug("Current item: " + r.TextColor(item, HighlightColorFont) + " isOperator")
				r.printDebug(aux_stack)

				if item == "!" {
					value, err := strconv.ParseFloat(aux_stack[i-1], 64)
					if err != nil {
						fmt.Printf("Error value1 as %s", aux_stack[i-1])
					}
					result_calc := r.SimpleCalculate(value, 0, item)

					aux_stack[i] = fmt.Sprintf("%f", result_calc)
					aux_stack = r.RemoveFromStackByIndex(aux_stack, i-1)

					i = 0
					continue
				}

				//Get item of index-2, to get value1 - stack: [X Y O] -> get value of X
				value1, err := strconv.ParseFloat(aux_stack[i-2], 64)
				if err != nil {
					fmt.Printf("Error value1 as %s", aux_stack[i-2])
				}
				//Get item of index-1, to get value2 - stack: [X Y O] -> get value of Y
				value2, err := strconv.ParseFloat(aux_stack[i-1], 64)
				if err != nil {
					fmt.Printf("Error value1 as %s", aux_stack[i-1])
				}
				//calculate using value1, value2, and item (as operator)
				result_calc := r.SimpleCalculate(value1, value2, item)

				//replace item (actual index), put result from calc
				aux_stack[i] = fmt.Sprintf("%f", result_calc)
				//remove value1 and value2 from array
				aux_stack = r.RemoveFromStackByIndex(aux_stack, i-1)
				aux_stack = r.RemoveFromStackByIndex(aux_stack, i-2)

				//set i=0 to check since the start of list
				i = 0

			} else {
				//Go to next item of list
				r.printDebug("# Current item: " + item + " isNumeric")
			}
		}
	}
	if len(aux_stack) == 1 {
		aux_result, err := strconv.ParseFloat(aux_stack[0], 64)
		if err != nil {
			fmt.Printf("Error value1 as %s", aux_stack[0])
		}
		r.result_string = aux_stack[0]
		r.result = aux_result
		return r.result
	}
	return -1
}

func (r *RpnGo) CalculateExpression(expression string) float64 {
	r.SetExpression(expression)
	r.ConvertExpressionToStack()
	r.ConvertToRPN()
	r.CalculateRPN()
	return r.result
}

func (r *RpnGo) RemoveFromStackByIndex(list []string, index int) []string {
	return append(list[:index], list[index+1:]...)
}

func (r *RpnGo) ShowVerboseResult() {
	r.printInfo("EXPRESSION: " + r.expression)
	r.printInfo("RPN EXPRESSION: " + r.rpn_expression)
	r.printInfo("RESULT: " + r.result_string)
}

func (r *RpnGo) ShowResultAsTest() {
	temp_result := fmt.Sprintf(`{expression: "%v", rpn_expression: "%v", result: %f, isResultCorrect: true}`,
		r.expression, r.rpn_expression, r.result)
	r.printInfo(temp_result)
}

func main() {

	expression := ""
	arg_debug := ""

	if len(os.Args) == 1 {
		fmt.Println("Add expression to command. Ex.: rpngo 9+6")
	}
	if len(os.Args) > 1 {
		expression = os.Args[1]
	}
	if len(os.Args) > 2 {
		arg_debug = os.Args[2]
	}

	rpn := RpnGo{}
	rpn.SetDebug(false)
	if arg_debug == "-d" {
		rpn.SetDebug(true)
	}
	if arg_debug == "-v" {
		rpn.SetVerbose(true)
	}
	rpn.CalculateExpression(expression)
	rpn.ShowVerboseResult()
	fmt.Println(rpn.GetResult())

	/* for expression != "exit" {

		if expression == "" {
			fmt.Println("\n\nEnter the expression like: 7^2+9 : [type 'exit' to finish]")
			fmt.Scanf("%s\n", &expression)
		}

		if expression == "exit" {
			return
		}

		if len(expression) > 0 {

			rpn := RpnGo{}
			rpn.SetDebug(false)
			rpn.CalculateExpression(expression)
			rpn.ShowVerboseResult()
			//rpn.ShowResultAsTest()
			expression = ""
		}
	} */

}
