// Copyright 2018 Gabriel Boorse

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// 	http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"time"

	"github.com/gnboorse/centipede"
)

// MapColoringAustralia simple implementation of the map coloring problem for Australia
func MapColoringAustralia() {

	colors := centipede.Domain{"red", "green", "blue"}

	// set a variable for each of the provinces
	// domain for the variable is the range of colors
	vars := centipede.Variables{
		// each has:   <name>,      <domain>
		centipede.NewVariable("WA", colors),
		centipede.NewVariable("NT", colors),
		centipede.NewVariable("Q", colors),
		centipede.NewVariable("NSW", colors),
		centipede.NewVariable("V", colors),
		centipede.NewVariable("SA", colors),
		centipede.NewVariable("T", colors),
	}

	// bordering provinces cannot be equal.
	// See https://en.wikipedia.org/wiki/States_and_territories_of_Australia
	constraints := centipede.Constraints{
		centipede.NotEquals("WA", "NT"),
		centipede.NotEquals("WA", "SA"),
		centipede.NotEquals("NT", "SA"),
		centipede.NotEquals("NT", "Q"),
		centipede.NotEquals("Q", "SA"),
		centipede.NotEquals("Q", "NSW"),
		centipede.NotEquals("NSW", "V"),
		centipede.NotEquals("NSW", "SA"),
		centipede.NotEquals("V", "SA"),
	}

	// create the solver with a maximum depth of 500
	solver := centipede.NewCSPSolver(vars, constraints, 500)
	begin := time.Now()
	success := solver.Solve() // run the solution
	elapsed := time.Since(begin)

	if success {
		fmt.Printf("Found solution in %s\n", elapsed)
		for _, variable := range solver.State.Vars {
			// print out values for each variable
			fmt.Printf("Variable %v = %v\n", variable.Name, variable.Value)
		}
	} else {
		fmt.Printf("Could not find solution in %s\n", elapsed)
	}
}
