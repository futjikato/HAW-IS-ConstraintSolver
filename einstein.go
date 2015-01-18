package main

import "github.com/futjikato/Constraints/lib"
import "fmt"

func main() {
  // constraint network
  n := new(constraint.Network)

  // jede nationalität
  brite := new(constraint.Variable)
  brite.Name = "Brite"
  brite.Values = []interface{}{1, 2, 3, 4, 5}

  schwede := new(constraint.Variable)
  schwede.Name = "Schwede"
  schwede.Values = []interface{}{1, 2, 3, 4, 5}

  daene := new(constraint.Variable)
  daene.Name = "Däne"
  daene.Values = []interface{}{1, 2, 3, 4, 5}

  norweger := new(constraint.Variable)
  norweger.Name = "Norweger"
  norweger.Values = []interface{}{1, 2, 3, 4, 5}

  deutscher := new(constraint.Variable)
  deutscher.Name = "Deutscher"
  deutscher.Values = []interface{}{1, 2, 3, 4, 5}

  // jede nationalität darf nur einmal vorhanden sein
  n.AllDifferent(brite, schwede, daene, norweger, deutscher)

  // hausfarben
  blau := new(constraint.Variable)
  blau.Name = "Blau"
  blau.Values = []interface{}{1, 2, 3, 4, 5}

  rot := new(constraint.Variable)
  rot.Name = "rot"
  rot.Values = []interface{}{1, 2, 3, 4, 5}

  gruen := new(constraint.Variable)
  gruen.Name = "Grün"
  gruen.Values = []interface{}{1, 2, 3, 4, 5}

  weiss := new(constraint.Variable)
  weiss.Name = "Weiß"
  weiss.Values = []interface{}{1, 2, 3, 4, 5}

  gelb := new(constraint.Variable)
  gelb.Name = "Gelb"
  gelb.Values = []interface{}{1, 2, 3, 4, 5}

  n.AllDifferent(blau, rot, gruen, weiss, gelb)

  // Wer hält was
  vogel := new(constraint.Variable)
  vogel.Name = "Vogel"
  vogel.Values = []interface{}{1, 2, 3, 4, 5}

  hund := new(constraint.Variable)
  hund.Name = "Hund"
  hund.Values = []interface{}{1, 2, 3, 4, 5}

  katze := new(constraint.Variable)
  katze.Name = "Katze"
  katze.Values = []interface{}{1, 2, 3, 4, 5}

  pferd := new(constraint.Variable)
  pferd.Name = "Pferd"
  pferd.Values = []interface{}{1, 2, 3, 4, 5}

  fisch := new(constraint.Variable)
  fisch.Name = "Fisch"
  fisch.Values = []interface{}{1, 2, 3, 4, 5}

  n.AllDifferent(vogel, hund, katze, pferd, fisch)

  // Wer trinkt was
  tee := new(constraint.Variable)
  tee.Name = "Tee"
  tee.Values = []interface{}{1, 2, 3, 4, 5}

  kaffee := new(constraint.Variable)
  kaffee.Name = "Kaffee"
  kaffee.Values = []interface{}{1, 2, 3, 4, 5}

  milch := new(constraint.Variable)
  milch.Name = "Milch"
  milch.Values = []interface{}{1, 2, 3, 4, 5}

  bier := new(constraint.Variable)
  bier.Name = "Bier"
  bier.Values = []interface{}{1, 2, 3, 4, 5}

  wasser := new(constraint.Variable)
  wasser.Name = "Wasser"
  wasser.Values = []interface{}{1, 2, 3, 4, 5}

  n.AllDifferent(tee, kaffee, milch, bier, wasser)

  // Wer raucht was
  pallmall := new(constraint.Variable)
  pallmall.Name = "Pall Mall"
  pallmall.Values = []interface{}{1, 2, 3, 4, 5}

  dunhill := new(constraint.Variable)
  dunhill.Name = "Dunhill"
  dunhill.Values = []interface{}{1, 2, 3, 4, 5}

  malboro := new(constraint.Variable)
  malboro.Name = "Malboro"
  malboro.Values = []interface{}{1, 2, 3, 4, 5}

  rothmanns := new(constraint.Variable)
  rothmanns.Name = "Rothmanns"
  rothmanns.Values = []interface{}{1, 2, 3, 4, 5}

  winfield := new(constraint.Variable)
  winfield.Name = "Winfield"
  winfield.Values = []interface{}{1, 2, 3, 4, 5}

  n.AllDifferent(pallmall, dunhill, malboro, rothmanns, winfield)

  // spezwifische positionen

  mittleresHaus := new(constraint.Variable)
  mittleresHaus.Name = "Mittleres Haus"
  mittleresHaus.Values = []interface{}{3}

  erstesHaus := new(constraint.Variable)
  erstesHaus.Name = "erstes Haus"
  erstesHaus.Values = []interface{}{1}

  n.AddVariable(mittleresHaus)
  n.AddVariable(erstesHaus)

  // rules

  // Der Brite lebt im roten Haus.
  rule1 := new(constraint.Constraint)
  rule1.Left = brite
  rule1.Right = rot
  rule1.Op = new(constraint.OperationEqual)
  n.AddConstraint(rule1)

  // Der Schwede hält sich einen Hund.
  rule2 := new(constraint.Constraint)
  rule2.Left = schwede
  rule2.Right = hund
  rule2.Op = new(constraint.OperationEqual)
  n.AddConstraint(rule2)

  // Der Däne trinkt gern Tee.
  rule3 := new(constraint.Constraint)
  rule3.Left = daene
  rule3.Right = tee
  rule3.Op = new(constraint.OperationEqual)
  n.AddConstraint(rule3)

  // Das grüne Haus steht links neben dem weißen Haus.
  rule4 := new(constraint.Constraint)
  rule4.Left = gruen
  rule4.Right = weiss
  rule4.Op = new(constraint.OperationLinksNeben)
  n.AddConstraint(rule4)

  // Der Besitzer des grünen Hauses trinkt Kaffee.
  rule5 := new(constraint.Constraint)
  rule5.Left = gruen
  rule5.Right = kaffee
  rule5.Op = new(constraint.OperationEqual)
  n.AddConstraint(rule5)

  // Die Person, die Pall Mall raucht, hat einen Vogel.
  rule6 := new(constraint.Constraint)
  rule6.Left = pallmall
  rule6.Right = vogel
  rule6.Op = new(constraint.OperationEqual)
  n.AddConstraint(rule6)

  // Der Mann im mittleren Haus trinkt Milch.
  rule7 := new(constraint.Constraint)
  rule7.Left = mittleresHaus
  rule7.Right = milch
  rule7.Op = new(constraint.OperationEqual)
  n.AddConstraint(rule7)

  // Der Bewohner des gelben Hauses raucht Dunhill.
  rule8 := new(constraint.Constraint)
  rule8.Left = gelb
  rule8.Right = dunhill
  rule8.Op = new(constraint.OperationEqual)
  n.AddConstraint(rule8)

  // Der Norweger lebt im ersten Haus.
  rule9 := new(constraint.Constraint)
  rule9.Left = norweger
  rule9.Right = erstesHaus
  rule9.Op = new(constraint.OperationEqual)
  n.AddConstraint(rule9)

  // Der Malboro-Raucher wohnt neben der Person mit der Katze.
  rule10 := new(constraint.Constraint)
  rule10.Left = malboro
  rule10.Right = katze
  rule10.Op = new(constraint.OperationNachbar)
  n.AddConstraint(rule10)

  // Der Mann mit dem Pferd lebt neben der Person, die Dunhill raucht.
  rule11 := new(constraint.Constraint)
  rule11.Left = pferd
  rule11.Right = dunhill
  rule11.Op = new(constraint.OperationNachbar)
  n.AddConstraint(rule11)

  // Der Winfield-Raucher trinkt gern Bier.
  rule12 := new(constraint.Constraint)
  rule12.Left = winfield
  rule12.Right = bier
  rule12.Op = new(constraint.OperationEqual)
  n.AddConstraint(rule12)

  // Der Norweger wohnt neben dem blauen Haus.
  rule13 := new(constraint.Constraint)
  rule13.Left = norweger
  rule13.Right = blau
  rule13.Op = new(constraint.OperationNachbar)
  n.AddConstraint(rule13)

  // Der Deutsche raucht Rothmanns.
  rule14 := new(constraint.Constraint)
  rule14.Left = deutscher
  rule14.Right = rothmanns
  rule14.Op = new(constraint.OperationEqual)
  n.AddConstraint(rule14)

  // Der Malboro-Raucher hat einen Nachbarn, der Wasser trinkt.
  rule15 := new(constraint.Constraint)
  rule15.Left = malboro
  rule15.Right = wasser
  rule15.Op = new(constraint.OperationNachbar)
  n.AddConstraint(rule15)


  n.Resolve()

  fmt.Println("Brite", brite.Values)
  fmt.Println("Schwede", schwede.Values)
  fmt.Println("Däne", daene.Values)
  fmt.Println("Norweger", norweger.Values)
  fmt.Println("Deutscher", deutscher.Values)

  fmt.Println("\nFisch", fisch.Values)
  fmt.Println("Vogel", vogel.Values)
  fmt.Println("Hund", hund.Values)
  fmt.Println("Katze", katze.Values)
  fmt.Println("Pferd", pferd.Values)

  fmt.Println("\nBlau", blau.Values)
  fmt.Println("Rot", rot.Values)
  fmt.Println("Grün", gruen.Values)
  fmt.Println("Weiß", weiss.Values)
  fmt.Println("Gelb", gelb.Values)

  fmt.Println("\nTee", tee.Values)
  fmt.Println("Kaffee", kaffee.Values)
  fmt.Println("Milch", milch.Values)
  fmt.Println("Bier", bier.Values)
  fmt.Println("Wasser", wasser.Values)

  fmt.Println("\nPall Mall", pallmall.Values)
  fmt.Println("Dunhill", dunhill.Values)
  fmt.Println("Malboro", malboro.Values)
  fmt.Println("Rothmanns", rothmanns.Values)
  fmt.Println("Winfield", winfield.Values)
}
