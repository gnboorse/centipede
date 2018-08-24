package main

import (
	"fmt"
)

// IntConstraints collection type for IntConstraint
type IntConstraints []IntConstraint

// AllSatisfied check if a collection of IntConstraints are satisfied
func (constraints *IntConstraints) AllSatisfied(variables IntVariables) bool {
	flag := true
	for _, constraint := range *constraints {
		flag = flag && constraint.Satisfied(variables)
	}
	return flag
}

// IntVariablesConstraintFunction function used to determine validity of IntVariables
type IntVariablesConstraintFunction func(variables IntVariables) bool

// IntConstraint CSP constraint considering integer variables
type IntConstraint struct {
	Vars               VariableNames
	constraintFunction IntVariablesConstraintFunction
}

// Satisfied checks to see if the given IntConstraint is satisfied by the variables presented
func (constraint *IntConstraint) Satisfied(variables IntVariables) bool {
	constraintVariablesSatisfied := true
	domainSatisfied := true

	for _, varname := range constraint.Vars {
		// make sure IntVariables contains an object for each name in IntConstraint.Vars
		constraintVariablesSatisfied = constraintVariablesSatisfied && (variables.Contains(varname))
	}

	for _, variable := range variables {
		// make sure each IntVariable being passed in has a value consistent with its domain or is empty
		domainSatisfied = domainSatisfied && (variable.Domain.Contains(variable.Value) || variable.Empty)
	}
	if !constraintVariablesSatisfied {
		panic(fmt.Sprintf("Insufficient variables provided. Expected %v", constraint.Vars))
	}
	if !domainSatisfied {
		panic("Variables do not satisfy the domains given.")
	}
	// now finally call the constraint function
	return constraint.constraintFunction(variables)
}

// EqualsInt IntConstraint generator that checks if two vars are equal
func EqualsInt(var1 VariableName, var2 VariableName) IntConstraint {
	return IntConstraint{VariableNames{var1, var2}, func(variables IntVariables) bool {
		if variables.Find(var1).Empty || variables.Find(var2).Empty {
			return true
		}
		return variables.Find(var1).Value == variables.Find(var2).Value
	}}
}

// NotEqualsInt IntConstraint generator that checks if two vars are not equal
func NotEqualsInt(var1 VariableName, var2 VariableName) IntConstraint {
	return IntConstraint{VariableNames{var1, var2}, func(variables IntVariables) bool {
		if variables.Find(var1).Empty || variables.Find(var2).Empty {
			return true
		}
		return variables.Find(var1).Value != variables.Find(var2).Value
	}}
}

// AllEqualsInt IntConstraint generator that checks that all given variables are equal
func AllEqualsInt(varnames ...VariableName) IntConstraint {
	if len(varnames) <= 0 {
		panic("Not enough variable names provided!")
	}
	return IntConstraint{varnames, func(variables IntVariables) bool {
		foundFirst := false
		var first IntVariable
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

// AllUniqueInt IntConstraint generator to check if all variable values are unique
func AllUniqueInt(varnames ...VariableName) IntConstraint {
	if len(varnames) <= 0 {
		panic("Not enough variable names provided!")
	}
	return IntConstraint{varnames, func(variables IntVariables) bool {
		uniqueMap := make(map[int]struct{})
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
