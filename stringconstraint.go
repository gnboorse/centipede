package main

import (
	"fmt"
)

// StringConstraints collection type for StringConstraint
type StringConstraints []StringConstraint

// AllSatisfied check if a collection of StringConstraints are satisfied
func (constraints *StringConstraints) AllSatisfied(variables StringVariables) bool {
	flag := true
	for _, constraint := range *constraints {
		flag = flag && constraint.Satisfied(variables)
	}
	return flag
}

// StringVariablesConstraintFunction function used to determine validity of StringVariables
type StringVariablesConstraintFunction func(variables StringVariables) bool

// StringConstraint CSP constraint considering integer variables
type StringConstraint struct {
	Vars               VariableNames
	constraintFunction StringVariablesConstraintFunction
}

// Satisfied checks to see if the given StringConstraint is satisfied by the variables presented
func (constraint *StringConstraint) Satisfied(variables StringVariables) bool {
	constraintVariablesSatisfied := true
	domainSatisfied := true

	for _, varname := range constraint.Vars {
		// make sure StringVariables contains an object for each name in StringConstraint.Vars
		constraintVariablesSatisfied = constraintVariablesSatisfied && (variables.Contains(varname))
	}

	for _, variable := range variables {
		// make sure each StringVariable being passed in has a value consistent with its domain or is empty
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

// EqualsString StringConstraint generator that checks if two vars are equal
func EqualsString(var1 VariableName, var2 VariableName) StringConstraint {
	return StringConstraint{VariableNames{var1, var2}, func(variables StringVariables) bool {
		if variables.Find(var1).Empty || variables.Find(var2).Empty {
			return true
		}
		return variables.Find(var1).Value == variables.Find(var2).Value
	}}
}

// NotEqualsString StringConstraint generator that checks if two vars are not equal
func NotEqualsString(var1 VariableName, var2 VariableName) StringConstraint {
	return StringConstraint{VariableNames{var1, var2}, func(variables StringVariables) bool {
		if variables.Find(var1).Empty || variables.Find(var2).Empty {
			return true
		}
		return variables.Find(var1).Value != variables.Find(var2).Value
	}}
}

// LessThanString StringConstraint generator that checks if first variable is less than second variable
func LessThanString(var1 VariableName, var2 VariableName) StringConstraint {
	return StringConstraint{VariableNames{var1, var1}, func(variables StringVariables) bool {
		if variables.Find(var1).Empty || variables.Find(var2).Empty {
			return true
		}
		return variables.Find(var1).Value < variables.Find(var2).Value
	}}
}

// GreaterThanString StringConstraint generator that checks if first variable is less than second variable
func GreaterThanString(var1 VariableName, var2 VariableName) StringConstraint {
	return StringConstraint{VariableNames{var1, var1}, func(variables StringVariables) bool {
		if variables.Find(var1).Empty || variables.Find(var2).Empty {
			return true
		}
		return variables.Find(var1).Value > variables.Find(var2).Value
	}}
}

// AllEqualsString StringConstraint generator that checks that all given variables are equal
func AllEqualsString(varnames ...VariableName) StringConstraint {
	if len(varnames) <= 0 {
		panic("Not enough variable names provided!")
	}
	return StringConstraint{varnames, func(variables StringVariables) bool {
		foundFirst := false
		var first StringVariable
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

// AllUniqueString StringConstraint generator to check if all variable values are unique
func AllUniqueString(varnames ...VariableName) StringConstraint {
	if len(varnames) <= 0 {
		panic("Not enough variable names provided!")
	}
	return StringConstraint{varnames, func(variables StringVariables) bool {
		uniqueMap := make(map[string]struct{})
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
