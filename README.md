# Centipede - Constraint Satisfaction Problem Solver for Go

![badge](https://github.com/gnboorse/centipede/actions/workflows/go.yml/badge.svg)

Centipede is a Constraint Satisfaction Problem solver written in Golang. [Learn more about CSPs](https://en.wikipedia.org/wiki/Constraint_satisfaction_problem).

There is also a very informative slide deck about CSPs available from Stanford University [here](https://web.stanford.edu/class/cs227/Lectures/lec14.pdf).

## Features

- Problems are defined using sets of `Variable`, `Constraint`, and `Domain`. Some convenient generators have been provided for `Constraint` and `Domain`.
- `Variable` values can be set to values of any `comparable` data type in Go (using generic types). Mixing datatypes in variables that are compared to each other is not currently possible.
- The search algorithm used in this library is an implementation of [backtracking search](https://en.wikipedia.org/wiki/Backtracking).
- The solution of many complex problems can be simplified by enforcing [arc consistency](https://en.wikipedia.org/wiki/Local_consistency#Arc_consistency). This library provides an implementation of the popular [AC-3 algorithm](https://en.wikipedia.org/wiki/AC-3_algorithm) as `solver.State.MakeArcConsistent()`. Call this method before calling `solver.Solve()` to achieve best results.
  - See the [Sudoku solver](sudoku_test.go) for an example of how to use arc consistency.

## Project Status

**As of March 2022, this project has been updated to take advantage of the generic types that were added in Go 1.18**. Previously, this project was relying heavily on storing values using a type of `interface{}` and casting the values to their respective types (i.e. `value.(int)`) when necessary. This was extremely inconvenient, and now, three years after it was first created, generic types have greatly simplified library usage. For more information on how generic types work in Go, see the [Go 1.18 Release Notes](https://go.dev/doc/go1.18) as well as the very detailed [Type Parameters Proposal](https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md).

The project is very much a **work in progress**. Here are some planned future improvements:

- I have plans to implement the minimum remaining values (MRV) heuristic, the least constraining value (LCV) heuristic, and the degree heuristic.
- It would also be nice to have some better documentation.

## Examples

An example usage of the library is provided below:

```go
// some integer variables
vars := centipede.Variables[int]{
  centipede.NewVariable("A", centipede.IntRange(1, 10)),
  centipede.NewVariable("B", centipede.IntRange(1, 10)),
  centipede.NewVariable("C", centipede.IntRange(1, 10)),
  centipede.NewVariable("D", centipede.IntRange(1, 10)),
  centipede.NewVariable("E", centipede.IntRangeStep(0, 20, 2)), // even numbers < 20
}

// numeric constraints
constraints := centipede.Constraints[int]{
  // using some constraint generators
  centipede.Equals[int]("A", "D"), // A = D
  // here we implement a custom constraint
  centipede.Constraint[int]{Vars: centipede.VariableNames{"A", "E"}, // E = A * 2
    ConstraintFunction: func(variables *centipede.Variables[int]) bool {
      if variables.Find("E").Empty || variables.Find("A").Empty {
        return true
      }
      return variables.Find("E").Value == variables.Find("A").Value*2
    }},
}
constraints = append(constraints, centipede.AllUnique[int]("A", "B", "C", "E")...) // A != B != C != E

// solve the problem
solver := centipede.NewBackTrackingCSPSolver(vars, constraints)
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
```

Unit tests have been provided to serve as additional examples of how to use the library.

## Documentation

Godocs are available [here](https://pkg.go.dev/github.com/gnboorse/centipede) on pkg.go.dev.

## Installation

```bash
go get github.com/gnboorse/centipede@v1.0.0
```

So far, this project has only been tested on macOS and Linux.

## Running tests

```bash
go test -v
```

## Contributing

Feel free to make a pull request if you spot anything out of order or want to improve the project!
