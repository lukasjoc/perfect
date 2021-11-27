package main

import (
    "fmt"
    // "errors"
)

// sliceHasSameValues determines if a uint64 slice only as the
// the same values by comparing all values to the first one
func sliceHasSameValues(s []uint64) bool {
    if len(s) < 1 { return false }
    if len(s) == 1 { return true }
    if len(s) == 2 { return s[0]==s[1] }
    for _, sum := range s{
        if(s[0] != sum) {
            return false
        }
    }
    return true
}

// sliceSums returns the sum of all the uint64 values
// of a given slice
// if the slice is emty 0 is returned
func sliceSum(s []uint64) uint64 {
    if len(s) < 1 { return 0 }
    if len(s) == 1 { return s[0] }
    if len(s) == 2 { return s[0]+s[1] }
    var acc uint64
    for _, digit := range s{
        acc+=digit
    }
    return acc
}

/* TODO:
- write function to determine magic square
- write tests to test that function
- generate square with size and number limit
- [Size: 4, Limit: 3] 4^3 + 4^1
- check if a square is a perfect square
- print a Square in pretty format to the stdout
*/

// Idea for a more structured way of
// representing the row, cols and still
// keeping connections to the data between
// rows and cols
// pointers for row values that point to specific row
// values in specific rows col1 ->(r[0][0] -> r[1][0])

// A Row is a bunch of uint64s in a slice
// it also stores a sum of that slice
type Row struct {
    Values []uint64
    Sum uint64
}

// re-calculates and sets the sum of the current row
func (r *Row) calculateSum() uint64 {
    r.Sum = sliceSum(r.Values)
    return r.Sum
}

// SquareSpec defines Characteristics that define when a Square is perfect
type SquareSpec struct {
    perfectCols bool
    perfectRows bool
    perfectDiagonals bool
}

// Sqaure is a 2D Matrix of Rows
type Square struct {
    Shape []*Row
    Spec *SquareSpec
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
            if(len(tempCol) == len(s.Shape[0].Values)) {
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
        tempDiagBackward = append(tempDiagBackward, rowValues[len(s.Shape) -1 -row])
    }

    var diagSums []uint64
    diagSums = append(diagSums, sliceSum(tempDiagForward))
    diagSums = append(diagSums, sliceSum(tempDiagBackward))

    return sliceHasSameValues(diagSums)
}

// isPerfect tests a Square against its spec
func (s *Square) isPerfect() bool {
    panic("TODO: isPerfect is not implemented yet!!")
    return false
}

func main() {
    // A demo Square
    s := &Square{
        Shape: []*Row{
            // TODO: error handling to just provide
            // square amount of Rows and RowValues
            // 2 Rows 2 Values
            // 4.Rows 4 Values
            // etc
            &Row{Values: []uint64{1,1,2}},
            &Row{Values: []uint64{1,2,1}},
            &Row{Values: []uint64{1,1,1}},
        },

        Spec: &SquareSpec {
            perfectRows: true,
            perfectCols: true,
            perfectDiagonals: true,
        },

    }

    fmt.Printf("square generated = %#v\n", s)
    fmt.Printf("square has perfect rows = %v\n", s.hasPerfectRows())
    fmt.Printf("square has perfect cols = %v\n", s.hasPerfectCols())
    fmt.Printf("square has perfect diagnonals = %v\n", s.hasPerfectDiagonals())
    fmt.Printf("square is perfect = %v\n", s.isPerfect())
}

