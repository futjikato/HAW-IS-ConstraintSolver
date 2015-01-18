package constraint

import "fmt"

type Operation interface {
	validate(a interface{}, b interface{}) bool
}

type Constraint struct {
	Left   *Variable
	Right  *Variable
	Op     Operation
}

func (c *Constraint) ArcReduce() bool {
	hasChanged := false
	leftIter := c.Left.getValueIterator()
	// for each left value
	for i := 0; i < len(c.Left.Values); i++ {
		// check for matching partner in right value
		hasPartner := false
		leftVal := leftIter()
		rightIter := c.Right.getValueIterator()
		for j := 0; j < len(c.Right.Values); j++ {
			rightVal := rightIter()
			if c.Op.validate(leftVal, rightVal) {
				hasPartner = true
				break
			}
		}

		// if none is found rmove that left value
		if !hasPartner {
			fmt.Print("c")
			hasChanged = true
			c.Left.Values = append(c.Left.Values[:i], c.Left.Values[i+1:]...)
		}
	}

	return hasChanged
}

func inSlice(elem interface{}, slice []interface{}) bool {
	for _, se := range slice {
		if se == elem {
			return true
		}
	}
	return false
}
