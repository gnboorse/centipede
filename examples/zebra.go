package main

import (
	"fmt"
	"time"

	ctp "github.com/gnboorse/centipede"
)

//   Zebra puzzle https://en.wikipedia.org/wiki/Zebra_Puzzle

func Zebra() {

	colors := []string{"Yellow", "Blue", "Red", "Ivory", "Green"}
	nationality := []string{"Norwegian", "Ukrainian", "Englishman", "Spaniard", "Japanese"}
	drink := []string{"Water", "Tea", "Milk", "Orange juice", "Coffee"}
	smoke := []string{"Kools", "Chesterfield", "Old Gold", "Lucky Strike", "Parliament"}
	pet := []string{"Fox", "Horse", "Snails", "Dog", "Zebra"}
	categories := [][]string{colors, nationality, drink, smoke, pet}

	// initialize variables
	vars := make(ctp.Variables, 0)
	fiveDomain := ctp.IntRange(0, 5)
	constraints := make(ctp.Constraints, 0)

	// add uniqueness constraints for each category
	for _, category := range categories { 
		categoryVars := make(ctp.VariableNames, 0)
		for _, vName := range category {
			varName := ctp.VariableName(vName)
			vari := ctp.NewVariable(varName, fiveDomain)
			vars = append(vars, vari)
			categoryVars = append(categoryVars, varName)
		}
		constraints = append(constraints, ctp.AllUnique(categoryVars...)...)
	}

	// intRelConstraint checks if two int variables satisfy a binary relation
	intRelConstraint := func(var1 ctp.VariableName, var2 ctp.VariableName, rel func(int, int) bool) ctp.Constraint {
		return ctp.Constraint{Vars: ctp.VariableNames{var1, var2}, ConstraintFunction: func(variables *ctp.Variables) bool {
			if variables.Find(var1).Empty || variables.Find(var2).Empty {
				return true
			}
			v1 := variables.Find(var1).Value.(int)
			v2 := variables.Find(var2).Value.(int)
			return rel(v1, v2)
		}}
	}

	// nextToConstraint checks if two int vars differ by at most one
	nextToConstraint := func(var1 ctp.VariableName, var2 ctp.VariableName) ctp.Constraint {
		return intRelConstraint(var1, var2, func(v1, v2 int) bool { return v2 == v1+1 || v2 == v1-1 })
	}
	// offsetConstraint checks if int var1 plus offset equals var2
	offsetConstraint := func(var1 ctp.VariableName, var2 ctp.VariableName, offset int) ctp.Constraint {
		return intRelConstraint(var1, var2, func(v1, v2 int) bool { return v2 == v1+offset })
	}

	vars.SetValue("Milk", 2)
	vars.SetValue("Norwegian", 0)
	constraints = append(constraints,
		ctp.Equals("Englishman", "Red"),
		ctp.Equals("Spaniard", "Dog"),
		ctp.Equals("Coffee", "Green"),
		ctp.Equals("Ukrainian", "Tea"),
		offsetConstraint("Ivory", "Green", 1),
		ctp.Equals("Old Gold", "Snails"),
		ctp.Equals("Kools", "Yellow"),
		nextToConstraint("Chesterfield", "Fox"),
		nextToConstraint("Kools", "Horse"),
		nextToConstraint("Norwegian", "Blue"),
		ctp.Equals("Lucky Strike", "Orange juice"),
		ctp.Equals("Japanese", "Parliament"))

	// create solver
	solver := ctp.NewBackTrackingCSPSolver(vars, constraints)

	begin := time.Now()
	// simplify variable domains following initial assignment
	solver.State.MakeArcConsistent()
	success := solver.Solve() // run the solution
	elapsed := time.Since(begin)

	// output results and time elapsed
	if success {
		fmt.Printf("Found solution in %s\n", elapsed)
		for _, variable := range solver.State.Vars {
			fmt.Printf("Variable %v = %v\n", variable.Name, variable.Value)
		}
	} else {
		fmt.Printf("Could not find solution in %s\n", elapsed)
	}

}
