package main

// IntCSPState state object for CSP Solver
type IntCSPState struct {
	Vars        IntVariables
	Constraints IntConstraints
	MaxDepth    int
}

// StringCSPState state object for CSP Solver
type StringCSPState struct {
	Vars        StringVariables
	Constraints StringConstraints
	MaxDepth    int
}

// FloatCSPState state object for CSP Solver
type FloatCSPState struct {
	Vars        FloatVariables
	Constraints FloatConstraints
	MaxDepth    int
}
