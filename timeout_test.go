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

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimeout(t *testing.T) {
	vars := Variables[int]{
		NewVariable("A", IntRange(1, 10)),
		NewVariable("B", IntRange(1, 10)),
		NewVariable("C", IntRange(1, 10)),
	}

	constraints := Constraints[int]{
		Equals[int]("A", "B"), // A = B
		Constraint[int]{Vars: VariableNames{"A", "B"},
			ConstraintFunction: func(variables *Variables[int]) bool {
				time.Sleep(10 * time.Millisecond)
				if variables.Find("A").Empty || variables.Find("B").Empty {
					return true
				}
				return variables.Find("A").Value > variables.Find("B").Value
			}},
	}

	// solve the problem
	solver := NewBackTrackingCSPSolver(vars, constraints)
	d := time.Now().Add(20 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.TODO(), d)
	defer cancel()
	_, err := solver.Solve(ctx)
	assert.Equal(t, ErrExecutionCanceled, err)
}

func TestNoTimeout(t *testing.T) {
	vars := Variables[int]{
		NewVariable("A", IntRange(1, 10)),
		NewVariable("B", IntRange(1, 10)),
		NewVariable("C", IntRange(1, 10)),
	}

	constraints := Constraints[int]{
		Equals[int]("A", "B"), // A = B
		Constraint[int]{Vars: VariableNames{"A", "C"},
			ConstraintFunction: func(variables *Variables[int]) bool {
				time.Sleep(1 * time.Millisecond)
				if variables.Find("A").Empty || variables.Find("C").Empty {
					return true
				}
				return variables.Find("A").Value > variables.Find("C").Value
			}},
	}

	// solve the problem
	solver := NewBackTrackingCSPSolver(vars, constraints)
	d := time.Now().Add(200 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.TODO(), d)
	defer cancel()
	success, err := solver.Solve(ctx)
	assert.Nil(t, err)
	assert.True(t, success)
}
