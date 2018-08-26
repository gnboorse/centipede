# Centipede - Constraint Satisfaction Problem Solver for Go

Centipede is a Constraint Satisfaction Problem solver written in Golang. [Learn more about CSPs](https://en.wikipedia.org/wiki/Constraint_satisfaction_problem). 

There is also a very informative slide deck about CSPs available from Stanford University [here](https://web.stanford.edu/class/cs227/Lectures/lec14.pdf).

## Project Status

Currently, this is very much a **work in progress**. Here are some of its limitations:

- Numeric comparison constraint (less than, greater than, etc.) generators are not yet supported, but are on the way. Variables currently use the Go `interface{}` type for their actual values, so equality and inequality are supported out of the box.
- The search algorithm in use right now by the `CSPSolver` is a very simple implementation of [backtracking search](https://en.wikipedia.org/wiki/Backtracking), but I have future plans to optimize and improve this using [Arc consistency](https://en.wikipedia.org/wiki/Local_consistency#Arc_consistency). 
- I have plans to implement the minimum remaining values (MRV) heuristic and the least constraining value (LCV) heuristic.
- Additionally, I would like to implement some type of [iterative deepening search](https://en.wikipedia.org/wiki/Iterative_deepening_depth-first_search) in the solver. Right now, you can specify the `MaxDepth` field, but IDS could potentially reach a solution faster if it searched for a solution iteratively, increasing `MaxDepth` until it finds a solution.
- Unit tests need to be written. It would also be nice to have some better documentation.

## Examples

For example code, see the `examples/` directory.

## Installation

```bash
go get github.com/gnboorse/centipede
```

So far, this project has only been tested on macOS and Linux.

## Contributing

Feel free to make a pull request if you spot anything out of order or want to improve the project.

Go is not my primary programming language, but I have been wanting to learn it for a while now. Feel free to fix anything that isn't idiomatic Go. I come from a Java/Python background. 

# 