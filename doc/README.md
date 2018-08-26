# centipede
--
    import "github.com/gnboorse/centipede"


## Usage

#### type CSPSolver

    type CSPSolver struct {
    	State CSPState
    }


CSPSolver struct for holding solver state

#### func  NewCSPSolver

    func NewCSPSolver(vars Variables, constraints Constraints, maxDepth int) CSPSolver

NewCSPSolver create a solver

#### func (*CSPSolver) Solve

    func (solver *CSPSolver) Solve() bool

Solve solves for values in the CSP

#### type CSPState

    type CSPState struct {
    	Vars        Variables
    	Constraints Constraints
    	MaxDepth    int
    }


CSPState state object for CSP Solver

#### type Constraint

    type Constraint struct {
    	Vars               VariableNames
    	ConstraintFunction VariablesConstraintFunction
    }


Constraint CSP constraint considering integer variables

#### func  AllEquals

    func AllEquals(varnames ...VariableName) Constraint

AllEquals Constraint generator that checks that all given variables are equal

#### func  AllUnique

    func AllUnique(varnames ...VariableName) Constraint

AllUnique Constraint generator to check if all variable values are unique

#### func  Equals

    func Equals(var1 VariableName, var2 VariableName) Constraint

Equals Constraint generator that checks if two vars are equal

#### func  NotEquals

    func NotEquals(var1 VariableName, var2 VariableName) Constraint

NotEquals Constraint generator that checks if two vars are not equal

#### func  UnaryEquals

    func UnaryEquals(var1 VariableName, value interface{}) Constraint

UnaryEquals Unary constraint that checks if var1 equals some constant

#### func  UnaryNotEquals

    func UnaryNotEquals(var1 VariableName, value interface{}) Constraint

UnaryNotEquals Unary constraint that checks if var1 is not equal to some
constant

#### func (*Constraint) Satisfied

    func (constraint *Constraint) Satisfied(variables Variables) bool

Satisfied checks to see if the given Constraint is satisfied by the variables
presented

#### type Constraints

    type Constraints []Constraint


Constraints collection type for Constraint

#### func (*Constraints) AllSatisfied

    func (constraints *Constraints) AllSatisfied(variables Variables) bool

AllSatisfied check if a collection of Constraints are satisfied

#### type Domain

    type Domain []interface{}


Domain domain object

#### func  FloatRange

    func FloatRange(start float64, end float64) Domain

FloatRange returns a slice of integers in the desired range with a step of 1

#### func  FloatRangeStep

    func FloatRangeStep(start float64, end float64, step float64) Domain

FloatRangeStep returns a slice of integers in the desired range with the given
step

#### func  Generator

    func Generator(inputDomain Domain, fx func(interface{}) interface{}) Domain

Generator generates a Domain from another input domain and a function f(x). For
example:

#### func  IntRange

    func IntRange(start int, end int) Domain

IntRange returns a slice of integers in the desired range with a step of 1

#### func  IntRangeStep

    func IntRangeStep(start int, end int, step int) Domain

IntRangeStep returns a slice of integers in the desired range with the given
step

#### func  TimeRange

    func TimeRange(start time.Time, end time.Time) Domain

TimeRange get the range of days from the start to the end time

#### func  TimeRangeStep

    func TimeRangeStep(start time.Time, end time.Time, step time.Duration) Domain

TimeRangeStep get the range of time between start to end with step as a Duration
(in nanoseconds).

#### func (*Domain) Contains

    func (domain *Domain) Contains(value interface{}) bool

Contains slice contains method for Domain

#### type Variable

    type Variable struct {
    	Name   VariableName
    	Value  interface{}
    	Domain Domain
    	Empty  bool
    }


Variable indicates a CSP variable of interface{} type

#### func  NewVariable

    func NewVariable(name VariableName, domain Domain) Variable

NewVariable constructor for Variable type

#### func (*Variable) SetValue

    func (variable *Variable) SetValue(value interface{})

SetValue setter for Variable value field

#### func (*Variable) Unset

    func (variable *Variable) Unset()

Unset the variable

#### type VariableName

    type VariableName string


VariableName is our string type for names of variables

#### type VariableNames

    type VariableNames []VariableName


VariableNames collection type for VariableName

#### func (*VariableNames) Contains

    func (varnames *VariableNames) Contains(varname VariableName) bool

Contains slice contains method for VariableNames

#### type Variables

    type Variables []Variable


Variables collection type for interface{} type variables

#### func (*Variables) Complete

    func (variables *Variables) Complete() bool

Complete indicates if all variables have been assigned to

#### func (*Variables) Contains

    func (variables *Variables) Contains(name VariableName) bool

Contains slice contains method for Variables

#### func (*Variables) Find

    func (variables *Variables) Find(name VariableName) Variable

Find find an Variable by name in an Variables collection

#### func (*Variables) SetValue

    func (variables *Variables) SetValue(name VariableName, value interface{})

SetValue setter for Variables collection

#### func (*Variables) Unassigned

    func (variables *Variables) Unassigned() Variables

Unassigned return all unassigned variables

#### type VariablesConstraintFunction

    type VariablesConstraintFunction func(variables Variables) bool


VariablesConstraintFunction function used to determine validity of Variables
