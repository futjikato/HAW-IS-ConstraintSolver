package constraint

/**
 * Equal operator
 */
type OperationEqual struct{}

func (o *OperationEqual) validate(a interface{}, b interface{}) bool {
  return a == b
}

/**
* Not Equal operator
*/
type OperationNotEqual struct{}

func (o *OperationNotEqual) validate(a interface{}, b interface{}) bool {
  return a != b
}


/**
 * Greater ( > ) operator
 */
type OperationGreater struct{}

func (o *OperationGreater) validate(a interface{}, b interface{}) bool {
  aNum := a.(int)
  bNum := b.(int)

  return aNum > bNum
}
