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
	return reduce(&solver.State, 1)
}

// IterativeDeepeningSolve solve for values in the CSP with an interative deepening strategy
func (solver *CSPSolver) IterativeDeepeningSolve() bool {
	max := solver.State.MaxDepth
	freshState := solver.State // state at beginning of solve
	for i := 1; i <= max; i++ {
		solver.State = freshState // reset to beginning state
		solver.State.MaxDepth = i
		fmt.Printf("Attempting to solve with max depth %v\n", i)
		if solver.Solve() {
			return true
		}
	}
	return false
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
				complete := state.Vars.Complete()
				satisfied := state.Constraints.AllSatisfied(&state.Vars)
				tooDeep := depth >= state.MaxDepth

				if complete && satisfied {
					// we have a full solution
					return true
				} else if complete && !satisfied {
					// we have filled it in completely.
					// keep looping over the domain, but if that fails, we'll bottom out
					continue
				} else if !complete && satisfied {
					// we have hit our max limit.
					// continue domain loop instead of going further down
					if tooDeep {
						continue
					}
					// go down a level to assign to another variable
					if reduce(state, depth+1) {
						return true
					}
				} else { // !complete && !satisfied
					continue // keep looping over the domain, but if that fails we'll bottom out
				}
			}
			// unset variable and try with a different one first
			state.Vars[i].Unset()
		}

	}
	return false
}
