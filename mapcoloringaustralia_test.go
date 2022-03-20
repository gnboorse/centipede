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

	"github.com/stretchr/testify/assert"
)

func TestMapColoringAustralia(t *testing.T) {
	colors := Domain[string]{"red", "green", "blue"}

	// set a variable for each of the provinces
	// domain for the variable is the range of colors
	vars := Variables[string]{
		// each has:   <name>,      <domain>
		NewVariable("WA", colors),
		NewVariable("NT", colors),
		NewVariable("Q", colors),
		NewVariable("NSW", colors),
		NewVariable("V", colors),
		NewVariable("SA", colors),
		NewVariable("T", colors),
	}

	// bordering provinces cannot be equal.
	// See https://en.wikipedia.org/wiki/States_and_territories_of_Australia
	constraints := Constraints[string]{
		NotEquals[string]("WA", "NT"),
		NotEquals[string]("WA", "SA"),
		NotEquals[string]("NT", "SA"),
		NotEquals[string]("NT", "Q"),
		NotEquals[string]("Q", "SA"),
		NotEquals[string]("Q", "NSW"),
		NotEquals[string]("NSW", "V"),
		NotEquals[string]("NSW", "SA"),
		NotEquals[string]("V", "SA"),
	}

	// create the solver with a maximum depth of 500
	solver := NewBackTrackingCSPSolver(vars, constraints)
	success, err := solver.Solve(context.TODO()) // run the solution
	assert.Nil(t, err)

	assert.True(t, success)
	values := map[string]string{}
	for _, variable := range solver.State.Vars {
		values[string(variable.Name)] = variable.Value
	}

	assert.Equal(t, values["WA"], "red")
	assert.Equal(t, values["NT"], "green")
	assert.Equal(t, values["Q"], "red")
	assert.Equal(t, values["NSW"], "green")
	assert.Equal(t, values["V"], "red")
	assert.Equal(t, values["SA"], "blue")
	assert.Equal(t, values["T"], "red")
}
