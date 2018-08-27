# centipede
--
    import "github.com/gnboorse/centipede"


## Usage

#### type BackTrackingCSPSolver

    type BackTrackingCSPSolver struct {
    	State CSPState
    }


BackTrackingCSPSolver struct for holding solver state

#### func  NewBackTrackingCSPSolver

    func NewBackTrackingCSPSolver(vars Variables, constraints Constraints) BackTrackingCSPSolver

NewBackTrackingCSPSolver create a solver

#### func (*BackTrackingCSPSolver) Solve

    func (solver *BackTrackingCSPSolver) Solve() bool

Solve solves for values in the CSP

#### type CSPState

    type CSPState struct {
    	Vars        Variables
    	Constraints Constraints
    }


CSPState state object for CSP Solver

#### func (*CSPState) MakeArcConsistent

    func (state *CSPState) MakeArcConsistent()

MakeArcConsistent algorithm based off of AC-3 used to make the given CSP fully
arc consistent. https://en.wikipedia.org/wiki/AC-3_algorithm

#### func (*CSPState) SimplifyPreAssignment

    func (state *CSPState) SimplifyPreAssignment()

SimplifyPreAssignment basic constraint propagation algorithm used to simplify
variable domains before solving based on variables already assigned to.
Condition: if a variable has been assigned to with a given value, remove that
value from the domain of all variables mutually exclusive to it, i.e. if A != B
and B = 2, remove 2 from the domain of A. Use of this algorith is not
recommended. Enforce arc consistency instead.

#### type Constraint

    type Constraint struct {
    	Vars               VariableNames
    	ConstraintFunction VariablesConstraintFunction
    }


Constraint CSP constraint considering integer variables

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

    func (constraint *Constraint) Satisfied(variables *Variables) bool

Satisfied checks to see if the given Constraint is satisfied by the variables
presented

#### type Constraints

    type Constraints []Constraint


Constraints collection type for Constraint

#### func  AllEquals

    func AllEquals(varnames ...VariableName) Constraints

AllEquals Constraint generator that checks that all given variables are equal

#### func  AllUnique

    func AllUnique(varnames ...VariableName) Constraints

AllUnique Constraint generator to check if all variable values are unique

#### func (*Constraints) AllSatisfied

    func (constraints *Constraints) AllSatisfied(variables *Variables) bool

AllSatisfied check if a collection of Constraints are satisfied

#### func (*Constraints) FilterByName

    func (constraints *Constraints) FilterByName(name VariableName) Constraints

FilterByName return all constraints related to a particular variable name

#### func (*Constraints) FilterByOrder

    func (constraints *Constraints) FilterByOrder(order int) Constraints

FilterByOrder return all constraints with the given order (number of related
variables)

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

#### func (Domain) Remove

    func (domain Domain) Remove(value interface{}) Domain

Remove given a value and return the updated domain

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

#### func (*Variable) SetDomain

    func (variable *Variable) SetDomain(domain Domain)

SetDomain set the domain of the given variable

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

#### func (*Variables) SetDomain

    func (variables *Variables) SetDomain(name VariableName, domain Domain)

SetDomain set the domain of the given variable by name

#### func (*Variables) SetValue

    func (variables *Variables) SetValue(name VariableName, value interface{})

SetValue setter for Variables collection

#### func (*Variables) Unassigned

    func (variables *Variables) Unassigned() int

Unassigned return the number of unassigned variables

#### func (*Variables) Unset

    func (variables *Variables) Unset(name VariableName)

Unset unset a variable with the given name

#### type VariablesConstraintFunction

    type VariablesConstraintFunction func(variables *Variables) bool


VariablesConstraintFunction function used to determine validity of Variables
