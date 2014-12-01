package main

import "constraint"
import "fmt"

func main() {
	c1 := new(constraint.Constraint)
	c2 := new(constraint.Constraint)
	c3 := new(constraint.Constraint)

	var1 := new(constraint.Variable)
	var1.Values = []interface{}{1, 2, 4}

	var2 := new(constraint.Variable)
	var2.Values = []interface{}{8, 3, 4, 5}

	var3 := new(constraint.Variable)
	var3.Values = []interface{}{4}

	var4 := new(constraint.Variable)
	var4.Values = []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	op1 := new(constraint.OperationGreater)
	op2 := new(constraint.OperationEqual)
	op_plus := new(constraint.OperationArithPlus)

	c1.Left = var1
	c1.Right = var2
	c1.Op = op1

	c2.Right = var1
	c2.Left = var3
	c2.Op = op2

	c3.Left = var2
	c3.Right = var3
	c3.Result = var4
	c3.Op = op_plus

	n := new(constraint.Network)
	n.AddConstraint(c1)
	n.AddConstraint(c2)
	n.AddConstraint(c3)

	n.Resolve()

	fmt.Println(var1.Result)
	fmt.Println(var2.Result)
	fmt.Println(var3.Result)
	fmt.Println(var4.Result)
}
