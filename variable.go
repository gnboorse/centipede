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

import "fmt"

// VariableName is our string type for names of variables
type VariableName string

// VariableNames collection type for VariableName
type VariableNames []VariableName

// Contains slice contains method for VariableNames
func (varnames *VariableNames) Contains(varname VariableName) bool {
	for _, item := range *varnames {
		if item == varname {
			return true
		}
	}
	return false
}

// Variable indicates a CSP variable of interface{} type
type Variable struct {
	Name   VariableName
	Value  interface{}
	Domain Domain
	Empty  bool
}

// NewVariable constructor for Variable type
func NewVariable(name VariableName, domain Domain) Variable {
	return Variable{name, 0, domain, true}
}

// SetValue setter for Variable value field
func (variable *Variable) SetValue(value interface{}) {
	variable.Value = value
	variable.Empty = false
}

// Unset the variable
func (variable *Variable) Unset() {
	variable.Empty = true
	var i interface{}
	variable.Value = i
}

// Variables collection type for interface{} type variables
type Variables []Variable

// SetValue setter for Variables collection
func (variables *Variables) SetValue(name VariableName, value interface{}) {
	foundIndex := -1

	for index, variable := range *variables {
		if variable.Name == name {
			foundIndex = index
		}
	}
	if !(foundIndex >= 0) {
		panic(fmt.Sprintf("Variable not found by name %v in variables %v", name, variables))
	} else {
		(*variables)[foundIndex].Value = value
		(*variables)[foundIndex].Empty = false

	}
}

// Find find an Variable by name in an Variables collection
func (variables *Variables) Find(name VariableName) Variable {
	for _, variable := range *variables {
		if variable.Name == name {
			return variable
		}
	}
	panic(fmt.Sprintf("Variable not found by name %v in variables %v", name, variables))
}

// Contains slice contains method for Variables
func (variables *Variables) Contains(name VariableName) bool {
	for _, variable := range *variables {
		if variable.Name == name {
			return true
		}
	}
	return false
}

// Unassigned return all unassigned variables
func (variables *Variables) Unassigned() Variables {
	unassigned := make(Variables, 0)
	for _, variable := range *variables {
		if variable.Empty {
			unassigned = append(unassigned, variable)
		}
	}
	return unassigned
}

// Complete indicates if all variables have been assigned to
func (variables *Variables) Complete() bool {
	return len(variables.Unassigned()) == 0
}
