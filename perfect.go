package main

import (
	"fmt"
	"golang.org/x/exp/rand"
	"time"
)

// sliceHasSameValues determines if a uint64 slice only as the
// the same values by comparing all values to the first one
func sliceHasSameValues(s []uint64) bool {
	if len(s) < 1 {
		return false
	}
	if len(s) == 1 {
		return true
	}
	if len(s) == 2 {
		return s[0] == s[1]
	}
	for _, sum := range s {
		if s[0] != sum {
			return false
		}
	}
	return true
}

// sliceSums returns the sum of all the uint64 values
// of a given slice
// if the slice is emty 0 is returned
func sliceSum(s []uint64) uint64 {
	if len(s) < 1 {
		return 0
	}
	if len(s) == 1 {
		return s[0]
	}
	if len(s) == 2 {
		return s[0] + s[1]
	}
	var acc uint64
	for _, digit := range s {
		acc += digit
	}
	return acc
}

// A Row is a bunch of uint64s in a slice
// it also stores a sum of that slice
type Row struct {
	Values []uint64
	Sum    uint64
}

// re-calculates and sets the sum of the current row
func (r *Row) calculateSum() uint64 {
	r.Sum = sliceSum(r.Values)
	return r.Sum
}

// SquareSpec defines Characteristics that define when a Square is perfect
type SquareSpec struct {
	PerfectCols      bool
	PerfectRows      bool
	PerfectDiagonals bool
}

// Returns the default state for the SquareSpec
// -> false, false, false
func (sqs *SquareSpec) hasFieldsDefault() bool {
	return !sqs.PerfectCols && !sqs.PerfectRows && !sqs.PerfectDiagonals
}

// Sqaure is a 2D Matrix of Rows
type Square struct {
	Shape  []*Row
	Spec   *SquareSpec
	Format uint64
}

// TODO: hasMethods should be implemented on the shape
// type which should be created. Right now they use
// the pointer rec. for the Square just for Shape Access
// which is not really nice, because they dont use all the
// other fields from Square

// Dermines if a Square has perfect rows
func (s *Square) hasPerfectRows() bool {
	// TODO: error handling when Values is
	// emty or nil
	var rowSums []uint64
	for _, row := range s.Shape {
		rowSums = append(rowSums, row.calculateSum())
	}
	return sliceHasSameValues(rowSums)
}

// Dermines if a Square has perfect cols
func (s *Square) hasPerfectCols() bool {
	// TODO: error handling when Values is
	// emty or nil
	var colsSums []uint64
	var tempCol []uint64
	for row := 0; row < len(s.Shape); row++ {
		for col := 0; col < len(s.Shape[0].Values); col++ {
			tempCol = append(tempCol, s.Shape[col].Values[row])
			if len(tempCol) == len(s.Shape[0].Values) {
				colsSums = append(colsSums, sliceSum(tempCol))
				tempCol = nil
			}
		}
	}
	return sliceHasSameValues(colsSums)
}

// Dermines if a Square has perfect diagonals
func (s *Square) hasPerfectDiagonals() bool {
	// TODO: error handling when Values is
	// emty or nil
	var tempDiagForward []uint64
	var tempDiagBackward []uint64
	for row := 0; row < len(s.Shape); row++ {
		rowValues := s.Shape[row].Values
		tempDiagForward = append(tempDiagForward, rowValues[row])
		tempDiagBackward = append(tempDiagBackward, rowValues[len(s.Shape)-1-row])
	}

	var diagSums []uint64
	diagSums = append(diagSums, sliceSum(tempDiagForward))
	diagSums = append(diagSums, sliceSum(tempDiagBackward))

	return sliceHasSameValues(diagSums)
}

// isPerfect tests a Square against its spec
// FIXME: this is horrible.
func (s *Square) isPerfect() (isPerfect bool) {
	if s.Spec.hasFieldsDefault() {
		return true
	}
	if s.Spec.PerfectRows && s.hasPerfectRows() {
		isPerfect = true
	}
	if s.Spec.PerfectRows && !s.hasPerfectRows() {
		return false
	}
	if !s.Spec.PerfectRows && s.hasPerfectRows() {
		return false
	}

	if s.Spec.PerfectCols && s.hasPerfectCols() {
		isPerfect = true
	}
	if s.Spec.PerfectCols && !s.hasPerfectCols() {
		return false
	}
	if !s.Spec.PerfectCols && s.hasPerfectCols() {
		return false
	}

	if s.Spec.PerfectDiagonals && s.hasPerfectDiagonals() {
		isPerfect = true
	}
	if s.Spec.PerfectDiagonals && !s.hasPerfectDiagonals() {
		return false
	}
	if !s.Spec.PerfectDiagonals && s.hasPerfectDiagonals() {
		return false
	}

	return isPerfect
}

// generates a new `perfect` square by the SquareSpec
func (s *Square) NextBySpec(limit uint64) {
	for {
		s.generateRandom(limit)
		if s.isPerfect() {
			s.showValues()
			fmt.Printf("square has perfect rows = %v\n", s.hasPerfectRows())
			fmt.Printf("square has perfect cols = %v\n", s.hasPerfectCols())
			fmt.Printf("square has perfect diagnonals = %v\n", s.hasPerfectDiagonals())
			fmt.Printf("square is perfect = `%v` by Spec = %#v\n", s.isPerfect(), s.Spec)
			break
		}
	}
}

// Prints the values to the console
// TODO: make the shifting work for 10s,100s,...
func (s *Square) showValues() {
	for i := 0; i < len(s.Shape[0].Values); i++ {
		for _, d := range s.Shape[i].Values {
			shiftstr := ""
			if d < 10 {
				shiftstr = " "
			}
			fmt.Printf("┋%s%d", shiftstr, d)
		}
		fmt.Printf("┋\n")
	}
}

// Generate a new pseudo-random square
// TODO: optimize this:
//       - remove the possiblity for equal squares
//       - concurrent calculations for this
//       -
func (s *Square) generateRandom(limit uint64) (shape []*Row) {
	rowTemp := []uint64{}
	var i, j uint64
	source := rand.NewSource(uint64(time.Now().UnixNano()))
	random := rand.New(source)
	for i = 0; i < s.Format; i++ {
		for j = 0; j < s.Format; j++ {
			rowTemp = append(rowTemp, random.Uint64n(limit))
			if uint64(len(rowTemp)) == s.Format {
				shape = append(shape, &Row{Values: rowTemp})
				rowTemp = nil
			}
		}
	}
	s.Shape = shape
	return shape
}

func main() {
	// Defining the square
	// -> Format: (4x4)
	// -> Spec: must have perfect:
	//    - rows
	//    - cols
	//    - diagonals
	s := &Square{
		Format: 4,
		Spec: &SquareSpec{
			PerfectRows:      true,
			PerfectCols:      true,
			PerfectDiagonals: true,
		},
	}

	// This Returns the next pseudo-random
	// Square by brute-forcing all combinations
	// until a combination fulfills the spec defined
	// above
	s.NextBySpec(7)
}
