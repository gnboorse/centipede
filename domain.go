package main

import (
	"fmt"
	"math"
)

// Domain domain object
type Domain []interface{}

// FloatRange returns a slice of integers in the desired range with a step of 1
func FloatRange(start float64, end float64) Domain {
	return FloatRangeStep(start, end, 1.0)
}

// FloatRangeStep returns a slice of integers in the desired range with the given step
func FloatRangeStep(start float64, end float64, step float64) Domain {
	rangeLength := int(math.Ceil((end - start) / step))
	domain := make(Domain, rangeLength, rangeLength)
	for i := int(0); i < rangeLength; i++ {
		domain[i] = float64(i)*step + start
	}
	return domain
}

// IntRange returns a slice of integers in the desired range with a step of 1
func IntRange(start int, end int) Domain {
	return IntRangeStep(start, end, 1)
}

// IntRangeStep returns a slice of integers in the desired range with the given step
func IntRangeStep(start int, end int, step int) Domain {
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

// Generator generates a Domain from another input domain
// and a function f(x). For example:
func Generator(inputDomain Domain, fx func(interface{}) interface{}) Domain {
	outputDomain := make(Domain, 0)
	for _, input := range inputDomain {
		outputDomain = append(outputDomain, fx(input))
	}
	return outputDomain
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
