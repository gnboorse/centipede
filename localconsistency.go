package centipede

import "fmt"

// SimplifyPreAssignment basic constraint propagation algorithm used to
// simplify variable domains before solving based on variables already
// assigned to. Condition: if a variable has been assigned to with
// a given value, remove that value from the domain of all variables
// mutually exclusive to it, i.e. if A != B and B = 2, remove 2
// from the domain of A.
func (state *CSPState) SimplifyPreAssignment() {

	for _, variable := range state.Vars {
		if !variable.Empty { // assigned to
			// get all constraints associated with this variable
			assignedConstraints := state.Constraints.FilterByName(variable.Name)
			for _, assignedConstraint := range assignedConstraints {
				for _, constraintVarName := range assignedConstraint.Vars {
					// don't compare the variable in question to itself
					if constraintVarName == variable.Name {
						continue
					}
					constrainedVariable := state.Vars.Find(constraintVarName)
					// continue if this is one of the variables already assigned to
					if !constrainedVariable.Empty {
						continue
					}
					// check to see if the assigned value is a possibility
					// for the unassigned variable we're comparing too
					if constrainedVariable.Domain.Contains(variable.Value) {
						resultBefore := assignedConstraint.ConstraintFunction(&state.Vars)
						state.Vars.SetValue(constrainedVariable.Name, variable.Value)
						resultAfter := assignedConstraint.ConstraintFunction(&state.Vars)
						state.Vars.Unset(constrainedVariable.Name)
						if resultBefore && !resultAfter {
							// safe to assume that variable and constrainedVariable
							// cannot both have this value. Remove this value from
							// the domain of constrainedVariable
							restrictedDomain := constrainedVariable.Domain.Remove(variable.Value)
							state.Vars.SetDomain(constrainedVariable.Name, restrictedDomain)
							// if domain has only one value, set the value of the variable to
							// avoid further complexity
							if len(restrictedDomain) == 1 {
								state.Vars.SetValue(constrainedVariable.Name, restrictedDomain[0])
							}
						}
					}
				}
			}
		}
	}
}

// MakeArcConsistent algorithm loosely based off of AC-3 used to make the
// given CSP fully arc consistent.
func (state *CSPState) MakeArcConsistent() {
	// create queue and fill it with constraints
	queue := make([]int, 0)
	for i := range state.Constraints {
		queue = append(queue, i) // queue contains indices
	}
	// loop until the queue is empty
	for len(queue) > 0 {
		// pop first item off of queue
		index := queue[0]
		queue = queue[1:]
		constraint := state.Constraints[index]
		fmt.Println(constraint)

		//todo: need to fill this out
	}
}

// todo: add support here for Node consistency, true Arc consistency, and some kind of satisfiability algorithm.
