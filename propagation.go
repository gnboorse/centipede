package centipede

// DomainRemoval a struct indicating removing an item from a domain
type DomainRemoval[T comparable] struct {
	// VariableName the variable whose domain we are pruning
	VariableName
	// Value the value to prune from its domain
	Value T
}

// DomainRemovals list type for DomainRemoval
type DomainRemovals[T comparable] []DomainRemoval[T]

// VariableAssignment a struct indicating assigning a value to a variable
type VariableAssignment[T comparable] struct {
	// VariableName the variable that was assigned to
	VariableName
	// Value the value being assigned
	Value T
}

// Propagation type representing a domain propagation
type Propagation[T comparable] struct {
	Vars VariableNames
	PropagationFunction[T]
}

// Propagations list type for Propagation
type Propagations[T comparable] []Propagation[T]

// PropagationFunction used to determine domain pruning
type PropagationFunction[T comparable] func(assignment VariableAssignment[T], variables *Variables[T]) []DomainRemoval[T]

// Execute this Propagation
func (propagation *Propagation[T]) Execute(assignment VariableAssignment[T], variables *Variables[T]) []DomainRemoval[T] {
	return propagation.PropagationFunction(assignment, variables)
}

// Execute all propagations in the list
func (propagations *Propagations[T]) Execute(assignment VariableAssignment[T], variables *Variables[T]) []DomainRemoval[T] {
	removals := make(DomainRemovals[T], 0)
	for _, propagation := range *propagations {
		// if this assignment is relevant to us
		if propagation.Vars.Contains(assignment.VariableName) {
			removals = append(removals, propagation.Execute(assignment, variables)...)
		}
	}
	return removals
}
