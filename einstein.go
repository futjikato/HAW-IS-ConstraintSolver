package main

import "github.com/futjikato/Constraints/lib"
import "fmt"

func main() {
  // constraint network
  n := new(constraint.Network)

  // jede nationalität
  brite := new(constraint.Variable)
  brite.Values = []interface{}{1, 2, 3, 4, 5}

  schwede := new(constraint.Variable)
  schwede.Values = []interface{}{1, 2, 3, 4, 5}

  daene := new(constraint.Variable)
  daene.Values = []interface{}{1, 2, 3, 4, 5}

  norweger := new(constraint.Variable)
  norweger.Values = []interface{}{1, 2, 3, 4, 5}

  deutscher := new(constraint.Variable)
  deutscher.Values = []interface{}{1, 2, 3, 4, 5}

  // jede nationalität darf nur einmal vorhanden sein
  n.AllDifferent(brite, schwede, daene, norweger, deutscher)

  // hausfarben
  blau := new(constraint.Variable)
  blau.Values = []interface{}{1, 2, 3, 4, 5}

  rot := new(constraint.Variable)
  rot.Values = []interface{}{1, 2, 3, 4, 5}

  gruen := new(constraint.Variable)
  gruen.Values = []interface{}{1, 2, 3, 4, 5}

  weiss := new(constraint.Variable)
  weiss.Values = []interface{}{1, 2, 3, 4, 5}

  gelb := new(constraint.Variable)
  gelb.Values = []interface{}{1, 2, 3, 4, 5}

  n.AllDifferent(blau, rot, gruen, weiss, gelb)

  // rules
  briterot := new(constraint.Constraint)
  briterot.Left = brite
  briterot.Right = rot
  briterot.Op = new(constraint.OperationEqual)
  n.AddConstraint(briterot)

  n.Resolve()

  fmt.Println(brite.Values)
  fmt.Println(rot.Values)
  fmt.Println(schwede.Values)
  fmt.Println(daene.Values)
  fmt.Println(norweger.Values)
  fmt.Println(deutscher.Values)
}
