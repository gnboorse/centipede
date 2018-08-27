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

package centipede

import (
	"fmt"
)

// Constraint CSP constraint considering integer variables
type Constraint struct {
	Vars               VariableNames
	ConstraintFunction VariablesConstraintFunction
}

// Constraints collection type for Constraint
type Constraints []Constraint

// VariablesConstraintFunction function used to determine validity of Variables
type VariablesConstraintFunction func(variables *Variables) bool

// AllSatisfied check if a collection of Constraints are satisfied
func (constraints *Constraints) AllSatisfied(variables *Variables) bool {
	flag := true
	for _, constraint := range *constraints {
		flag = flag && constraint.Satisfied(variables)
	}
	return flag
}

// FilterByName return all constraints related to a particular variable name
func (constraints *Constraints) FilterByName(name VariableName) Constraints {
	filtered := make(Constraints, 0)
	for _, constraint := range *constraints {
		if constraint.Vars.Contains(name) {
			filtered = append(filtered, constraint)
		}
	}
	return filtered
}

// FilterByOrder return all constraints with the given order (number of related variables)
func (constraints *Constraints) FilterByOrder(order int) Constraints {
	filtered := make(Constraints, 0)
	for _, constraint := range *constraints {
		if len(constraint.Vars) == order {
			filtered = append(filtered, constraint)
		}
	}
	return filtered
}

// Satisfied checks to see if the given Constraint is satisfied by the variables presented
func (constraint *Constraint) Satisfied(variables *Variables) bool {
	constraintVariablesSatisfied := true
	domainSatisfied := true

	for _, varname := range constraint.Vars {
		// make sure Variables contains an object for each name in Constraint.Vars
		constraintVariablesSatisfied = constraintVariablesSatisfied && (variables.Contains(varname))
	}

	for _, variable := range *variables {
		// make sure each Variable being passed in has a value consistent with its domain or is empty
		domainSatisfied = domainSatisfied && (variable.Domain.Contains(variable.Value) || variable.Empty)
	}
	if !constraintVariablesSatisfied {
		panic(fmt.Sprintf("Insufficient variables provided. Expected %v", constraint.Vars))
	}
	if !domainSatisfied {
		panic("Variables do not satisfy the domains given.")
	}

	// now finally call the constraint function
	return constraint.ConstraintFunction(variables)
}

// Equals Constraint generator that checks if two vars are equal
func Equals(var1 VariableName, var2 VariableName) Constraint {
	return Constraint{VariableNames{var1, var2}, func(variables *Variables) bool {
		if variables.Find(var1).Empty || variables.Find(var2).Empty {
			return true
		}
		return variables.Find(var1).Value == variables.Find(var2).Value
	}}
}

// NotEquals Constraint generator that checks if two vars are not equal
func NotEquals(var1 VariableName, var2 VariableName) Constraint {
	return Constraint{VariableNames{var1, var2}, func(variables *Variables) bool {
		if variables.Find(var1).Empty || variables.Find(var2).Empty {
			return true
		}
		return variables.Find(var1).Value != variables.Find(var2).Value
	}}
}

// UnaryEquals Unary constraint that checks if var1 equals some constant
func UnaryEquals(var1 VariableName, value interface{}) Constraint {
	return Constraint{VariableNames{var1}, func(variables *Variables) bool {
		if variables.Find(var1).Empty {
			return true
		}
		return variables.Find(var1).Value == value
	}}
}

// UnaryNotEquals Unary constraint that checks if var1 is not equal to some constant
func UnaryNotEquals(var1 VariableName, value interface{}) Constraint {
	return Constraint{VariableNames{var1}, func(variables *Variables) bool {
		if variables.Find(var1).Empty {
			return true
		}
		return variables.Find(var1).Value != value
	}}
}

// // LessThan Constraint generator that checks if first variable is less than second variable
// func LessThan(var1 VariableName, var2 VariableName) Constraint {
// 	return Constraint{VariableNames{var1, var1}, func(variables Variables) bool {
// 		if variables.Find(var1).Empty || variables.Find(var2).Empty {
// 			return true
// 		}
// 		return variables.Find(var1).Value < variables.Find(var2).Value
// 	}}
// }

// // GreaterThan Constraint generator that checks if first variable is less than second variable
// func GreaterThan(var1 VariableName, var2 VariableName) Constraint {
// 	return Constraint{VariableNames{var1, var1}, func(variables Variables) bool {
// 		if variables.Find(var1).Empty || variables.Find(var2).Empty {
// 			return true
// 		}
// 		return variables.Find(var1).Value > variables.Find(var2).Value
// 	}}
// }

// AllEquals Constraint generator that checks that all given variables are equal
func AllEquals(varnames ...VariableName) Constraint {
	if len(varnames) <= 0 {
		panic("Not enough variable names provided!")
	}
	return Constraint{varnames, func(variables *Variables) bool {
		foundFirst := false
		var first Variable
		// find first non empty variable to compare all others to
		for _, varname := range varnames {
			next := variables.Find(varname)
			if !next.Empty {
				first = next
				foundFirst = true
			}
		}
		if !foundFirst {
			return true // all variables are empty
		}
		flag := true
		// compare all variables to the first non-empty one, ignoring empty variables
		for _, varname := range varnames {
			next := variables.Find(varname)
			flag = flag && (first.Value == next.Value || next.Empty)
		}
		return flag
	}}
}

// AllUnique Constraint generator to check if all variable values are unique
func AllUnique(varnames ...VariableName) Constraint {
	if len(varnames) <= 0 {
		panic("Not enough variable names provided!")
	}
	return Constraint{varnames, func(variables *Variables) bool {
		uniqueMap := make(map[interface{}]struct{})
		for _, varname := range varnames {
			next := variables.Find(varname)

			// if our variable isn't empty and we have already assigned to the map with its value
			if _, ok := uniqueMap[next.Value]; ok && !next.Empty {
				return false
			}
			uniqueMap[next.Value] = struct{}{}
		}
		return true
	}}
}
