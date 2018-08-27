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

// IntegerConstraints basic example of using int constraints
func IntegerConstraints() {
	// some integer variables
	vars := centipede.Variables{
		centipede.NewVariable("A", centipede.IntRange(1, 10)),
		centipede.NewVariable("B", centipede.IntRange(1, 10)),
		centipede.NewVariable("C", centipede.IntRange(1, 10)),
		centipede.NewVariable("D", centipede.IntRange(1, 10)),
		centipede.NewVariable("E", centipede.IntRangeStep(0, 20, 2)), // even numbers < 20
	}

	// numeric constraints
	constraints := centipede.Constraints{
		// using some constraint generators
		centipede.AllUnique("A", "B", "C", "E"), // A != B != C != E
		centipede.Equals("A", "D"),              // A = D
		// here we implement a custom constraint
		centipede.Constraint{Vars: centipede.VariableNames{"A", "E"}, // E = A * 2
			ConstraintFunction: func(variables *centipede.Variables) bool {
				// here we have to use type assertion for numeric methods since
				// Variable.Value is stored as interface{}
				if variables.Find("E").Empty || variables.Find("A").Empty {
					return true
				}
				return variables.Find("E").Value.(int) == variables.Find("A").Value.(int)*2
			}},
	}

	// solve the problem
	solver := centipede.NewBackTrackingCSPSolver(vars, constraints)
	begin := time.Now()
	success := solver.Solve() // run the solution
	elapsed := time.Since(begin)

	// output results and time elapsed
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
