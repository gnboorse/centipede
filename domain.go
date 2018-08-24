package main

import (
	"fmt"
)

// StringDomain domain of string objects
type StringDomain []string

// IntDomain domain of int objects
type IntDomain []int

// FloatDomain domain of float32 objects
type FloatDomain []float32

// IntRange returns a slice of integers in the desired range with a step of 1
func IntRange(start int, end int) IntDomain {
	return IntRangeStep(start, end, 1)
}

// IntRangeStep returns a slice of integers in the desired range with the given step
func IntRangeStep(start int, end int, step int) IntDomain {
	rangeLength := (end - start) / step
	mod := (end - start) % step
	if mod > 0 {
		rangeLength++
	}
	intDomain := make(IntDomain, rangeLength, rangeLength)
	for i := int(0); i < rangeLength; i++ {
		intDomain[i] = i*step + start
	}
	return intDomain
}

// IntGenerator generates a series of numbers on the domain of [start, end)
// with a step size defaulting to 1
func IntGenerator(start int, end int, fx func(int) int) IntDomain {
	return IntGeneratorVariableStep(start, end, 1, fx)
}

// IntGeneratorVariableStep generates a series of numbers on the domain of [start, end)
// with a given step size for the domain
func IntGeneratorVariableStep(start int, end int, step int, fx func(int) int) IntDomain {
	rangeLength := (end - start) / step
	mod := (end - start) % step
	intDomain := make(IntDomain, rangeLength, rangeLength)
	if mod > 0 {
		rangeLength++
	}
	for i := int(0); i < rangeLength; i++ {
		// set to function value
		intDomain[i] = fx(i*step + start)
	}
	return intDomain
}

// String to string override
func (intDomain *IntDomain) String() string {
	return fmt.Sprintf("%#v %v", *intDomain, len(*intDomain))
}

// Contains slice contains method for IntDomain
func (intDomain *IntDomain) Contains(value int) bool {
	for _, item := range *intDomain {
		if item == value {
			return true
		}
	}
	return false
}

// Contains slice contains method for StringDomain
func (stringDomain *StringDomain) Contains(value string) bool {
	for _, item := range *stringDomain {
		if item == value {
			return true
		}
	}
	return false
}

// Contains slice contains method for FloatDomain
func (floatDomain *FloatDomain) Contains(value float32) bool {
	for _, item := range *floatDomain {
		if item == value {
			return true
		}
	}
	return false
}
