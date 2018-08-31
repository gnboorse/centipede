# Centipede - Constraint Satisfaction Problem Solver for Go

Centipede is a Constraint Satisfaction Problem solver written in Golang. [Learn more about CSPs](https://en.wikipedia.org/wiki/Constraint_satisfaction_problem). 

There is also a very informative slide deck about CSPs available from Stanford University [here](https://web.stanford.edu/class/cs227/Lectures/lec14.pdf).

## Features

- Problems are defined using sets of `Variable`, `Constraint`, and `Domain`. Some convenient generators have been provided for `Constraint` and `Domain`.
- `Variable` values can be set to values of any data type in Go (using the `interface{}` feature). However, mixing datatypes in variables that are compared to each other is not recommended. The safest approach is to set all `Variable` domains to slices of the same type.
- The search algorithm used in this library is an implementation of [backtracking search](https://en.wikipedia.org/wiki/Backtracking). 
- The solution of many complex problems can be simplified by enforcing [arc consistency](https://en.wikipedia.org/wiki/Local_consistency#Arc_consistency). This library provides an implementation of the popular [AC-3 algorithm](https://en.wikipedia.org/wiki/AC-3_algorithm) as `solver.State.MakeArcConsistent()`. Call this method before calling `solver.Solve()` to achieve best results. 
  - See the [Sudoku solver](examples/sudoku.go) for an example of how to use arc consistency. 

## Project Status

Currently, this project is very much a **work in progress**. Here are some of its limitations:

- Numeric comparison constraint (less than, greater than, etc.) generators are not yet supported, but are on the way. Variables currently use the Go `interface{}` type for their actual values, so equality and inequality are supported out of the box.
- I have plans to implement the minimum remaining values (MRV) heuristic, the least constraining value (LCV) heuristic, and the degree heuristic.
- Unit tests need to be written. It would also be nice to have some better documentation.

## Examples

For example usage of this library, see the `examples/` directory.

## Documentation

Godocs are available [here](doc/README.md).

## Installation

```bash
go get github.com/gnboorse/centipede
```

So far, this project has only been tested on macOS and Linux.

## Contributing

Feel free to make a pull request if you spot anything out of order or want to improve the project.

Go is not my primary programming language, but I have been wanting to learn it for a while now. Feel free to fix anything that isn't idiomatic Go. I come from a Java/Python background. 

# 
