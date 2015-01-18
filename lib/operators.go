package constraint

import "math"

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

type OperationLinksNeben struct{}
func (o *OperationLinksNeben) validate(a interface{}, b interface{}) bool {
  aNum := a.(int)
  bNum := b.(int)
  return aNum - 1 == bNum
}

type OperationNachbar struct{}
func (o *OperationNachbar) validate(a interface{}, b interface{}) bool {
  aNum := a.(int)
  bNum := b.(int)
  aFloat := float64(aNum)
  bFloat := float64(bNum)
  return math.Abs(aFloat - bFloat) == 1
}
