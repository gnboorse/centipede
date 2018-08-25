package main

import (
	"fmt"
)

// FloatConstraints collection type for FloatConstraint
type FloatConstraints []FloatConstraint

// AllSatisfied check if a collection of FloatConstraints are satisfied
func (constraints *FloatConstraints) AllSatisfied(variables FloatVariables) bool {
	flag := true
	for _, constraint := range *constraints {
		flag = flag && constraint.Satisfied(variables)
	}
	return flag
}

// FloatVariablesConstraintFunction function used to determine validity of FloatVariables
type FloatVariablesConstraintFunction func(variables FloatVariables) bool

// FloatConstraint CSP constraint considering integer variables
type FloatConstraint struct {
	Vars               VariableNames
	constraintFunction FloatVariablesConstraintFunction
}

// Satisfied checks to see if the given FloatConstraint is satisfied by the variables presented
func (constraint *FloatConstraint) Satisfied(variables FloatVariables) bool {
	constraintVariablesSatisfied := true
	domainSatisfied := true

	for _, varname := range constraint.Vars {
		// make sure FloatVariables contains an object for each name in FloatConstraint.Vars
		constraintVariablesSatisfied = constraintVariablesSatisfied && (variables.Contains(varname))
	}

	for _, variable := range variables {
		// make sure each FloatVariable being passed in has a value consistent with its domain or is empty
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

// EqualsFloat FloatConstraint generator that checks if two vars are equal
func EqualsFloat(var1 VariableName, var2 VariableName) FloatConstraint {
	return FloatConstraint{VariableNames{var1, var2}, func(variables FloatVariables) bool {
		if variables.Find(var1).Empty || variables.Find(var2).Empty {
			return true
		}
		return variables.Find(var1).Value == variables.Find(var2).Value
	}}
}

// NotEqualsFloat FloatConstraint generator that checks if two vars are not equal
func NotEqualsFloat(var1 VariableName, var2 VariableName) FloatConstraint {
	return FloatConstraint{VariableNames{var1, var2}, func(variables FloatVariables) bool {
		if variables.Find(var1).Empty || variables.Find(var2).Empty {
			return true
		}
		return variables.Find(var1).Value != variables.Find(var2).Value
	}}
}

// LessThanFloat FloatConstraint generator that checks if first variable is less than second variable
func LessThanFloat(var1 VariableName, var2 VariableName) FloatConstraint {
	return FloatConstraint{VariableNames{var1, var1}, func(variables FloatVariables) bool {
		if variables.Find(var1).Empty || variables.Find(var2).Empty {
			return true
		}
		return variables.Find(var1).Value < variables.Find(var2).Value
	}}
}

// GreaterThanFloat FloatConstraint generator that checks if first variable is less than second variable
func GreaterThanFloat(var1 VariableName, var2 VariableName) FloatConstraint {
	return FloatConstraint{VariableNames{var1, var1}, func(variables FloatVariables) bool {
		if variables.Find(var1).Empty || variables.Find(var2).Empty {
			return true
		}
		return variables.Find(var1).Value > variables.Find(var2).Value
	}}
}

// AllEqualsFloat FloatConstraint generator that checks that all given variables are equal
func AllEqualsFloat(varnames ...VariableName) FloatConstraint {
	if len(varnames) <= 0 {
		panic("Not enough variable names provided!")
	}
	return FloatConstraint{varnames, func(variables FloatVariables) bool {
		foundFirst := false
		var first FloatVariable
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

// AllUniqueFloat FloatConstraint generator to check if all variable values are unique
func AllUniqueFloat(varnames ...VariableName) FloatConstraint {
	if len(varnames) <= 0 {
		panic("Not enough variable names provided!")
	}
	return FloatConstraint{varnames, func(variables FloatVariables) bool {
		uniqueMap := make(map[float32]struct{})
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
