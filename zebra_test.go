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

func TestZebra(t *testing.T) {

	colors := []string{"Yellow", "Blue", "Red", "Ivory", "Green"}
	nationality := []string{"Norwegian", "Ukrainian", "Englishman", "Spaniard", "Japanese"}
	drink := []string{"Water", "Tea", "Milk", "Orange juice", "Coffee"}
	smoke := []string{"Kools", "Chesterfield", "Old Gold", "Lucky Strike", "Parliament"}
	pet := []string{"Fox", "Horse", "Snails", "Dog", "Zebra"}
	categories := [][]string{colors, nationality, drink, smoke, pet}

	// initialize variables
	vars := make(Variables[int], 0)
	fiveDomain := IntRange(0, 5)
	constraints := make(Constraints[int], 0)

	// add uniqueness constraints for each category
	for _, category := range categories {
		categoryVars := make(VariableNames, 0)
		for _, vName := range category {
			varName := VariableName(vName)
			vari := NewVariable(varName, fiveDomain)
			vars = append(vars, vari)
			categoryVars = append(categoryVars, varName)
		}
		constraints = append(constraints, AllUnique[int](categoryVars...)...)
	}

	// intRelConstraint checks if two int variables satisfy a binary relation
	intRelConstraint := func(var1 VariableName, var2 VariableName, rel func(int, int) bool) Constraint[int] {
		return Constraint[int]{Vars: VariableNames{var1, var2}, ConstraintFunction: func(variables *Variables[int]) bool {
			if variables.Find(var1).Empty || variables.Find(var2).Empty {
				return true
			}
			v1 := variables.Find(var1).Value
			v2 := variables.Find(var2).Value
			return rel(v1, v2)
		}}
	}

	// nextToConstraint checks if two int vars differ by at most one
	nextToConstraint := func(var1 VariableName, var2 VariableName) Constraint[int] {
		return intRelConstraint(var1, var2, func(v1, v2 int) bool { return v2 == v1+1 || v2 == v1-1 })
	}
	// offsetConstraint checks if int var1 plus offset equals var2
	offsetConstraint := func(var1 VariableName, var2 VariableName, offset int) Constraint[int] {
		return intRelConstraint(var1, var2, func(v1, v2 int) bool { return v2 == v1+offset })
	}

	vars.SetValue("Milk", 2)
	vars.SetValue("Norwegian", 0)
	constraints = append(constraints,
		Equals[int]("Englishman", "Red"),
		Equals[int]("Spaniard", "Dog"),
		Equals[int]("Coffee", "Green"),
		Equals[int]("Ukrainian", "Tea"),
		offsetConstraint("Ivory", "Green", 1),
		Equals[int]("Old Gold", "Snails"),
		Equals[int]("Kools", "Yellow"),
		nextToConstraint("Chesterfield", "Fox"),
		nextToConstraint("Kools", "Horse"),
		nextToConstraint("Norwegian", "Blue"),
		Equals[int]("Lucky Strike", "Orange juice"),
		Equals[int]("Japanese", "Parliament"))

	// create solver
	solver := NewBackTrackingCSPSolver(vars, constraints)

	// simplify variable domains following initial assignment
	solver.State.MakeArcConsistent(context.TODO())
	success, err := solver.Solve(context.TODO()) // run the solution
	assert.Nil(t, err)

	assert.True(t, success)

	values := map[string]int{}
	for _, variable := range solver.State.Vars {
		values[string(variable.Name)] = variable.Value
	}

	assert.Equal(t, values["Yellow"], 0)
	assert.Equal(t, values["Blue"], 1)
	assert.Equal(t, values["Red"], 2)
	assert.Equal(t, values["Ivory"], 3)
	assert.Equal(t, values["Green"], 4)
	assert.Equal(t, values["Norwegian"], 0)
	assert.Equal(t, values["Ukrainian"], 1)
	assert.Equal(t, values["Englishman"], 2)
	assert.Equal(t, values["Spaniard"], 3)
	assert.Equal(t, values["Japanese"], 4)
	assert.Equal(t, values["Water"], 0)
	assert.Equal(t, values["Tea"], 1)
	assert.Equal(t, values["Milk"], 2)
	assert.Equal(t, values["Orange juice"], 3)
	assert.Equal(t, values["Coffee"], 4)
	assert.Equal(t, values["Kools"], 0)
	assert.Equal(t, values["Chesterfield"], 1)
	assert.Equal(t, values["Old Gold"], 2)
	assert.Equal(t, values["Lucky Strike"], 3)
	assert.Equal(t, values["Parliament"], 4)
	assert.Equal(t, values["Fox"], 0)
	assert.Equal(t, values["Horse"], 1)
	assert.Equal(t, values["Snails"], 2)
	assert.Equal(t, values["Dog"], 3)
	assert.Equal(t, values["Zebra"], 4)
}
