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

package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gnboorse/centipede"
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
func Sudoku() {

	// initialize variables
	vars := make(centipede.Variables, 0)
	constraints := make(centipede.Constraints, 0)

	letters := [9]string{"A", "B", "C", "D", "E", "F", "G", "H", "I"}

	// configure variables and block constraints
	for _, letter := range letters {
		letterVars := make(centipede.VariableNames, 0)
		for i := 1; i <= 9; i++ {
			varName := centipede.VariableName(letter + strconv.Itoa(i))
			// add vars like A1, A2, A3 ... A9, B1, B2, B3 ... B9 ... I9
			vars = append(vars, centipede.NewVariable(varName, getDomain(varName)))
			letterVars = append(letterVars, varName)
		}
		// for each block, add uniqueness constraint within block
		constraints = append(constraints, centipede.AllUnique(letterVars...))
	}

	// add horizontal constraints
	rowLetterSets := [3][3]string{{"A", "B", "C"}, {"D", "E", "F"}, {"G", "H", "I"}}
	rowNumberSets := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	for _, letterSet := range rowLetterSets {
		for _, numberSet := range rowNumberSets {
			rowVarNames := make(centipede.VariableNames, 0)
			for _, letter := range letterSet {
				for _, number := range numberSet {
					varName := centipede.VariableName(letter + strconv.Itoa(number))
					rowVarNames = append(rowVarNames, varName)
				}
			}
			// add uniqueness constraints
			constraints = append(constraints, centipede.AllUnique(rowVarNames...))
		}
	}

	// add vertical constraints
	columnLetterSets := [3][3]string{{"A", "D", "G"}, {"B", "E", "H"}, {"C", "F", "I"}}
	columnNumberSets := [3][3]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}}
	for _, letterSet := range columnLetterSets {
		for _, numberSet := range columnNumberSets {
			columnVarNames := make(centipede.VariableNames, 0)
			for _, letter := range letterSet {
				for _, number := range numberSet {
					varName := centipede.VariableName(letter + strconv.Itoa(number))
					columnVarNames = append(columnVarNames, varName)
				}
			}
			// add uniqueness constraints
			constraints = append(constraints, centipede.AllUnique(columnVarNames...))
		}
	}

	// solve the problem
	solver := centipede.NewCSPSolver(vars, constraints, 500)
	begin := time.Now()
	success := solver.Solve() // run the solution
	elapsed := time.Since(begin)

	// output results and time elapsed
	if success {
		fmt.Printf("Found solution in %s\n", elapsed)
		for _, variable := range solver.State.Vars {
			// print out values for each variable
			fmt.Printf("Variable %v = %v\n", variable.Name, variable.Value)
		}
	} else {
		fmt.Printf("Could not find solution in %s\n", elapsed)
	}

}

// prevent domain for certain variables from going
// outside of known values in the puzzle
// values here populated from: https://en.wikipedia.org/wiki/Sudoku#/media/File:Sudoku_Puzzle_by_L2G-20050714_standardized_layout.svg
func getDomain(varName centipede.VariableName) centipede.Domain {
	switch varName {
	case "A1":
		return centipede.Domain{5}
	case "A2":
		return centipede.Domain{3}
	case "A4":
		return centipede.Domain{6}
	case "A8":
		return centipede.Domain{9}
	case "A9":
		return centipede.Domain{8}
	case "B2":
		return centipede.Domain{7}
	case "B4":
		return centipede.Domain{1}
	case "B5":
		return centipede.Domain{9}
	case "B6":
		return centipede.Domain{5}
	case "C8":
		return centipede.Domain{6}
	case "D1":
		return centipede.Domain{8}
	case "D4":
		return centipede.Domain{4}
	case "D7":
		return centipede.Domain{7}
	case "E2":
		return centipede.Domain{6}
	case "E4":
		return centipede.Domain{8}
	case "E6":
		return centipede.Domain{3}
	case "E8":
		return centipede.Domain{2}
	case "F3":
		return centipede.Domain{3}
	case "F6":
		return centipede.Domain{1}
	case "F9":
		return centipede.Domain{6}
	case "G2":
		return centipede.Domain{6}
	case "H4":
		return centipede.Domain{4}
	case "H5":
		return centipede.Domain{1}
	case "H6":
		return centipede.Domain{9}
	case "H8":
		return centipede.Domain{8}
	case "I1":
		return centipede.Domain{2}
	case "I2":
		return centipede.Domain{8}
	case "I6":
		return centipede.Domain{5}
	case "I8":
		return centipede.Domain{7}
	case "I9":
		return centipede.Domain{9}
	default:
		return centipede.IntRange(1, 10)
	}
}
