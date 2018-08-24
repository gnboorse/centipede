package main

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
