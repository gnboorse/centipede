// Copyright 2022 Gabriel Boorse

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

import "context"

// BackTrackingCSPSolver struct for holding solver state
type BackTrackingCSPSolver[T comparable] struct {
	State CSPState[T]
}

// NewBackTrackingCSPSolver create a solver
func NewBackTrackingCSPSolver[T comparable](vars Variables[T], constraints Constraints[T]) BackTrackingCSPSolver[T] {
	return BackTrackingCSPSolver[T]{CSPState[T]{vars, constraints, []Propagation[T]{}}}
}

// NewBackTrackingCSPSolverWithPropagation create a solver
func NewBackTrackingCSPSolverWithPropagation[T comparable](vars Variables[T], constraints Constraints[T], propagations Propagations[T]) BackTrackingCSPSolver[T] {
	return BackTrackingCSPSolver[T]{CSPState[T]{vars, constraints, propagations}}
}

// Solve solves for values in the CSP
func (solver *BackTrackingCSPSolver[T]) Solve(ctx context.Context) (bool, error) {
	b, err := RunWithContext[bool](ctx, func() bool {
		return reduce(&solver.State)
	})
	if b != nil && *b {
		return true, nil
	}
	return false, err
}

// implements backtracking search
func reduce[T comparable](state *CSPState[T]) bool {
	// iterate over unassigned variables
	for i := range state.Vars {
		// ignore variables that have been set
		if state.Vars[i].Empty {
			// iterate over options in the domain
			domainRemovals := make(DomainRemovals[T], 0)
			variableDomain := state.Vars[i].Domain
			for _, option := range variableDomain {
				// undo any attempts to do domain propagation
				state.Vars.ResetDomainRemovalEvaluation(domainRemovals)

				// set variable
				state.Vars[i].SetValue(option)

				// get the propagations
				domainRemovals = state.Propagations.Execute(VariableAssignment[T]{state.Vars[i].Name, option}, &state.Vars)
				// propagate through the rest of the variables
				state.Vars.EvaluateDomainRemovals(domainRemovals)

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
	return false
}
