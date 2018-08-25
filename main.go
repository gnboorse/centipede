package main

import "fmt"

func main() {
	constraints := IntConstraints{AllUniqueInt("A", "B", "D")}
	vars := IntVariables{
		NewIntVariable("A", IntRange(1, 5)),
		NewIntVariable("B", IntRange(1, 5)),
		NewIntVariable("C", IntRange(1, 10)),
		NewIntVariable("D", IntRange(1, 10)),
	}
	fmt.Printf("Satisfied? %#v, %#v\n", constraints.AllSatisfied(vars), vars)
	vars.SetValue("A", 1)
	vars.SetValue("B", 2)
	vars.SetValue("C", 2)
	fmt.Printf("Satisfied? %#v, %#v\n", constraints.AllSatisfied(vars), vars)
}
