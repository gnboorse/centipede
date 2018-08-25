# Centipede - Constraint Satisfaction Problem Solver for Go

Centipede is a Constraint Satisfaction Problem solver written in Golang. [Learn more about CSPs](https://en.wikipedia.org/wiki/Constraint_satisfaction_problem). 

There is also a very informative slide deck about CSPs available from Stanford University [here](https://web.stanford.edu/class/cs227/Lectures/lec14.pdf).

## Project Status

Currently, this is very much a **work in progress**. Here are some of its limitations:

- Centipede only fully supports solving in finite domains with `int` type variables. The beginning of `string` and `float32` support has been begun but not finished. If only Go supported generics, this would be done by now...
- The search algorithm in use right now by the `IntCSPSolver` is a very simple implementation of [backtracking search](https://en.wikipedia.org/wiki/Backtracking), but I have future plans to optimize and improve this using [Arc consistency](https://en.wikipedia.org/wiki/Local_consistency#Arc_consistency). 
- I have plans to implement the minimum remaining values (MRV) heuristic and the least constraining value (LCV) heuristic.
- Additionally, I would like to implement some type of [iterative deepening search](https://en.wikipedia.org/wiki/Iterative_deepening_depth-first_search) in the solver. Right now, you can specify the `MaxDepth` field, but IDS could potentially reach a solution faster if it searched for a solution iteratively, increasing `MaxDepth` until it finds a solution.
- The project needs to be cleaned up into a real package and unit tests need to be written. It would also be nice to have some better documentation.

## Example

This is a short example of solving the map-coloring problem for Australia using Centipede.

```go
// simple implementation of the map coloring problem for Australia
colors := [3]string{"red", "green", "blue"}

// set a variable for each of the provinces
// domain for the variable is the range of colors
vars := IntVariables{
    // each has:   <name>,      <domain>
    NewIntVariable("WA", IntRange(1, len(colors)+1)),
    NewIntVariable("NT", IntRange(1, len(colors)+1)),
    NewIntVariable("Q", IntRange(1, len(colors)+1)),
    NewIntVariable("NSW", IntRange(1, len(colors)+1)),
    NewIntVariable("V", IntRange(1, len(colors)+1)),
    NewIntVariable("SA", IntRange(1, len(colors)+1)),
    NewIntVariable("T", IntRange(1, len(colors)+1)),
}

// bordering provinces cannot be equal.
// See https://en.wikipedia.org/wiki/States_and_territories_of_Australia
constraints := IntConstraints{
    NotEqualsInt("WA", "NT"),
    NotEqualsInt("WA", "SA"),
    NotEqualsInt("NT", "SA"),
    NotEqualsInt("NT", "Q"),
    NotEqualsInt("Q", "SA"),
    NotEqualsInt("Q", "NSW"),
    NotEqualsInt("NSW", "V"),
    NotEqualsInt("NSW", "SA"),
    NotEqualsInt("V", "SA"),
}

// create the solver with a maximum depth of 500
solver := NewIntCSPSolver(vars, constraints, 500)
begin := time.Now()
success := solver.Solve() // run the solution
elapsed := time.Since(begin)

if success {
    fmt.Printf("Found solution in %s\n", elapsed)
    for _, variable := range solver.State.Vars {
        // print out values for each variable
        fmt.Printf("Variable %v = %v", variable.Name, colors[variable.Value-1])
    }
} else {
    fmt.Printf("Could not find solution in %s\n", elapsed)
}

// expected output is:

// Found solution in 46.207µs
// Variable WA = red
// Variable NT = green
// Variable Q = red
// Variable NSW = green
// Variable V = red
// Variable SA = blue
// Variable T = red
```

## Usage

Have [Go](https://golang.org/) installed. Build process currently has been tested on Linux and macOS

```bash
./build.sh && cspsolver
```

## Contributing

Feel free to make a pull request if you spot anything out of order or want to improve the project.

Go is not my primary programming language, but I have been wanting to learn it for a while now. Feel free to fix anything that isn't idiomatic Go. I come from a Java/Python background. 

