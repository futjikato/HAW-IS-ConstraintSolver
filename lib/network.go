package constraint

import "fmt"
import "container/list"

type Network struct {
  Constraints []*Constraint
  Variables []*Variable
  deathCount int
}

func (n *Network) AddVariable(v *Variable) {
  if n.Variables == nil {
    n.Variables = make([]*Variable, 0)
  }

  n.Variables = append(n.Variables, v)
}

func (n *Network) AddConstraint(c *Constraint) {
  if n.Constraints == nil {
    n.Constraints = make([]*Constraint, 0)
  }

  n.Constraints = append(n.Constraints, c)
}

func (n *Network) AllDifferent(nvars ...*Variable) {
  for _, v := range nvars {
    n.AddVariable(v)
  }
  for i := 0; i < len(nvars); i++ {
    for j := 0; j < len(nvars); j++ {
      if(j > i) {
        c := new(Constraint)
        c.Left = nvars[i]
        c.Right = nvars[j]
        c.Op = new(OperationNotEqual)

        n.Constraints = append(n.Constraints, c)
      }
    }
  }
}

func (n *Network) Resolve() {
  fmt.Println("Solve with", len(n.Constraints), "constraints and", len(n.Variables), "variables:")

  // inital add all constraints to stack
  stack := list.New()
  for i := 0; i < len(n.Constraints); i++ {
    stack.PushBack(n.Constraints[i])
  }
  if !n.doAC3LA(stack) {
    fmt.Println("Inital AC3LA consistency check failed.")
    return
  }

  // iterative solving
  stack = list.New()
  n.iterSolve(stack)
  fmt.Println("\nDone!\n")
}

func (n *Network) iterSolve(stack *list.List) bool {
  fmt.Println("iterSolve")
  if !n.doAC3LA(stack) {
    return false
  }

  k := n.getVariableWithoutDecision(0)
  for k != -1 {
    v := n.Variables[k]
    fmt.Println("got", v.Name, "as variable to setup next")
    availableValues := v.Values
    for _, vv := range availableValues {
      backup := n.backup()
      fmt.Println("set", v.Name, vv)
      v.Values = append(v.Values[:0], vv)
      innerStack := list.New()
      connected := n.getConstraintsWithVariable(v)
      innerStack.PushBackList(connected)
      if n.iterSolve(innerStack) {
        return true
      } else {
        fmt.Println("Restore")
        n.restore(backup)
      }
    }
    return false
  }

  return true
}

func (n *Network) doAC3LA(stack *list.List) bool {
  hasValues := true
  for stack.Len() > 0 && hasValues {
    e := stack.Front()
    stack.Remove(e)

    var c *Constraint
    c = e.Value.(*Constraint)

    if c.ArcReduce() {
      neighbors := n.getNeighbors(c)
      stack.PushBackList(neighbors)
      hasValues = c.Left.hasValues()
      fmt.Println("reduced ", c.Left.Name, "and has value", hasValues);
    }
  }

  return hasValues
}

func (n *Network) getNeighbors(c *Constraint) *list.List {
  neighbors := list.New()
  for i := 0; i < len(n.Constraints); i++ {
    if n.Constraints[i].Right == c.Left && c.Right != n.Constraints[i].Left {
      neighbors.PushBack(n.Constraints[i])
    }
  }

  return neighbors
}

func (n *Network) getConstraintsWithVariable(v *Variable) *list.List {
  connected := list.New()
  for i := 0; i < len(n.Constraints); i++ {
    if n.Constraints[i].Right == v && n.Constraints[i].Left != v {
      connected.PushBack(n.Constraints[i])
    }
  }

  return connected
}

func (n *Network) getVariableWithoutDecision(i int) int {
  for k, v := range n.Variables[i:] {
    if len(v.Values) > 1 {
      fmt.Println(v.Name, "has", v.Values)
      return k
    }
  }

  return -1
}

func (n *Network) backup() []interface{} {
  backup := make([]interface{}, len(n.Variables))
  for k, v := range n.Variables {
    copyV := make([]interface{}, len(v.Values))
    copy(copyV, v.Values)
    backup[k] = copyV
  }

  return backup
}

func (n *Network) restore(backup []interface{}) {
  for k, copyVV := range backup {
    copyV := copyVV.([]interface{})
    n.Variables[k].Values = copyV
  }
}
