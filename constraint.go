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

		if !variable.Domain.Contains(variable.Value) && !variable.Empty {
			fmt.Printf("Variable %v with domain %v does not support value %v\n", variable.Name, variable.Domain, variable.Value)
		}
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
func AllEquals(varnames ...VariableName) Constraints {
	return mapCombinationsToBinaryConstraint(varnames, Equals)
}

// AllUnique Constraint generator to check if all variable values are unique
func AllUnique(varnames ...VariableName) Constraints {
	return mapCombinationsToBinaryConstraint(varnames, NotEquals)
}

func mapCombinationsToBinaryConstraint(varnames VariableNames, fx func(VariableName, VariableName) Constraint) Constraints {
	if len(varnames) <= 0 {
		panic("Not enough variable names provided!")
	}
	constraints := make(Constraints, 0)
	// map of commutative, unique pairs
	uniqueMap := make(map[[2]VariableName]struct{})
	for _, name1 := range varnames {
		for _, name2 := range varnames {
			// if we've already seen this pair before, continue
			if _, ok := uniqueMap[[2]VariableName{name1, name2}]; ok {
				continue
			} else if _, ok := uniqueMap[[2]VariableName{name2, name1}]; ok {
				continue
			}
			// we don't want to make constraints for A == A or A != A
			if name1 == name2 {
				continue
			}
			uniqueMap[[2]VariableName{name1, name2}] = struct{}{}
			constraints = append(constraints, fx(name1, name2))
		}
	}
	return constraints
}
