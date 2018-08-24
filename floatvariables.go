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
func (floatVariable *FloatVariable) SetValue(value float32) {
	floatVariable.Value = value
	floatVariable.Empty = false
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
