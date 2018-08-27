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

// BackTrackingCSPSolver struct for holding solver state
type BackTrackingCSPSolver struct {
	State CSPState
}

// NewBackTrackingCSPSolver create a solver
func NewBackTrackingCSPSolver(vars Variables, constraints Constraints) BackTrackingCSPSolver {
	return BackTrackingCSPSolver{CSPState{vars, constraints}}
}

// Solve solves for values in the CSP
func (solver *BackTrackingCSPSolver) Solve() bool {
	return reduce(&solver.State)
}

// implements backtracking search
func reduce(state *CSPState) bool {
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

				if complete && satisfied {
					// we have a full solution
					return true
				} else if complete && !satisfied {
					// we have filled it in completely.
					// keep looping over the domain, but if that fails, we'll bottom out
					continue
				} else if !complete && satisfied {
					// go down a level to assign to another variable
					// fmt.Printf("Set variable with %v left unset, %#v\n", state.Vars.Unassigned(), state.Vars[i])
					printVars(&state.Vars)
					if reduce(state) {
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

func printVars(vars *Variables) {
	fmt.Printf("\n\n ==>")
	for _, v := range *vars {
		fmt.Printf(" (%v = %v) ", v.Name, v.Value)
	}
	fmt.Printf("\n")
}
