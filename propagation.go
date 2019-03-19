package centipede

// DomainRemoval a struct indicating removing an item from a domain
type DomainRemoval struct {
	// VariableName the variable whose domain we are pruning
	VariableName
	// Value the value to prune from its domain
	Value interface{}
}

// DomainRemovals list type for DomainRemoval
type DomainRemovals []DomainRemoval

// VariableAssignment a struct indicating assigning a value to a variable
type VariableAssignment struct {
	// VariableName the variable that was assigned to
	VariableName
	// Value the value being assigned
	Value interface{}
}

// Propagation type representing a domain propagation
type Propagation struct {
	Vars VariableNames
	PropagationFunction
}

// Propagations list type for Propagation
type Propagations []Propagation

// PropagationFunction used to determine domain pruning
type PropagationFunction func(assignment VariableAssignment, variables *Variables) []DomainRemoval

// Execute this Propagation
func (propagation *Propagation) Execute(assignment VariableAssignment, variables *Variables) []DomainRemoval {
	return propagation.PropagationFunction(assignment, variables)
}

// Execute all propagations in the list
func (propagations *Propagations) Execute(assignment VariableAssignment, variables *Variables) []DomainRemoval {
	removals := make(DomainRemovals, 0)
	for _, propagation := range *propagations {
		// if this assignment is relevant to us
		if propagation.Vars.Contains(assignment.VariableName) {
			removals = append(removals, propagation.Execute(assignment, variables)...)
		}
	}
	return removals
}
