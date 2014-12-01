package constraint

type Variable struct {
	Name   string
	Values []interface{}
	Result interface{}
}

func (v *Variable) getValueIterator() func() interface{} {
	currentValue := 0
	return func() interface{} {
		if len(v.Values) > currentValue {
			retVal := v.Values[currentValue]
			currentValue += 1
			return retVal
		}

		return nil
	}
}

type Operation interface {
	validate(a interface{}, b interface{}) bool
	reverseValidate(b interface{}, a interface{}) bool
	execute(a interface{}, b interface{}) interface{}
}

type OperationEqual struct{}

func (o *OperationEqual) validate(a interface{}, b interface{}) bool {
	return a == b
}
func (o *OperationEqual) reverseValidate(b interface{}, a interface{}) bool {
	return b == a
}
func (o *OperationEqual) execute(_ interface{}, _ interface{}) interface{} {
	return nil
}

type OperationGreater struct{}

func (o *OperationGreater) validate(a interface{}, b interface{}) bool {
	aNum := a.(int)
	bNum := b.(int)

	return aNum > bNum
}
func (o *OperationGreater) reverseValidate(b interface{}, a interface{}) bool {
	aNum := a.(int)
	bNum := b.(int)

	return bNum < aNum
}
func (o *OperationGreater) execute(_ interface{}, _ interface{}) interface{} {
	return nil
}

type OperationArithPlus struct{}

func (o *OperationArithPlus) validate(_ interface{}, _ interface{}) bool {
	return true
}
func (o *OperationArithPlus) reverseValidate(_ interface{}, _ interface{}) bool {
	return true
}
func (o *OperationArithPlus) execute(a interface{}, b interface{}) interface{} {
	aNum := a.(int)
	bNum := b.(int)
	return aNum + bNum
}

type Constraint struct {
	Left   *Variable
	Right  *Variable
	Result *Variable
	Op     Operation
}

func (c *Constraint) Check() {
	leftIter := c.Left.getValueIterator()
	okValues := make([]interface{}, 0)
	for i := 0; i < len(c.Left.Values); i++ {
		leftVal := leftIter()
		if leftVal != nil {
			rightIter := c.Right.getValueIterator()
			for j := 0; j < len(c.Right.Values); j++ {
				rightVal := rightIter()
				if rightVal != nil && c.Op.validate(leftVal, rightVal) {
					okValues = append(okValues, leftVal)
					break
				}
			}
		}
	}
	c.Left.Values = okValues

	rightIter := c.Right.getValueIterator()
	okValues = make([]interface{}, 0)
	for i := 0; i < len(c.Right.Values); i++ {
		rightVal := rightIter()
		if rightVal != nil {
			leftIter := c.Left.getValueIterator()
			for j := 0; j < len(c.Left.Values); j++ {
				leftVal := leftIter()
				if leftVal != nil && c.Op.reverseValidate(rightVal, leftVal) {
					okValues = append(okValues, rightVal)
					break
				}
			}
		}
	}
	c.Right.Values = okValues

	if c.Result != nil {
		okValues = make([]interface{}, 0)
		for i := 0; i < len(c.Left.Values); i++ {
			for j := 0; j < len(c.Right.Values); j++ {
				r := c.Op.execute(c.Left.Values[i], c.Right.Values[j])
				if inSlice(r, c.Result.Values) && r != nil {
					okValues = append(okValues, r)
				}
			}
		}
		c.Result.Values = okValues
	}

}

func (c *Constraint) Resolve() {
	leftIter := c.Left.getValueIterator()
	for i := 0; i < len(c.Left.Values); i++ {
		leftVal := leftIter()
		if leftVal != nil {
			c.Left.Result = leftVal
			break
		}
	}

	rightIter := c.Right.getValueIterator()
	for j := 0; j < len(c.Right.Values); j++ {
		rightVal := rightIter()
		if rightVal != nil && c.Op.validate(c.Left.Result, rightVal) {
			c.Right.Result = rightVal
			break
		}
	}

	if c.Result != nil {
		c.Result.Result = c.Op.execute(c.Left.Result, c.Right.Result)
	}
}

type Network struct {
	Constraints []*Constraint
}

func (n *Network) AddConstraint(c *Constraint) {
	if n.Constraints == nil {
		n.Constraints = make([]*Constraint, 0)
	}

	n.Constraints = append(n.Constraints, c)
}

func (n *Network) Resolve() {
	for i := 0; i < len(n.Constraints); i++ {
		n.Constraints[i].Check()
	}

	for i := 0; i < len(n.Constraints); i++ {
		n.Constraints[i].Resolve()
	}
}

func inSlice(elem interface{}, slice []interface{}) bool {
	for _, se := range slice {
		if se == elem {
			return true
		}
	}
	return false
}
