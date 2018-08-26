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

// CSPSolver struct for holding solver state
type CSPSolver struct {
	State CSPState
}

// NewCSPSolver create a solver
func NewCSPSolver(vars Variables, constraints Constraints, maxDepth int) CSPSolver {
	return CSPSolver{CSPState{vars, constraints, maxDepth}}
}

// Solve solves for values in the CSP
func (solver *CSPSolver) Solve() bool {
	return reduce(&solver.State, 0)
}

// implements backtracking search
func reduce(state *CSPState, depth int) bool {
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
