package main

import (
    "fmt"
    "golang.org/x/exp/rand"
    "time"
)

func rangeN(b, s, n uint64, include bool) (xs []uint64) {
    var i uint64 = 0;
    if include { n = n+1 }
    for i = b; i < n; i+=s {
        xs = append(xs, i)
    }
    return xs
}

// factorial of a number n
func fac(n uint64) (rs uint64) {
    var i uint64
    for _, i = range(rangeN(1, 1, n, true)) {
        if(rs==0){
            rs = i
        }else {
            rs*=i
        }
    }
    return rs
}

// https://rosettacode.org/wiki/Permutations#Go
// permuations of 1..N
func slicePermutations(begin, limit uint64) (perms [][]uint64) {
    a := rangeN(begin, 1, limit, true)
    fmt.Println("rangeN: ", a)
    n := uint64(len(a) - 1)
    var permTemp []uint64
    var c, i, j uint64
    for c = 1; c < fac(limit); c++ {
        i = n - 1
        j = n
        for a[i] > a[i+1] { i-- }
        for a[j] < a[i]   { j-- }
        a[i], a[j] = a[j], a[i]
        j = n
        i += 1
        for i < j {
            a[i], a[j] = a[j], a[i]
            i++
            j--
        }
        for _, aval := range a {
            permTemp = append(permTemp, aval)
        }
        perms = append(perms, permTemp)
        permTemp = nil
    }
    return perms
}

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
    // TODO: Primes: false,
    // TODO: PerfectMiddleDiagonals: false,
    // TODO: PerfectCorners: false,
    // TODO: PerfectUpperLowerTriangles: false
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
    next uint64
    perms [][]uint64
    Format uint64
}

// Return if the square has duplicate cell values
func (s *Square) hasDuplicateCellValues() bool {
    seen := map[uint64]uint64{};
    for _, row := range s.Shape {
        for _, value := range row.Values {
            if _, ok := seen[value]; ok {
                return false
            }
            seen[value] = value
        }
    }
    return false
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
    // TODO: make option
    //if s.hasDuplicateCellValues() == false {
    //    return false
    //}

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

// Generate a new pseudo-random square with duplicate cell values
// TODO: optimize this:
//       - remove the possiblity for equal squares
//       - concurrent calculations for this
//       - NOTE: when using the fact that each cell has to
//               contain a different number this approach does not work
//               anymore..
func (s *Square) generateRandomWithDups(limit uint64) (shape []*Row) {
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

// generatePerms generates all the permutations of a limit
// by coverting the limit to a list and returning all permutations
// of that list.
// TODO: save already generated permutations in a map
func (s *Square) generateFromPerms(start,end uint64) (shape []*Row) {
    if s.perms == nil {
        s.perms = slicePermutations(start, end)
    }
    shapeRow := []uint64{}
    s.next += 1

    var i uint64
    for i = 0; i < uint64(len(s.perms[s.next])); i++ {
        shapeRow = append(shapeRow, s.perms[s.next][i])
        if(uint64(len(shapeRow)) == s.Format) {
            shape = append(shape, &Row{Values: shapeRow})
            shapeRow = nil
        }
    }
    // fmt.Printf("shape(%s): %#d\n", shape)
    s.Shape = shape;
    return shape;
}


func main() {
    s := &Square{
        // -> Format: (3x3)
        Format: 3,

        // -> Spec: must have perfect:
        //    - can have duplicate numbers
        //    - rows
        //    - cols
        //    - diagonals
        Spec: &SquareSpec{true, true, true},
    }

    for {
    	// s.generateRandomWithDups(limit)
        s.generateFromPerms(100, 109)
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
