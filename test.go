package main

import "constraint"
import "fmt"

func main() {
	c1 := new(constraint.Constraint)
	c2 := new(constraint.Constraint)

	var1 := new(constraint.Variable)
	var1.Values = []interface{}{1, 2, 4}

	var2 := new(constraint.Variable)
	var2.Values = []interface{}{8, 3, 4, 5}

	var3 := new(constraint.Variable)
	var3.Values = []interface{}{4}

	op1 := new(constraint.OperationGreater)
	op2 := new(constraint.OperationEqual)

	c1.Left = var1
	c1.Right = var2
	c1.Op = op1

	c2.Right = var1
	c2.Left = var3
	c2.Op = op2

	n := new(constraint.Network)
	n.AddConstraint(c1)
	n.AddConstraint(c2)

	n.Resolve()

	fmt.Println(var1.Result)
	fmt.Println(var2.Result)
	fmt.Println(var3.Result)
}
