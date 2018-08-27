package centipede

import (
	"fmt"
)

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

// MakeArcConsistent algorithm based off of AC-3 used to make the
// given CSP fully arc consistent.
// https://en.wikipedia.org/wiki/AC-3_algorithm
func (state *CSPState) MakeArcConsistent() {
	// create queue of indices and fill it with constraints
	queue := make([]int, 0)
	for i := range state.Constraints {
		queue = append(queue, i)
	}
	// loop until the queue is empty
	for len(queue) > 0 {
		// pop first item off of queue
		index := queue[0]
		queue = queue[1:]
		constraint := state.Constraints[index]
		// only consider binary constraints
		if len(constraint.Vars) == 2 {
			// must be arc consistent both ways
			change1, domain1 := arcReduce(constraint.Vars[0], constraint.Vars[1], constraint, state)
			change2, domain2 := arcReduce(constraint.Vars[1], constraint.Vars[0], constraint, state)

			if change1 {
				if len(domain1) == 0 {
					panic(fmt.Sprintf("Domain reduced to empty slice for constraint %v", constraint))
				}
				state.Vars.SetDomain(constraint.Vars[0], domain1)
				// add all neighbors of X excluding Y
				for index2, constraint2 := range state.Constraints {
					if constraint2.Vars.Contains(constraint.Vars[0]) && !constraint2.Vars.Contains(constraint.Vars[1]) {
						queue = append(queue, index2)
					}
				}
			}

			if change2 {
				if len(domain2) == 0 {
					panic(fmt.Sprintf("Domain reduced to empty slice for constraint %v", constraint))
				}
				state.Vars.SetDomain(constraint.Vars[1], domain2)
				// add all neighbors of X excluding Y
				for index2, constraint2 := range state.Constraints {
					if constraint2.Vars.Contains(constraint.Vars[1]) && !constraint2.Vars.Contains(constraint.Vars[0]) {
						queue = append(queue, index2)
					}
				}
			}
		}
	}
}

// arcReduce reduce the domain of both vars on a binary constraint using
// arc consistency
func arcReduce(nameX, nameY VariableName, constraint Constraint, state *CSPState) (bool, Domain) {
	var modifiedDomain Domain
	X := state.Vars.Find(nameX)
	Y := state.Vars.Find(nameY)
	// if X is already assigned to, domain of X is simply the value of X
	var dxValues []interface{} // values of X
	if !X.Empty {
		dxValues = []interface{}{X.Value}
	} else {
		dxValues = X.Domain
	}
	// if Y is already assigned to, domain of Y is simply the value of Y
	var dyValues []interface{} // values of Y
	if !Y.Empty {
		dyValues = []interface{}{Y.Value}
	} else {
		dyValues = Y.Domain
	}

	modifiedDomain = dxValues
	change := false

	// iterate over values of X,
	for _, vx := range dxValues {
		foundvy := false
		for _, vy := range dyValues {
			tempVars := Variables{Variable{X.Name, vx, dxValues, false}, Variable{Y.Name, vy, dyValues, false}}
			if constraint.ConstraintFunction(&tempVars) {
				foundvy = true
				break
			}
		}
		if !foundvy { // no corresponding vy for vx
			modifiedDomain = modifiedDomain.Remove(vx)
			change = true
		}
	}
	return change, modifiedDomain
}

// todo: add support here for Node consistency
