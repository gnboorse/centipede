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
	log.Debugf("Depth %v of solve.\n", depth)
	log.Debugf("Current state is: %v\n", state.Vars)
	// iterate over unassigned variables
	for i, variable := range state.Vars {
		// ignore variables that have been set
		if variable.Empty {
			log.Debugf("Variable %v, Depth %v\n", variable.Name, depth)
			// iterate over options in the domain
			for _, option := range variable.Domain {
				// set variable
				state.Vars[i].SetValue(option)
				log.Debugf("Value %v, Variable %v, Depth %v -> %v\n", option, variable.Name, depth, state.Vars)
				// check if this is valid
				if state.Constraints.AllSatisfied(state.Vars) {
					// check if complete
					if state.Vars.Complete() {
						// we have a full solution of valid values
						log.Debugf("Solution found at depth %v\n", depth)
						return true
					} else if depth >= state.MaxDepth {
						// don't descend too far, bottom out recursion
						log.Debugf("Hit max depth at depth %v\n", depth)
						return false
					} else {
						// go down another level
						log.Debugf("Going down another level\n")
						if reduce(state, depth+1) {
							return true
						} // else continue with the domain loop (Backtrack)
					}
				}
			}
			// unset variable and try with a different one first
			log.Debugf("Unsetting variable %v at depth %v\n", variable.Name, depth)
			state.Vars[i].Unset()
		}

	}
	log.Debugf("Bottomed out at depth %v\n", depth)
	return false
}
