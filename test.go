package main

import "constraint"
import "fmt"

func main() {
	c := new(constraint.Constraint)

	var1 := new(constraint.Variable)
	var1.Values = []interface{}{1, 2, 4}

	var2 := new(constraint.Variable)
	var2.Values = []interface{}{8, 3, 4, 5}

	op := new(constraint.OperationGreater)

	c.Left = var1
	c.Right = var2
	c.Op = op

	c.Resolve()
	fmt.Println(c.Left.Result)
	fmt.Println(c.Right.Result)
}
