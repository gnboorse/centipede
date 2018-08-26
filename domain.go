package main

import (
	"fmt"
)

// Domain domain object
type Domain []interface{}

// Range returns a slice of integers in the desired range with a step of 1
func Range(start int, end int) Domain {
	return RangeStep(start, end, 1)
}

// RangeStep returns a slice of integers in the desired range with the given step
func RangeStep(start int, end int, step int) Domain {
	rangeLength := (end - start) / step
	mod := (end - start) % step
	if mod > 0 {
		rangeLength++
	}
	domain := make(Domain, rangeLength, rangeLength)
	for i := int(0); i < rangeLength; i++ {
		domain[i] = i*step + start
	}
	return domain
}

// Generator generates a series of numbers on the domain of [start, end)
// with a step size defaulting to 1
func Generator(start int, end int, fx func(int) int) Domain {
	return GeneratorVariableStep(start, end, 1, fx)
}

// GeneratorVariableStep generates a series of numbers on the domain of [start, end)
// with a given step size for the domain
func GeneratorVariableStep(start int, end int, step int, fx func(int) int) Domain {
	rangeLength := (end - start) / step
	mod := (end - start) % step
	domain := make(Domain, rangeLength, rangeLength)
	if mod > 0 {
		rangeLength++
	}
	for i := int(0); i < rangeLength; i++ {
		// set to function value
		domain[i] = fx(i*step + start)
	}
	return domain
}

// String to string override
func (domain *Domain) String() string {
	return fmt.Sprintf("%#v %v", *domain, len(*domain))
}

// Contains slice contains method for Domain
func (domain *Domain) Contains(value interface{}) bool {
	for _, item := range *domain {
		if item == value {
			return true
		}
	}
	return false
}
