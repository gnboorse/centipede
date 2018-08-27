# Centipede Examples

Contained in this directory are examples of how to use Centipede for solving problems.

## Running the examples

From this `examples/` directory, run:

```bash
go build . && ./examples
```

## Examples

- [`mapcoloringaustralia.go`](mapcoloringaustralia.go) - This example solves the classic [map-coloring problem](https://en.wikipedia.org/wiki/Four_color_theorem) in three colors for the provinces of Australia.
- [`integerconstraints.go`](integerconstraints.go) - This example sets a number of variables `A`, `B`, `C`, `D`, and `E`, and searches for a solution with the following constraints: `A != B != C != E`, `A == D`, and `E == 2 * A`
- [`sudoku.go`](sudoku.go) - This example sets up and solves [this Sudoku puzzle](https://en.wikipedia.org/wiki/Sudoku#/media/File:Sudoku_Puzzle_by_L2G-20050714_standardized_layout.svg). This is a good example of a complex set of constraints being solved quite easily with the help of Arc consistency.

