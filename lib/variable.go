package constraint

type Variable struct {
  Name   string
  Values []interface{}
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

func (v *Variable) hasValues() bool {
  for _, v := range v.Values {
    if v != nil {
      return true
    }
  }
  return false
}
