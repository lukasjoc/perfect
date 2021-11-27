package main

import (
    "fmt"
    // "errors"
)

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
    // TODO: error handling when Values is
    // emty or nil
    for _, digit := range r.Values {
        r.Sum += digit
    }

    fmt.Println("row sum: ", r.Values, r.Sum)

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
    var perfectRows bool

    // TODO: error handling when Values is
    // emty or nil
    // Way to check for equality is not really
    // nice but for the first try its enough

    // Defines a sum that is used to determine
    // if each row has the same sum
    var rowSums []uint64

    for _, row := range s.Shape {
        rowSums = append(rowSums, row.calculateSum())
    }

    // checks the first sum against all the other
    // ones. if the dont match then it sets the
    // perfectRows flag to false and returns the
    // check
    for _, sum := range rowSums {
        if(rowSums[0] != sum) {
            perfectRows = false
            return perfectRows
        }
        perfectRows = true
    }

    return perfectRows
}

// Dermines if a Square has perfect cols
func (s *Square) hasPerfectCols() bool {
    var perfectCols bool

    // TODO: error handling when Values is
    // emty or nil
    // Way to check for equality is not really
    // nice but for the first try its enough

    // Defines a sum that is used to determine
    // if each col has the same sum
    var colsSums []uint64

    var tempCol []uint64
    for row := 0; row < len(s.Shape); row++ {
        for col := 0; col < len(s.Shape[0].Values); col++ {
            tempCol = append(tempCol, s.Shape[col].Values[row])
            if(len(tempCol) == len(s.Shape[0].Values)) {
                var colSum uint64 = 0
                for _, digit := range tempCol {
                    colSum += digit
                }
                fmt.Println("col sum: ", tempCol, colSum)
                colsSums = append(colsSums, colSum)
                tempCol = nil
            }
        }
    }

    // checks the first sum against all the other
    // ones. if the dont match then it sets the
    // perfectRows flag to false and returns the
    // check
    for _, sum := range colsSums {
        if(colsSums[0] != sum) {
            perfectCols = false
            return perfectCols
        }
        perfectCols = true
    }

    return perfectCols
}

// Dermines if a Square has perfect diagnonals
func (s *Square) hasPerfectDiagonals() bool {
    panic("TODO: hasPerfectDiagonals is not implemented yet!!")
    return false
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
            &Row{Values: []uint64{12,18,24}},
            &Row{Values: []uint64{12,18,24}},
            &Row{Values: []uint64{12,18,24}},
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

