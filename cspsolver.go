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
	return BackTrackingCSPSolver{CSPState{vars, constraints, []Propagation{}}}
}

// NewBackTrackingCSPSolverWithPropagation create a solver
func NewBackTrackingCSPSolverWithPropagation(vars Variables, constraints Constraints, propagations Propagations) BackTrackingCSPSolver {
	return BackTrackingCSPSolver{CSPState{vars, constraints, propagations}}
}

// Solve solves for values in the CSP
func (solver *BackTrackingCSPSolver) Solve() bool {
	return reduce(&solver.State)
}

// implements backtracking search
func reduce(state *CSPState) bool {
	// iterate over unassigned variables
	for i := range state.Vars {
		// ignore variables that have been set
		if state.Vars[i].Empty {
			// iterate over options in the domain
			domainRemovals := make(DomainRemovals, 0)
			variableDomain := state.Vars[i].Domain
			for _, option := range variableDomain {
				// undo any attempts to do domain propagation
				state.Vars.ResetDomainRemovalEvaluation(domainRemovals)

				// set variable
				fmt.Printf("Setting variable %v with value %v\n", state.Vars[i], option)
				state.Vars[i].SetValue(option)

				// get the propagations
				domainRemovals = state.Propagations.Execute(VariableAssignment{state.Vars[i].Name, option}, &state.Vars)
				// propagate through the rest of the variables
				state.Vars.EvaluateDomainRemovals(domainRemovals)

				// check if this is valid
				complete := state.Vars.Complete()
				satisfied := state.Constraints.AllSatisfied(&state.Vars)
				fmt.Printf("Completed = %v, Satisfied = %v\n", complete, satisfied)

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
					if reduce(state) {
						return true
					}
				} else { // !complete && !satisfied
					continue // keep looping over the domain, but if that fails we'll bottom out
				}
			}
			// reset domain removals
			state.Vars.ResetDomainRemovalEvaluation(domainRemovals)
			// unset variable and try with a different one first
			state.Vars[i].Unset()

		}

	}
	fmt.Printf("- Bottoming out of reduction loop.\n")
	return false
}
