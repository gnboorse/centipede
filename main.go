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
	vars := Variables{
		// each has:   <name>,      <domain>
		NewVariable("WA", Range(1, len(colors)+1)),
		NewVariable("NT", Range(1, len(colors)+1)),
		NewVariable("Q", Range(1, len(colors)+1)),
		NewVariable("NSW", Range(1, len(colors)+1)),
		NewVariable("V", Range(1, len(colors)+1)),
		NewVariable("SA", Range(1, len(colors)+1)),
		NewVariable("T", Range(1, len(colors)+1)),
	}

	// bordering provinces cannot be equal.
	// See https://en.wikipedia.org/wiki/States_and_territories_of_Australia
	constraints := Constraints{
		NotEquals("WA", "NT"),
		NotEquals("WA", "SA"),
		NotEquals("NT", "SA"),
		NotEquals("NT", "Q"),
		NotEquals("Q", "SA"),
		NotEquals("Q", "NSW"),
		NotEquals("NSW", "V"),
		NotEquals("NSW", "SA"),
		NotEquals("V", "SA"),
	}

	// create the solver with a maximum depth of 500
	solver := NewCSPSolver(vars, constraints, 500)
	begin := time.Now()
	success := solver.Solve() // run the solution
	elapsed := time.Since(begin)

	if success {
		fmt.Printf("Found solution in %s\n", elapsed)
		for _, variable := range solver.State.Vars {
			// print out values for each variable
			fmt.Printf("Variable %v = %v\n", variable.Name, colors[variable.Value.(int)-1])
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
