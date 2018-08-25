package main

// IntCSPSolver struct for holding solver state
type IntCSPSolver struct {
	State IntCSPState
}

// NewIntCSPSolver create a solver
func NewIntCSPSolver(vars IntVariables, constraints IntConstraints, maxDepth int) IntCSPSolver {
	return IntCSPSolver{IntCSPState{vars, constraints, maxDepth}}
}

// Solve solves for values in the CSP
func (solver *IntCSPSolver) Solve() bool {
	return reduce(&solver.State, 0)
}

// implements backtracking search
func reduce(state *IntCSPState, depth int) bool {
	// iterate over unassigned variables
	for i, variable := range state.Vars {
		// ignore variables that have been set
		if variable.Empty {
			// iterate over options in the domain
			for _, option := range variable.Domain {
				// set variable
				state.Vars[i].SetValue(option)
				// check if this is valid
				if state.Constraints.AllSatisfied(state.Vars) {
					// check if complete
					if state.Vars.Complete() {
						// we have a full solution of valid values
						return true
					} else if depth >= state.MaxDepth {
						// don't descend too far, bottom out recursion
						return false
					} else {
						// go down another level
						if reduce(state, depth+1) {
							return true
						} // else continue with the domain loop (Backtrack)
					}
				}
			}
			// unset variable and try with a different one first
			state.Vars[i].Unset()
		}

	}
	return false
}
