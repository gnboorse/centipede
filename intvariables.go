package main

import "fmt"

// IntVariable indicates a CSP variable of int type
type IntVariable struct {
	Name   VariableName
	Value  int
	Domain IntDomain
	Empty  bool
}

// NewIntVariable constructor for IntVariable type
func NewIntVariable(name VariableName, domain IntDomain) IntVariable {
	return IntVariable{name, 0, domain, true}
}

// SetValue setter for IntVariable value field
func (variable *IntVariable) SetValue(value int) {
	variable.Value = value
	variable.Empty = false
}

// Unset the variable
func (variable *IntVariable) Unset() {
	variable.Empty = true
	var i int
	variable.Value = i
}

// IntVariables collection type for int type variables
type IntVariables []IntVariable

// SetValue setter for IntVariables collection
func (variables *IntVariables) SetValue(name VariableName, value int) {
	foundIndex := -1

	for index, variable := range *variables {
		if variable.Name == name {
			foundIndex = index
		}
	}
	if !(foundIndex >= 0) {
		panic(fmt.Sprintf("IntVariable not found by name %v in variables %v", name, variables))
	} else {
		(*variables)[foundIndex].Value = value
		(*variables)[foundIndex].Empty = false

	}
}

// Find find an IntVariable by name in an IntVariables collection
func (variables *IntVariables) Find(name VariableName) IntVariable {
	for _, variable := range *variables {
		if variable.Name == name {
			return variable
		}
	}
	panic(fmt.Sprintf("IntVariable not found by name %v in variables %v", name, variables))
}

// Contains slice contains method for IntVariables
func (variables *IntVariables) Contains(name VariableName) bool {
	for _, variable := range *variables {
		if variable.Name == name {
			return true
		}
	}
	return false
}

// Unassigned return all unassigned variables
func (variables *IntVariables) Unassigned() IntVariables {
	unassigned := make(IntVariables, 0)
	for _, variable := range *variables {
		if variable.Empty {
			unassigned = append(unassigned, variable)
		}
	}
	return unassigned
}

// Complete indicates if all variables have been assigned to
func (variables *IntVariables) Complete() bool {
	return len(variables.Unassigned()) == 0
}
