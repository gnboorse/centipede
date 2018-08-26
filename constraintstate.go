package centipede

// CSPState state object for CSP Solver
type CSPState struct {
	Vars        Variables
	Constraints Constraints
	MaxDepth    int
}
