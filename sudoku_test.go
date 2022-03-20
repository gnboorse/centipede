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
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Sudoku implementation of a Sudoku puzzle solver. The particular puzzle being
// solved here has been taken straight from the wikipedia page.
// Puzzle: https://en.wikipedia.org/wiki/Sudoku#/media/File:Sudoku_Puzzle_by_L2G-20050714_standardized_layout.svg
// Solution: https://en.wikipedia.org/wiki/Sudoku#/media/File:Sudoku_Puzzle_by_L2G-20050714_solution_standardized_layout.svg
// In this solution, cells are labeled by the 3x3 sector (out of the 9 large boxes) as a letter in the range A-I,
// and an integer in the range 1-9 indicating the cell's position in the sector.
// For reference, here is an example grid:
// [A1 A2 A3 B1 B2 B3 C1 C2 C3]
// [A4 A5 A6 B4 B5 B6 C4 C5 C6]
// [A7 A8 A9 B7 B8 B9 C7 C8 C9]
// [D1 D2 D3 E1 E2 E3 F1 F2 F3]
// [D4 D5 D6 E4 E5 E6 F4 F5 F6]
// [D7 D8 D9 E7 E8 E9 F7 F8 F9]
// [G1 G2 G3 H1 H2 H3 I1 I2 I3]
// [G4 G5 G6 H4 H5 H6 I4 I5 I6]
// [G7 G8 G9 H7 H8 H9 I7 I8 I9]
// This solution enforces Arc consistency on all binary constraints in the problem,
// resulting in a very fast solve
func TestSudoku(t *testing.T) {

	// initialize variables
	vars := make(Variables[int], 0)
	constraints := make(Constraints[int], 0)

	letters := [9]string{"A", "B", "C", "D", "E", "F", "G", "H", "I"}
	tenDomain := IntRange(1, 10)

	// configure variables and block constraints
	for _, letter := range letters {
		letterVars := make(VariableNames, 0)
		for i := 1; i <= 9; i++ {
			varName := VariableName(letter + strconv.Itoa(i))
			// add vars like A1, A2, A3 ... A9, B1, B2, B3 ... B9 ... I9
			vars = append(vars, NewVariable(varName, tenDomain))
			letterVars = append(letterVars, varName)
		}
		// for each block, add uniqueness constraint within block
		constraints = append(constraints, AllUnique[int](letterVars...)...)
	}

	// add horizontal constraints
	rowLetterSets := [3][3]string{{"A", "B", "C"}, {"D", "E", "F"}, {"G", "H", "I"}}
	rowNumberSets := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	for _, letterSet := range rowLetterSets {
		for _, numberSet := range rowNumberSets {
			rowVarNames := make(VariableNames, 0)
			for _, letter := range letterSet {
				for _, number := range numberSet {
					varName := VariableName(letter + strconv.Itoa(number))
					rowVarNames = append(rowVarNames, varName)
				}
			}
			// add uniqueness constraints
			constraints = append(constraints, AllUnique[int](rowVarNames...)...)
		}
	}

	// add vertical constraints
	columnLetterSets := [3][3]string{{"A", "D", "G"}, {"B", "E", "H"}, {"C", "F", "I"}}
	columnNumberSets := [3][3]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}}
	for _, letterSet := range columnLetterSets {
		for _, numberSet := range columnNumberSets {
			columnVarNames := make(VariableNames, 0)
			for _, letter := range letterSet {
				for _, number := range numberSet {
					varName := VariableName(letter + strconv.Itoa(number))
					columnVarNames = append(columnVarNames, varName)
				}
			}
			// add uniqueness constraints
			constraints = append(constraints, AllUnique[int](columnVarNames...)...)
		}
	}
	// set values already known
	vars.SetValue("A1", 5)
	vars.SetValue("A2", 3)
	vars.SetValue("A4", 6)
	vars.SetValue("A8", 9)
	vars.SetValue("A9", 8)
	vars.SetValue("B2", 7)
	vars.SetValue("B4", 1)
	vars.SetValue("B5", 9)
	vars.SetValue("B6", 5)
	vars.SetValue("C8", 6)
	vars.SetValue("D1", 8)
	vars.SetValue("D4", 4)
	vars.SetValue("D7", 7)
	vars.SetValue("E2", 6)
	vars.SetValue("E4", 8)
	vars.SetValue("E6", 3)
	vars.SetValue("E8", 2)
	vars.SetValue("F3", 3)
	vars.SetValue("F6", 1)
	vars.SetValue("F9", 6)
	vars.SetValue("G2", 6)
	vars.SetValue("H4", 4)
	vars.SetValue("H5", 1)
	vars.SetValue("H6", 9)
	vars.SetValue("H8", 8)
	vars.SetValue("I1", 2)
	vars.SetValue("I2", 8)
	vars.SetValue("I6", 5)
	vars.SetValue("I8", 7)
	vars.SetValue("I9", 9)

	// create solver
	solver := NewBackTrackingCSPSolver(vars, constraints)

	// simplify variable domains following initial assignment
	solver.State.MakeArcConsistent(context.TODO())
	success, err := solver.Solve(context.TODO()) // run the solution
	assert.Nil(t, err)

	assert.True(t, success)

	// check that we have a valid sudoku solution

	for _, letterSet := range rowLetterSets {
		for _, numberSet := range rowNumberSets {
			sum := 0
			for _, letter := range letterSet {
				for _, number := range numberSet {
					varName := VariableName(letter + strconv.Itoa(number))
					variable := solver.State.Vars.Find(varName)
					sum += variable.Value
				}
			}
			assert.Equal(t, 45, sum)
		}
	}

	for _, letterSet := range columnLetterSets {
		for _, numberSet := range columnNumberSets {
			sum := 0
			for _, letter := range letterSet {
				for _, number := range numberSet {
					varName := VariableName(letter + strconv.Itoa(number))
					variable := solver.State.Vars.Find(varName)
					sum += variable.Value
				}
			}
			assert.Equal(t, 45, sum)
		}
	}

	for _, letter := range letters {
		sum := 0
		for num := 1; num <= 9; num++ {
			varName := VariableName(letter + strconv.Itoa(num))
			variable := solver.State.Vars.Find(varName)
			sum += variable.Value
		}
		assert.Equal(t, 45, sum)
	}

}
