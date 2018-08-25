package main

import "fmt"

// FloatVariable indicates a CSP variable of float32 type
type FloatVariable struct {
	Name   VariableName
	Value  float32
	Domain FloatDomain
	Empty  bool
}

// NewFloatVariable constructor for FloatVariable type
func NewFloatVariable(name VariableName, domain FloatDomain) FloatVariable {
	return FloatVariable{name, 0.0, domain, true}
}

// SetValue setter for FloatVariable value field
func (variable *FloatVariable) SetValue(value float32) {
	variable.Value = value
	variable.Empty = false
}

// Unset the variable
func (variable *FloatVariable) Unset() {
	variable.Empty = true
	var f float32
	variable.Value = f
}

// FloatVariables collection type for float type variables
type FloatVariables []FloatVariable

// SetValue setter for FloatVariables collection
func (variables *FloatVariables) SetValue(name VariableName, value float32) {
	foundIndex := -1

	for index, variable := range *variables {
		if variable.Name == name {
			foundIndex = index
		}
	}
	if !(foundIndex >= 0) {
		panic(fmt.Sprintf("FloatVariable not found by name %v in variables %v", name, variables))
	} else {
		(*variables)[foundIndex].Value = value
		(*variables)[foundIndex].Empty = false

	}
}

// Find find a FloatVariable by name in a FloatVariables collection
func (variables *FloatVariables) Find(name VariableName) FloatVariable {
	for _, variable := range *variables {
		if variable.Name == name {
			return variable
		}
	}
	panic(fmt.Sprintf("FloatVariable not found by name %v in variables %v", name, variables))
}

// Contains slice contains method for FloatVariables
func (variables *FloatVariables) Contains(name VariableName) bool {
	for _, variable := range *variables {
		if variable.Name == name {
			return true
		}
	}
	return false
}

// Unassigned return all unassigned variables
func (variables *FloatVariables) Unassigned() FloatVariables {
	unassigned := make(FloatVariables, 0)
	for _, variable := range *variables {
		if variable.Empty {
			unassigned = append(unassigned, variable)
		}
	}
	return unassigned
}

// Complete indicates if all variables have been assigned to
func (variables *FloatVariables) Complete() bool {
	return len(variables.Unassigned()) == 0
}
