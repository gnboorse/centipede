// Copyright 2022 Gabriel Boorse

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// 	http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package centipede

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntegerConstraints(t *testing.T) {

	// some integer variables
	vars := Variables[int]{
		NewVariable("A", IntRange(1, 10)),
		NewVariable("B", IntRange(1, 10)),
		NewVariable("C", IntRange(1, 10)),
		NewVariable("D", IntRange(1, 10)),
		NewVariable("E", IntRangeStep(0, 20, 2)), // even numbers < 20
	}

	// numeric constraints
	constraints := Constraints[int]{
		// using some constraint generators
		Equals[int]("A", "D"), // A = D
		// here we implement a custom constraint
		Constraint[int]{Vars: VariableNames{"A", "E"}, // E = A * 2
			ConstraintFunction: func(variables *Variables[int]) bool {
				// here we have to use type assertion for numeric methods since
				// Variable.Value is stored as interface{}
				if variables.Find("E").Empty || variables.Find("A").Empty {
					return true
				}
				return variables.Find("E").Value == variables.Find("A").Value*2
			}},
	}
	constraints = append(constraints, AllUnique[int]("A", "B", "C", "E")...) // A != B != C != E

	// solve the problem
	solver := NewBackTrackingCSPSolver(vars, constraints)
	success, err := solver.Solve(context.TODO()) // run the solution
	assert.Nil(t, err)

	assert.True(t, success)
	values := map[string]int{}
	for _, variable := range solver.State.Vars {
		values[string(variable.Name)] = variable.Value
	}

	assert.Equal(t, values["A"], 1)
	assert.Equal(t, values["B"], 3)
	assert.Equal(t, values["C"], 4)
	assert.Equal(t, values["D"], 1)
	assert.Equal(t, values["E"], 2)
}
