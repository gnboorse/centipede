package main

import (
	"fmt"
	"time"
)

func main() {
	// simple implementation of the map coloring problem for Australia
	colors := [3]string{"red", "green", "blue"}

	// set a variable for each of the provinces
	// domain for the variable is the range of colors
	vars := IntVariables{
		// each has:   <name>,      <domain>
		NewIntVariable("WA", IntRange(1, len(colors)+1)),
		NewIntVariable("NT", IntRange(1, len(colors)+1)),
		NewIntVariable("Q", IntRange(1, len(colors)+1)),
		NewIntVariable("NSW", IntRange(1, len(colors)+1)),
		NewIntVariable("V", IntRange(1, len(colors)+1)),
		NewIntVariable("SA", IntRange(1, len(colors)+1)),
		NewIntVariable("T", IntRange(1, len(colors)+1)),
	}

	// bordering provinces cannot be equal.
	// See https://en.wikipedia.org/wiki/States_and_territories_of_Australia
	constraints := IntConstraints{
		NotEqualsInt("WA", "NT"),
		NotEqualsInt("WA", "SA"),
		NotEqualsInt("NT", "SA"),
		NotEqualsInt("NT", "Q"),
		NotEqualsInt("Q", "SA"),
		NotEqualsInt("Q", "NSW"),
		NotEqualsInt("NSW", "V"),
		NotEqualsInt("NSW", "SA"),
		NotEqualsInt("V", "SA"),
	}

	// create the solver with a maximum depth of 500
	solver := NewIntCSPSolver(vars, constraints, 500)
	begin := time.Now()
	success := solver.Solve() // run the solution
	elapsed := time.Since(begin)

	if success {
		fmt.Printf("Found solution in %s\n", elapsed)
		for _, variable := range solver.State.Vars {
			// print out values for each variable
			fmt.Printf("Variable %v = %v", variable.Name, colors[variable.Value-1])
		}
	} else {
		fmt.Printf("Could not find solution in %s\n", elapsed)
	}

	// expected output is:

	// Found solution in 46.207µs
	// Variable WA = red
	// Variable NT = green
	// Variable Q = red
	// Variable NSW = green
	// Variable V = red
	// Variable SA = blue
	// Variable T = red
}
